package service_test

import (
	"testing"

	"github.com/Meduzz/dsl/endpoint"
	"github.com/Meduzz/dsl/service"
)

func TestService(t *testing.T) {
	t.Run("all appenders are appending", func(t *testing.T) {
		s := service.NewService("test", func(sb service.ServiceBuilder) {
			sb.AddDependency("test", "test")
			sb.AddTags("a=1")
			sb.Container("test")
			sb.SetDescription("Very descriptive")
			sb.AddEndpoint("test", func(eb endpoint.EndpointBuilder) {
				eb.GET("/test")
			})
		})

		if len(s.Dependency) == 0 {
			t.Errorf("no dependency was created")
		}

		if s.Description == "" {
			t.Errorf("description was not set")
		}

		if len(s.Tags) == 0 {
			t.Errorf("no tags was created")
		}

		if s.Packaging == nil {
			t.Errorf("packaging was not set")
		}

		if len(s.Endpoints) == 0 {
			t.Errorf("no endpoint was created")
		}
	})
}
