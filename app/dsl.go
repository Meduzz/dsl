package app

import (
	"github.com/Meduzz/dsl/policy"
	"github.com/Meduzz/dsl/service"
)

func NewApp(name string) *App {
	app := &App{}

	app.Name = name

	return app
}

func (a *App) AddService(service *service.Service) *service.Service {
	a.Services = append(a.Services, service)
	return service
}

func (a *App) SetPolicy() *policy.Policy {
	p := &policy.Policy{}
	a.Policy = p
	return p
}
