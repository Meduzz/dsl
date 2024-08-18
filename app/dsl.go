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

func (a *App) AddService(name string, kind service.ServiceKind) *service.Service {
	s := &service.Service{}

	s.Name = name
	s.Kind = kind

	a.Services = append(a.Services, s)
	return s
}

func (a *App) GetPolicy() *policy.Policy {
	if a.Policy != nil {
		return a.Policy
	}

	p := &policy.Policy{}
	a.Policy = p
	return p
}
