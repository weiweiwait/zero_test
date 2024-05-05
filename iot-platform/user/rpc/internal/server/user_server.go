// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"iot-platform/user/rpc/internal/logic"
	"iot-platform/user/rpc/internal/svc"
	"iot-platform/user/rpc/types/user"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) Auth(ctx context.Context, in *user.UserAuthRequest) (*user.UserAuthReply, error) {
	l := logic.NewAuthLogic(ctx, s.svcCtx)
	return l.Auth(in)
}
