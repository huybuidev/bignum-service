package main

import (
	"net/http"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"bignum-service/app/bignumrpc"
	"bignum-service/domain/bignum/bignummem"
	"bignum-service/lib/jrpcserver"
	"bignum-service/service/numobj"
)

func main() {
	setupLogging()

	s := jrpcserver.NewServerWithStringArrayParams(map[string]string{
		"create":   "BigNumRPCService.Create",
		"update":   "BigNumRPCService.Update",
		"delete":   "BigNumRPCService.Delete",
		"add":      "BigNumRPCService.Add",
		"subtract": "BigNumRPCService.Subtract",
		"multiply": "BigNumRPCService.Multiply",
		"divide":   "BigNumRPCService.Divide",
	})

	numObjSvc, err := numobj.NewService(numobj.WithBignumMemRepo(bignummem.NewBignumMemoryRepository()))
	if err != nil {
		log.Fatal().Err(err).Msg("failed to init number object service")
	}
	bignumRPCSvc := bignumrpc.New(numObjSvc)

	s.Register(bignumRPCSvc)
	http.Handle("/rpc", s)
	http.ListenAndServe(":8080", nil)
}

func setupLogging() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
}
