// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"go-micro.dev/v4"
	email2 "paper-translation/app/email/service/email"
	"paper-translation/pkg/email"
	"paper-translation/pkg/service"
)

// Injectors from wire.go:

func InitApp() micro.Service {
	registry := service.NewRegistry()
	config := service.NewConfig()
	pool := email.NewEmailPool(config)
	emailService := email2.NewEmailService(pool, config)
	microService := NewService(registry, config, emailService)
	return microService
}