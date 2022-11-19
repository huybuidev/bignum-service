package main

import (
	"net/http"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"bignum-service/pkg/bignum"
	"bignum-service/pkg/jrpcserver"
)

func main() {
	setupLogging()

	s := jrpcserver.NewServerWithStringArrayParams(map[string]string{
		"create": "BigNumService.Create",
	})

	s.Register(&bignum.BigNumService{})
	http.Handle("/rpc", s)
	http.ListenAndServe(":8080", nil)
}

func setupLogging() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
}
