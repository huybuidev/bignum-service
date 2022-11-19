package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"bignum-service/pkg/jrpcserver"
)

// BigNumSvc is the RPC service placeholder for the Big Number Service
type BigNumSvc struct{}

func (bns *BigNumSvc) Create(params []string, reply *string) error {
	log.Debug().Interface("params", params).Str("method", "create").Msg("")

	// First param is the number object's name
	// The number object's value must be provided (the second param)
	if len(params) < 2 {
		return errors.New("value must be provided")
	}

	a := &big.Float{}
	_, ok := a.SetString(params[1])
	if !ok {
		return fmt.Errorf("invalid big float number %s", params[1])
	}

	*reply = "created"
	return nil
}

func paramsParser(inputParamsJSONRaw []byte) (rpcParamsJSONRaw []byte, err error) {
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

func main() {
	setupLogging()

	s := jrpcserver.New(map[string]string{
		"create": "BigNumSvc.Create",
	}, paramsParser)

	BigNumSvc := new(BigNumSvc)
	s.Register(BigNumSvc)
	http.Handle("/rpc", s)
	http.ListenAndServe(":8080", nil)
}

func setupLogging() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
}
