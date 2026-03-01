package app

import (
	"github.com/Meduzz/dsl/policy"
	"github.com/Meduzz/dsl/service"
)

type (
	App struct {
		Name        string             `json:"name"`
		Description string             `json:"description,omitempty"`
		Services    []*service.Service `json:"services"`
		Policy      *policy.Policy     `json:"policy,omitempty"`
		Domain      string             `json:"domain,omitempty"`
		ContextPath string             `json:"context,omitempty"`
	}

	AppBuilder interface {
		SetDescription(string)
		SetDomain(string)
		SetContextPath(string)
		AddService(string, func(service.ServiceBuilder))
		Policy(func(policy.PolicyBuilder))
	}

	appBuilder struct {
		app *App
	}
)

var (
	_ AppBuilder = &appBuilder{}
)
