// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"go-micro.dev/v4/web"
	"paper-translation/pkg/oss"
	"paper-translation/pkg/service"
)

// Injectors from wire.go:

func InitApp() web.Service {
	registry := service.NewRegistry()
	config := service.NewConfig()
	fileService := NewFileService(registry)
	paperService := NewPaperService(registry)
	client := oss.NewAliYunOSS(config)
	engine := NewRoute(fileService, paperService, client)
	webService := NewService(registry, config, engine)
	return webService
}