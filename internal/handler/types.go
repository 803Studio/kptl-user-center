package handler

import "github.com/803Studio/kptl-grpc-go/pkg/usercenter"

type KptlUserServer struct {
	usercenter.UnimplementedUserServer
}

var (
	_ usercenter.UserServer = (*KptlUserServer)(nil)
)
