// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"tgbot-service/internal/biz"
	"tgbot-service/internal/conf"
	"tgbot-service/internal/data"
	"tgbot-service/internal/server"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	grpcServer := server.NewGRPCServer(confServer, logger)
	db := data.NewDB(confData)
	dataData, cleanup, err := data.NewData(confData, logger, db)
	if err != nil {
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData)
	bot := data.NewTelebot(confData)
	userUseCase := biz.NewUserUseCase(userRepo, logger, bot)
	telebotBot := server.NewTelebot(userUseCase, bot)
	rabbitConn := server.NewRabbitConn(confData, userUseCase)
	app := newApp(logger, grpcServer, telebotBot, rabbitConn)
	return app, func() {
		cleanup()
	}, nil
}
