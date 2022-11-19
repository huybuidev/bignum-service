package jrpcserver

import (
	"bignum-service/pkg/kjsonrpc"
	"io"
	"log"
	"net/http"
	"net/rpc"
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
	log.Println("got a request")
	codec := kjsonrpc.NewServerCodec(&HttpConn{in: req.Body, out: w}, s.methodMap, s.paramsParser)
	log.Println("ServeCodec")
	s.rpcServer.ServeCodec(codec)
	log.Println("finished serving request")
}
