package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"bignum-service/pkg/jrpcserver"
)

type NumObject struct {
	Name  *string `json:"name"`
	Value *string `json:"value"`
}

type BigNumSvc struct{}

func (bns *BigNumSvc) Create(numObj *NumObject, reply *string) error {
	log.Debug().Interface("numObj", numObj).Str("method", "create").Msg("")
	if numObj.Value == nil {
		return errors.New("value must be provided")
	}
	*reply = "created"
	return nil
}

func paramsParser(inputParamsJSONRaw []byte) (rpcParamsJSONRaw []byte, err error) {
	jrpcParams := []string{}
	err = json.Unmarshal(inputParamsJSONRaw, &jrpcParams)
	if err != nil {
		return nil, err
	}
	if len(jrpcParams) > 2 || len(jrpcParams) == 0 {
		return nil, errors.New("only accept one or two params")
	}
	numObj := &NumObject{
		Name: &jrpcParams[0],
	}
	if len(jrpcParams) == 2 {
		numObj.Value = &jrpcParams[1]
	}
	jb, err := json.Marshal([]*NumObject{numObj})
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
