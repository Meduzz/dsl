package endpoint

import "github.com/Meduzz/helper/meta/schema"

func NewEndpoint(name string, cb func(EndpointBuilder)) *Endpoint {
	ep := &Endpoint{
		Name: name,
	}

	builder := &endpointBuilder{ep}

	cb(builder)

	return ep
}

func (e *endpointBuilder) SetDescription(desc string) {
	e.endpoint.Description = desc
}

func (e *endpointBuilder) SetPermission(permission string) {
	e.endpoint.Permission = permission
}

func (e *endpointBuilder) Topic(topic, consumerGroup string) {
	e.endpoint.Route = &Route{
		Kind:          RpcKind,
		Topic:         topic,
		ConsumerGroup: consumerGroup,
	}
}

func (e *endpointBuilder) GET(path string) {
	e.endpoint.Route = &Route{
		Kind:   HttpKind,
		Method: "GET",
		Path:   path,
	}
}

func (e *endpointBuilder) POST(path string) {
	e.endpoint.Route = &Route{
		Kind:   HttpKind,
		Method: "POST",
		Path:   path,
	}
}

func (e *endpointBuilder) PUT(path string) {
	e.endpoint.Route = &Route{
		Kind:   HttpKind,
		Method: "PUT",
		Path:   path,
	}
}

func (e *endpointBuilder) DELETE(path string) {
	e.endpoint.Route = &Route{
		Kind:   HttpKind,
		Method: "DELETE",
		Path:   path,
	}
}

func (e *endpointBuilder) PATCH(path string) {
	e.endpoint.Route = &Route{
		Kind:   HttpKind,
		Method: "PATCH",
		Path:   path,
	}
}

func (e *endpointBuilder) RequestHeader(name string) {
	e.endpoint.Route.In = append(e.endpoint.Route.In, param(HeaderKind, name, "", nil))
}

func (e *endpointBuilder) Query(name string) {
	e.endpoint.Route.In = append(e.endpoint.Route.In, param(QueryKind, name, "", nil))
}

func (e *endpointBuilder) QueryMap(name string) {
	e.endpoint.Route.In = append(e.endpoint.Route.In, param(QueryKind, name, "", nil))
}

func (e *endpointBuilder) Path(name string) {
	e.endpoint.Route.In = append(e.endpoint.Route.In, param(PathKind, name, "", nil))
}

func (e *endpointBuilder) RequestBody(contentType string, payload any) {
	e.endpoint.Route.In = append(e.endpoint.Route.In, param(BodyKind, "", contentType, payload))
}

func (e *endpointBuilder) ResponseHeader(name string) {
	e.endpoint.Route.Out = append(e.endpoint.Route.Out, param(HeaderKind, name, "", nil))
}

func (e *endpointBuilder) ResponseBody(contentType string, payload any) {
	e.endpoint.Route.Out = append(e.endpoint.Route.Out, param(BodyKind, "", contentType, payload))
}

func param(kind ParamKind, name, contentType string, typ any) *Param {
	p := &Param{
		Kind: kind,
	}

	if name != "" {
		p.Name = name
	}

	if typ != nil {
		p.Schema = schema.SchemaFor(typ)
		p.ContentType = contentType
	}

	return p
}
