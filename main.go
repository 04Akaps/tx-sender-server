package main

import (
	"flag"
	"github.com/04Akaps/tx-sender-server/cmd/auth"
	"github.com/04Akaps/tx-sender-server/cmd/tx"
)

var protoFlag = flag.String("proto", "gRPC/proto/auth.proto", "proto path")
var configFlag = flag.String("config", "./config.toml", "config path")

func main() {

	//protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative gRPC/proto/auth.proto

	txApp := tx.NewTxApp()
	authApp := auth.NewAuthApp()

	txApp.StartTxApp()
	authApp.StartAuthApp()
}
