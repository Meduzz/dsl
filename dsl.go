package dsl

import "github.com/Meduzz/dsl/app"

// Start here
func NewApp(name string, cb func(app.AppBuilder)) *app.App {
	a := app.NewApp(name, cb)

	return a
}
