// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"ciam-rebac/internal/biz"
	"ciam-rebac/internal/conf"
	"ciam-rebac/internal/data"
	"ciam-rebac/internal/server"
	"ciam-rebac/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	spiceDbRepository, cleanup, err := data.NewSpiceDbRepository(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	createRelationshipsUsecase := biz.NewCreateRelationshipsUsecase(spiceDbRepository, logger)
	readRelationshipsUsecase := biz.NewReadRelationshipsUsecase(spiceDbRepository, logger)
	deleteRelationshipsUsecase := biz.NewDeleteRelationshipsUsecase(spiceDbRepository, logger)
	relationshipsService := service.NewRelationshipsService(logger, createRelationshipsUsecase, readRelationshipsUsecase, deleteRelationshipsUsecase)
	healthService := service.NewHealthService()
	checkUsecase := biz.NewCheckUsecase(spiceDbRepository, logger)
	checkService := service.NewCheckService(logger, checkUsecase)
	getSubjectsUsecase := biz.NewGetSubjectsUseCase(spiceDbRepository, logger)
	lookupService := service.NewLookupSubjectsService(getSubjectsUsecase)
	grpcServer := server.NewGRPCServer(confServer, relationshipsService, healthService, checkService, lookupService, logger)
	httpServer := server.NewHTTPServer(confServer, relationshipsService, healthService, checkService, lookupService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
