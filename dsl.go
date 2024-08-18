package dsl

import "github.com/Meduzz/dsl/app"

// Start here
func NewApp(name, description string) *app.App {
	a := app.NewApp(name)
	a.Description = description

	return a
}
