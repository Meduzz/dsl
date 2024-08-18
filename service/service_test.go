package service_test

import (
	"testing"

	"github.com/Meduzz/dsl/api"
	"github.com/Meduzz/dsl/service"
)

type TestStruct struct {
	Name string `json:"name,omitempty"`
	Age  int
}

func TestService(t *testing.T) {
	t.Run("all appenders are appending", func(t *testing.T) {
		s := service.Service{}
		a := s.API()
		e := a.GET("/test")

		if len(a.Endpoints) != 1 {
			t.Error("endpoint was not appended")
		}

		s.Argv("test")

		if len(s.Params) != 1 {
			t.Error("param was not appended")
		}

		s.TCP(8080)

		if len(s.Ports) != 1 {
			t.Error("port was not appended")
		}

		e.QueryVariable("test")

		if len(e.Request) != 1 {
			t.Error("argument was not appended")
		}

		s.AddVolumes("/")

		if len(s.Volumes) != 1 {
			t.Error("volume was not appended")
		}
	})

	t.Run("Param.SetType sets the type", func(t *testing.T) {
		a := &api.Api{}
		e := a.GET("/")
		p := e.QueryVariable("test")
		p.SetType("a string")

		if p.Type != "string" {
			t.Errorf("type was %s", p.Type)
		}

		if p.Pointer || p.Array || p.Map {
			t.Error("pointer array map was set")
		}

		p = e.QueryVariable("test")
		p.SetType(42)

		if p.Type != "int" {
			t.Errorf("type was %s", p.Type)
		}

		if p.Pointer || p.Array || p.Map {
			t.Error("pointer array map was set")
		}

		p = e.QueryVariable("test")
		p.SetType(false)

		if p.Type != "bool" {
			t.Errorf("type was %s", p.Type)
		}

		if p.Pointer || p.Array || p.Map {
			t.Error("pointer array map was set")
		}

		p = e.QueryVariable("test")
		p.SetType(&TestStruct{})

		if p.Type != "service_test.TestStruct" {
			t.Errorf("type was %s", p.Type)
		}

		if p.Array || p.Map {
			t.Error("array map was set")
		}

		if !p.Pointer {
			t.Error("pointer was not set")
		}

		p = e.QueryVariable("test")
		p.ArrayOf(&TestStruct{})

		if p.Type != "service_test.TestStruct" {
			t.Errorf("type was %s", p.Type)
		}

		if p.Map {
			t.Error("map was set")
		}

		if !p.Pointer {
			t.Error("pointer was not set")
		}

		if !p.Array {
			t.Error("array was not set")
		}

		p = e.QueryVariable("test")
		p.MapOf(&TestStruct{})

		if p.Type != "service_test.TestStruct" {
			t.Errorf("type was %s", p.Type)
		}

		if !p.Map {
			t.Error("map was not set")
		}

		if !p.Pointer {
			t.Error("pointer was not set")
		}

		if p.Array {
			t.Error("array was set")
		}
	})
}
