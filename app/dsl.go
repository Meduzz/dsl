package app

import (
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
