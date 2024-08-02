// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"filesharer/internal/biz"
	"filesharer/internal/conf"
	"filesharer/internal/data"
	"filesharer/internal/server"
	"filesharer/internal/service"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/hashicorp/consul/api"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, bootstrap *conf.Bootstrap, logger log.Logger) (*kratos.App, func(), error) {
	registrar := InitRegistry(bootstrap)
	fileClient := data.NewFileClient(bootstrap)
	getConsulInfoClient := data.NewGetConsulInfoClient(bootstrap)
	dataData, cleanup, err := data.NewData(confData, logger, fileClient, getConsulInfoClient)
	if err != nil {
		return nil, nil, err
	}
	filesharerRepo := data.NewFilesharerRepo(dataData, logger)
	filesharerUsecase := biz.NewFilesharerUsecase(filesharerRepo, logger)
	fileService := service.NewFileService(filesharerUsecase)
	grpcServer := server.NewGRPCServer(confServer, fileService, logger)
	httpServer := server.NewHTTPServer(confServer, fileService, logger)
	app := newApp(logger, registrar, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}

// wire.go:

func InitRegistry(bc *conf.Bootstrap) registry.Registrar {
	cfg := bc.Consul

	defaultConfig := api.DefaultConfig()
	defaultConfig.Address = cfg.Addr
	client, err := api.NewClient(defaultConfig)
	if err != nil {
		panic(err)
	}
	reg := consul.New(client, consul.WithHealthCheck(true))

	return reg
}
