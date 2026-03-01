package service

import (
	"github.com/Meduzz/dsl/endpoint"
	serviceref "github.com/Meduzz/dsl/serviceRef"
	"github.com/Meduzz/helper/fp/slice"
	"github.com/Meduzz/helper/meta/schema"
)

func NewService(name string, cb func(ServiceBuilder)) *Service {
	s := &Service{}
	s.Name = name

	builder := &serviceBuilder{s}
	cb(builder)

	return s
}

func (s *serviceBuilder) SetDescription(description string) {
	s.service.Description = description
}

func (s *serviceBuilder) AddTags(tags ...string) {
	s.service.Tags = append(s.service.Tags, tags...)
}

func (s *serviceBuilder) AddEndpoint(name string, cb func(endpoint.EndpointBuilder)) {
	s.service.Endpoints = append(s.service.Endpoints, endpoint.NewEndpoint(name, cb))
}

func (s *serviceBuilder) AddDependencies(deps ...serviceref.ServiceRef) {
	s.service.Dependency = slice.Concat(s.service.Dependency, slice.Filter(deps, func(s serviceref.ServiceRef) bool {
		return s.Valid()
	}))
}

func (s *serviceBuilder) AddDependency(app, service string) {
	ref := serviceref.NewServiceRef(app, service)

	if ref.Valid() {
		s.service.Dependency = append(s.service.Dependency, ref)
	}
}

func (s *serviceBuilder) CLI(name string) {
	s.service.Packaging = &Packaging{
		Kind: CliKind,
		Name: name,
	}
}

func (s *serviceBuilder) Container(name string) {
	s.service.Packaging = &Packaging{
		Kind: ContainerKind,
		Name: name,
	}
}

func (s *serviceBuilder) AddEvent(topic string, event any) {
	sh := schema.SchemaFor(event)

	s.service.Events = append(s.service.Events, &Event{
		Topic:  topic,
		Schema: sh,
	})
}
