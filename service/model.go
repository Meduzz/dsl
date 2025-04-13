package service

import (
	"github.com/Meduzz/dsl/api"
	"github.com/Meduzz/dsl/deploy"
	"github.com/Meduzz/dsl/proxy"
)

type (
	// ServiceKind - Quickapi, Gin or RPC
	ServiceKind string

	// Service - describes a service
	Service struct {
		Name        string         `json:"name"`
		Description string         `json:"description,omitempty"`
		Kind        ServiceKind    `json:"kind"`
		Tags        []string       `json:"tags,omitempty"`
		Api         *api.Api       `json:"api,omitempty"`
		Deploy      *deploy.Deploy `json:"deploy,omitempty"`
		Proxy       *proxy.Proxy   `json:"proxy,omitempty"`
	}
)
