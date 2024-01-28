package config

import (
	"os/exec"
)

func InitProto(path string) error {
	//protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative gRPC/proto/auth.proto
	cmd := exec.Command("protoc",
		"--go_out=.",
		"--go_opt=paths=source_relative",
		"--go-grpc_out=.",
		"--go-grpc_opt=paths=source_relative",
		path)
	if _, err := cmd.CombinedOutput(); err != nil {
		panic(err)
	} else {
		return nil
	}

}
