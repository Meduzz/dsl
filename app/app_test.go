package app_test

import (
	"encoding/json"
	"testing"

	"github.com/Meduzz/dsl/app"
	"github.com/Meduzz/dsl/service"
)

func TestApp(t *testing.T) {
	app := app.NewApp("test")
	app.Description = "A very simple test app"
	s1 := app.AddService(service.NewService("service1", service.Gin))
	s1.AddPort(service.TCP(8080))
	root := s1.AddEndpoint(service.GET("/"))
	root.Name = "root"
	root.Description = "The root of the app, the first thing the visitor sees"
	root.Returns = service.BodyVariable("body", "text/html")
	root.Returns.SetType("")
	s1pa1 := s1.AddParam(service.Env("DB_URL"))
	s1pa1.Description = "The DSN to connect to the DB."

	bs, _ := json.Marshal(app)

	println(string(bs))
}
