package service

import (
	"github.com/Meduzz/dsl/endpoint"
	serviceref "github.com/Meduzz/dsl/serviceRef"
	"github.com/Meduzz/helper/meta/schema"
)

type (
	PackageKind string // container/cli

	// Service - describes a service
	Service struct {
		Name        string                  `json:"name"`
		Description string                  `json:"description,omitempty"`
		Tags        []string                `json:"tags,omitempty"`
		Endpoints   []*endpoint.Endpoint    `json:"endpoints"`
		Dependency  []serviceref.ServiceRef `json:"dependencies,omitempty"`
		Packaging   *Packaging              `json:"packaging,omitempty"`
		Events      []*Event                `json:"events,omitempty"`
	}

	ServiceBuilder interface {
		SetDescription(string)
		AddTags(...string)
		AddEndpoint(string, func(endpoint.EndpointBuilder))
		AddDependencies(...serviceref.ServiceRef)
		AddDependency(app, service string)
		CLI(string)
		Container(string)
		AddEvent(topic string, event any)
	}

	serviceBuilder struct {
		service *Service
	}

	Packaging struct {
		Kind PackageKind `json:"kind"`
		Name string      `json:"name"`
	}

	Event struct {
		Topic  string         `json:"topic"`
		Schema *schema.Schema `json:"schema"`
	}
)

var (
	_ ServiceBuilder = &serviceBuilder{}

	ContainerKind = PackageKind("container")
	CliKind       = PackageKind("cli")
)
