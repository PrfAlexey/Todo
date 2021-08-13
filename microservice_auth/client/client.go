package client

import (
	"Todo/microservice_auth/proto"
	"context"
	"github.com/labstack/echo"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

const (
	Localhost = "127.0.0.1"
)

type AuthClient struct {
	client proto.AuthClient
	gConn  *grpc.ClientConn
}

func NewAuthClient(port string) (IAuthClient, error) {
	gConn, err := grpc.Dial(
		Localhost+port,
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	return &AuthClient{client: proto.NewAuthClient(gConn), gConn: gConn}, nil
}

func (a *AuthClient) Login(login string, password string) (string, error) {
	usr := &proto.User{
		Login:    login,
		Password: password,
	}

	answer, err := a.client.Login(context.Background(), usr)
	if err != nil {
		return "", err
	}

	return answer.Value, nil
}

func (a *AuthClient) Check(value string) (uint64, error) {
	sessionValue := &proto.Session{Value: value}

	answer, err := a.client.Check(context.Background(), sessionValue)
	if err != nil {
		return 0, err
	}

	return answer.UserID, nil
}

func (a *AuthClient) Logout(value string) error {
	sessionValue := &proto.Session{Value: value}

	answer, err := a.client.Logout(context.Background(), sessionValue)
	if err != nil {
		return err
	}
	if answer.Flag {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	return nil
}

func (a *AuthClient) Close() {
	if err := a.gConn.Close(); err != nil {
		log.Println(err)
	}
}
