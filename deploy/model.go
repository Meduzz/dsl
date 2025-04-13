package deploy

type (
	Deploy struct {
		Binary  string    `json:"binary"`
		Options []*Option `json:"options"`
	}

	Option struct {
		Kind     string         `json:"kind"`
		Metadata map[string]any `json:"metadata"`
	}

	Property struct {
		Name string `json:"name"`
		Kind string `json:"kind"` // arg|env
	}

	Dialect struct {
		Name string `json:"name"`
		Kind string `json:"kind"`
		DSN  bool   `json:"dsn"`
	}
)
