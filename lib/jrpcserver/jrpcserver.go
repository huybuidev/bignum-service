package jrpcserver

import (
	"encoding/json"
	"io"
	"net/http"
	"net/rpc"

	"github.com/rs/zerolog/log"

	"bignum-service/lib/kjsonrpc"
)

type HttpConn struct {
	in  io.Reader
	out io.Writer
}

func (c *HttpConn) Read(p []byte) (n int, err error)  { return c.in.Read(p) }
func (c *HttpConn) Write(d []byte) (n int, err error) { return c.out.Write(d) }
func (c *HttpConn) Close() error                      { return nil }

type Server struct {
	rpcServer    *rpc.Server
	methodMap    map[string]string
	paramsParser kjsonrpc.ParamParserFunc
}

// New a JSON-RPC server that supports:
//
//  1. Custom method mapping
//     The standard net/rpc/json only support RPC-style of method value: `service.method`
//     In order to correctly invoke the method handler, we have to pass method value that
//     follows this style.
//
//     In order to support custom method name, we can make a custom method mapping:
//     customMethodName -> rpcTypeMethodName
//     E.g.: methodMap = map[string]string{"create": "BigNumSvc.Create"}
//     -> If the input method value in the JSON request is "create", it will be translate to "BigNumSvc.Create"
//     before passing down as RPC message.
//
//     If the input method could not be found in the methodMap (or methodMap is nil/empty),
//     The method value will be kept as it is.
//
//  2. Custom params parser function
//     JSON-RPC params is an array of values.
//     RPC params is an array of struct/object and only the first element will be used.
//     -> we need a custom params parser function which takes JSON raw bytes
//     for `params` as input. Then it is required to output JSON raw bytes
//     for `params` as correct format for RPC params.
func New(methodMap map[string]string, paramsParser ...kjsonrpc.ParamParserFunc) (jrpcServer *Server) {
	var paramsParserFunc kjsonrpc.ParamParserFunc
	if len(paramsParser) > 0 {
		paramsParserFunc = paramsParser[0]
	}
	return &Server{
		rpcServer:    rpc.NewServer(),
		methodMap:    methodMap,
		paramsParser: paramsParserFunc,
	}
}

func (s *Server) RegisterName(name string, rcvr any) {
	s.rpcServer.RegisterName(name, rcvr)
}

func (s *Server) Register(rcvr any) {
	s.rpcServer.Register(rcvr)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.Debug().Msg("got a request")
	codec := kjsonrpc.NewServerCodec(&HttpConn{in: req.Body, out: w}, s.methodMap, s.paramsParser)
	s.rpcServer.ServeCodec(codec)
	log.Debug().Msg("finished serving request")
}

func stringArrayParamsParser(inputParamsJSONRaw []byte) (rpcParamsJSONRaw []byte, err error) {
	// Input params should be an array of string
	jrpcParams := []string{}
	err = json.Unmarshal(inputParamsJSONRaw, &jrpcParams)
	if err != nil {
		return nil, err
	}
	// As normal RPC will only consider the first element in the field `params`,
	// We put the input params as array of string as the first element in the array
	jb, err := json.Marshal([][]string{jrpcParams})
	log.Debug().Str("params", string(jb)).Msg("")
	return jb, err
}

func NewServerWithStringArrayParams(methodMap map[string]string) (jrpcServer *Server) {
	return New(methodMap, stringArrayParamsParser)
}
