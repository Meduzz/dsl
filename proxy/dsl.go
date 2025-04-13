package proxy

func ProxyConfig(domain, context string) *Proxy {
	return &Proxy{
		Domain:  domain,
		Context: context,
	}
}

func (p *Proxy) WithMiddlewares(middlewares ...*Middleware) *Proxy {
	p.Middlewares = append(p.Middlewares, middlewares...)
	return p
}

func (p *Proxy) Annotations(annotations ...string) *Proxy {
	p.Tags = append(p.Tags, annotations...)
	return p
}

func StripPrefix(name, prefix string) *Middleware {
	return &Middleware{
		Kind: "path:strip",
		Name: name,
		Metadata: map[string]any{
			"prefix": prefix,
		},
	}
}

func AddPrefix(name, prefix string) *Middleware {
	return &Middleware{
		Kind: "path:add",
		Name: name,
		Metadata: map[string]any{
			"prefix": prefix,
		},
	}
}

func ForwardAuth(name, url string) *Middleware {
	return &Middleware{
		Kind: "auth:forward",
		Name: name,
		Metadata: map[string]any{
			"url": url,
		},
	}
}

func ReplacePath(name, path string) *Middleware {
	return &Middleware{
		Kind: "path:replace",
		Name: name,
		Metadata: map[string]any{
			"path": path,
		},
	}
}

func Error(name, service, query string, statuses ...int) *Middleware {
	return &Middleware{
		Kind: "error",
		Name: name,
		Metadata: map[string]any{
			"service":  service,
			"query":    query,
			"statuses": statuses,
		},
	}
}

func Chain(name string, middlewares ...string) *Middleware {
	return &Middleware{
		Kind: "chain",
		Name: name,
		Metadata: map[string]any{
			"middlewares": middlewares,
		},
	}
}
