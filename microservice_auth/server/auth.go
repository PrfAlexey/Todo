package server

import (
	"Todo/microservice_auth/proto"
	"Todo/microservice_auth/session"
	"Todo/microservice_auth/user"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type AuthServer struct {
	service     session.Service
	userService user.Service
}

func NewAuthServer(service session.Service, u user.Service) *AuthServer {
	return &AuthServer{service: service, userService: u}
}

func (a *AuthServer) Login(c context.Context, usr *proto.User) (*proto.LoginAnswer, error) {
	userID, err := a.userService.CheckUser(usr.Login, usr.Password)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	sessionValue, err1 := a.service.CreateSession(userID)
	if err1 != nil {
		return nil, status.Error(codes.InvalidArgument, err1.Error())
	}

	return &proto.LoginAnswer{Value: sessionValue, Flag: false}, nil
}

func (a *AuthServer) Check(c context.Context, s *proto.Session) (*proto.CheckAnswer, error) {
	userID, err := a.service.Check(s.Value)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &proto.CheckAnswer{UserID: userID, Flag: false}, nil
}

func (a *AuthServer) Logout(c context.Context, s *proto.Session) (*proto.LogoutAnswer, error) {
	_, err := a.service.Check(s.Value)
	if err != nil {
		return &proto.LogoutAnswer{}, status.Error(codes.Internal, err.Error())
	}

	err = a.service.Logout(s.Value)
	if err != nil {
		log.Println(err)
		return &proto.LogoutAnswer{}, err
	}

	return &proto.LogoutAnswer{Flag: false}, nil
}