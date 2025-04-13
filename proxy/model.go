package proxy

type (
	Proxy struct {
		Domain      string        `json:"domain"`
		Context     string        `json:"context"`
		Middlewares []*Middleware `json:"middlewares,omitempty"`
		Tags        []string      `json:"tags,omitempty"`
	}

	Middleware struct {
		Kind     string         `json:"kind"`
		Name     string         `json:"name"`
		Metadata map[string]any `json:"metadata"`
	}
)
