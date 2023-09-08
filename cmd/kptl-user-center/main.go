package main

import (
	"fmt"
	"github.com/803Studio/kptl-grpc-go/pkg/usercenter"
	"github.com/803Studio/kptl-user-center/internal/auth"
	"github.com/803Studio/kptl-user-center/internal/config"
	"github.com/803Studio/kptl-user-center/internal/data"
	"github.com/803Studio/kptl-user-center/internal/handler"
	"google.golang.org/grpc"
	"log"
	"net"
)

type InitFn func() error

func startRpcServer() error {
	socket, err := net.Listen("tcp", fmt.Sprintf(":%d", config.AppConfig.Port))
	if err != nil {
		return err
	}

	rpcServer := grpc.NewServer()
	usercenter.RegisterUserServer(rpcServer, &handler.KptlUserServer{})

	err = rpcServer.Serve(socket)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	for _, task := range []InitFn{
		config.Init,
		data.Init,
		auth.Init,
		startRpcServer,
	} {
		err := task()
		if err != nil {
			log.Fatal(err)
		}
	}
}
