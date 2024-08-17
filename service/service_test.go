package service_test

import (
	"testing"

	"github.com/Meduzz/dsl/service"
)

type TestStruct struct{}

func TestService(t *testing.T) {
	t.Run("all appenders are appending", func(t *testing.T) {
		s := service.NewService("test", service.Gin)

		e := s.AddEndpoint(service.GET("/test"))

		if len(s.Endpoints) != 1 {
			t.Error("endpoint was not appended")
		}

		s.AddParam(service.Argv("test"))

		if len(s.Params) != 1 {
			t.Error("param was not appended")
		}

		s.AddPort(service.TCP(8080))

		if len(s.Ports) != 1 {
			t.Error("port was not appended")
		}

		e.AddArgument(service.QueryVariable("test"))

		if len(e.Arguments) != 1 {
			t.Error("argument was not appended")
		}

		d := service.NewDeploy("test", "service")

		d.AddPortMap(service.NewPortMap(service.TCP(8080), 9000))

		if len(d.PortMaps) != 1 {
			t.Error("portMap was not appended")
		}

		d.AddVolume(service.NewVolume("/", "/"))

		if len(d.Volumes) != 1 {
			t.Error("volume was not appended")
		}
	})

	t.Run("Param.SetType sets the type", func(t *testing.T) {
		p := service.QueryVariable("test")
		p.SetType("a string")

		if p.Type != "string" {
			t.Errorf("type was %s", p.Type)
		}

		if p.Pointer || p.Array || p.Map {
			t.Error("pointer array map was set")
		}

		p = service.QueryVariable("test")
		p.SetType(42)

		if p.Type != "int" {
			t.Errorf("type was %s", p.Type)
		}

		if p.Pointer || p.Array || p.Map {
			t.Error("pointer array map was set")
		}

		p = service.QueryVariable("test")
		p.SetType(false)

		if p.Type != "bool" {
			t.Errorf("type was %s", p.Type)
		}

		if p.Pointer || p.Array || p.Map {
			t.Error("pointer array map was set")
		}

		p = service.QueryVariable("test")
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

		p = service.QueryVariable("test")
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

		p = service.QueryVariable("test")
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

	t.Run("portMaps are made of ports", func(t *testing.T) {
		port := service.TCP(8080)
		portMap := service.NewPortMap(port, 9000)

		if portMap.Protocol != port.Protocol {
			t.Error("protocol does not match")
		}

		if portMap.Container != port.Port {
			t.Error("container and port does not match")
		}

		if portMap.Host != 9000 {
			t.Error("host was not set correctly")
		}
	})
}
