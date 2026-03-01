package app

import (
	"github.com/Meduzz/dsl/policy"
	"github.com/Meduzz/dsl/service"
)

func NewApp(name string, cb func(AppBuilder)) *App {
	app := &App{}
	app.Name = name

	builder := &appBuilder{app}

	cb(builder)

	return app
}

func (a *appBuilder) SetDescription(description string) {
	a.app.Description = description
}

func (a *appBuilder) SetDomain(domain string) {
	a.app.Domain = domain
}

func (a *appBuilder) SetContextPath(path string) {
	a.app.ContextPath = path
}

func (a *appBuilder) AddService(name string, cb func(service.ServiceBuilder)) {
	a.app.Services = append(a.app.Services, service.NewService(name, cb))
}

func (a *appBuilder) Policy(cb func(policy.PolicyBuilder)) {
	a.app.Policy = &policy.Policy{}
	builder := policy.NewPolicyBuilder(a.app.Policy)

	cb(builder)
}
