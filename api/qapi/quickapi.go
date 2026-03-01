package qapi

import (
	"fmt"

	"github.com/Meduzz/dsl/endpoint"
	"github.com/Meduzz/dsl/service"
	"github.com/Meduzz/quickapi/model"
)

const (
	plain = "/%s/"
	id    = "/%s/:id"
)

var (
	mapper = make(map[string]string)
)

func init() {
	mapper["search"] = "read"
	mapper["patch"] = "update"
}

func Quickapi(builder service.ServiceBuilder, prefix, name string, entity model.Entity) {
	ep, url, permission := tupel("create", plain, prefix, name)
	builder.AddEndpoint(ep, func(eb endpoint.EndpointBuilder) {
		eb.POST(url)
		eb.RequestBody(endpoint.JSON, entity.Create())
		eb.ResponseBody(endpoint.JSON, entity.Create())
		eb.SetPermission(permission)
	})

	ep, url, permission = tupel("read", id, prefix, name)
	builder.AddEndpoint(ep, func(eb endpoint.EndpointBuilder) {
		eb.GET(url)
		eb.Path("id")
		eb.ResponseBody(endpoint.JSON, entity.Create())
		eb.SetPermission(permission)
	})

	ep, url, permission = tupel("update", id, prefix, name)
	builder.AddEndpoint(ep, func(eb endpoint.EndpointBuilder) {
		eb.PUT(url)
		eb.Path("id")
		eb.RequestBody(endpoint.JSON, entity.Create())
		eb.ResponseBody(endpoint.JSON, entity.Create())
		eb.SetPermission(permission)
	})

	ep, url, permission = tupel("delete", id, prefix, name)
	builder.AddEndpoint(ep, func(eb endpoint.EndpointBuilder) {
		eb.DELETE(url)
		eb.Path("id")
		eb.SetPermission(permission)
	})

	ep, url, permission = tupel("search", plain, prefix, name)
	builder.AddEndpoint(ep, func(eb endpoint.EndpointBuilder) {
		eb.GET(url)
		eb.Query("skip")
		eb.Query("take")
		eb.QueryMap("where")
		eb.QueryMap("sort")
		eb.ResponseBody(endpoint.JSON, entity.CreateArray())
		eb.SetPermission(permission)
	})

	ep, url, permission = tupel("patch", id, prefix, name)
	builder.AddEndpoint(ep, func(eb endpoint.EndpointBuilder) {
		eb.PATCH(url)
		eb.Path("id")
		eb.RequestBody(endpoint.JSON, entity.Create())
		eb.ResponseBody(endpoint.JSON, entity.Create())
		eb.SetPermission(permission)
	})
}

func createUrl(url, prefix, name string) string {
	if prefix != "" {
		return fmt.Sprintf("%s%s", prefix, fmt.Sprintf(url, name))
	}

	return fmt.Sprintf(url, name)
}

func endpointName(verb, name string) string {
	return fmt.Sprintf("%s/%s", name, verb)
}

func permission(verb, name string) string {
	format := "%s.%s"
	mapped, ok := mapper[verb]

	if ok {
		return fmt.Sprintf(format, name, mapped)
	}

	return fmt.Sprintf(format, name, verb)
}

func tupel(verb, format, prefix, name string) (string, string, string) {
	ep := endpointName(verb, name)
	perm := permission(verb, name)
	url := createUrl(format, prefix, name)

	return ep, url, perm
}
