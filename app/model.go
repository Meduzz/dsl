package app

import "github.com/Meduzz/dsl/service"

type (
	App struct {
		Name        string             `json:"name"`
		Description string             `json:"description,omitempty"`
		Services    []*service.Service `json:"services"`
	}
)
