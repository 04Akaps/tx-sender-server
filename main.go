package main

import (
	"flag"
	"github.com/04Akaps/tx-sender-server/cmd/app"
	"github.com/04Akaps/tx-sender-server/config"
	gRPC "github.com/04Akaps/tx-sender-server/gRPC/server"
)

var protoFlag = flag.String("proto", "gRPC/proto/auth.proto", "proto path")
var configFlag = flag.String("config", "./config.toml", "config path")
var logFlag = flag.String("log", "tx-server.txt", "log name")

func main() {
	flag.Parsed()

	// set Log File
	config.NewLog(*logFlag)

	// update Proto File Latest
	config.InitProto(*protoFlag)

	cfg := config.NewConfig(*configFlag)

	// Auth를 처리 할 gRPC 서버를 먼저 구동시켜 둔다.
	if err := gRPC.NewGrpcServer(cfg); err != nil {
		panic(err)
	} else {
		txApp := app.NewTxApp(cfg)
		txApp.StartTxApp()
	}
}
