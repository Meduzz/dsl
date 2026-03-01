package app_test

import (
	"encoding/json"
	"testing"

	"github.com/Meduzz/dsl/api/qapi"
	"github.com/Meduzz/dsl/app"
	"github.com/Meduzz/dsl/endpoint"
	"github.com/Meduzz/dsl/policy"
	"github.com/Meduzz/dsl/service"
	"github.com/Meduzz/quickapi/model"
)

type (
	Document struct {
		Name    string `json:"name"`
		Content string `json:"content,omitempty"`
		Created int64  `json:"created,omitempty"`
		Updated int64  `json:"updated,omitempty"`
	}

	DocumentEvent struct {
		Document string `json:"document"`
	}

	Folder struct {
		Folder string `json:"folder"`
		Parent string `json:"parent,omitempty"`
	}
)

var (
	_ model.Entity = Folder{}
)

func TestApp(t *testing.T) {
	app := app.NewApp("sheets", func(ab app.AppBuilder) {
		ab.SetDescription("A very simple sheet app")
		ab.SetDomain("docs.example.com")

		ab.AddService("documents", func(sb service.ServiceBuilder) {
			sb.AddEndpoint("listDocuments", func(eb endpoint.EndpointBuilder) {
				eb.GET("/")
				eb.QueryMap("where")
				eb.QueryMap("sort")
				eb.Query("skip")
				eb.Query("take")
				eb.SetDescription("List documents")
			})
			sb.AddEndpoint("createDocument", func(eb endpoint.EndpointBuilder) {
				eb.POST("/")
				eb.SetDescription("Create document")
			})
			sb.AddEndpoint("loadDocument", func(eb endpoint.EndpointBuilder) {
				eb.GET("/:id")
				eb.Path("id")
				eb.SetDescription("Load a single doc by id")
			})

			sb.AddEvent("document.created", &DocumentEvent{})
		})

		ab.AddService("folders", func(sb service.ServiceBuilder) {
			qapi.Quickapi(sb, "/api", "folder", Folder{})
		})

		ab.Policy(func(p policy.PolicyBuilder) {

			// define our relations
			owns := p.Relation("owner")
			edits := p.Relation("edit")
			views := p.Relation("view")
			parents := p.Relation("parent")

			// define our namespaces
			document := p.Namespace("document")
			folder := p.Namespace("folder")
			user := p.Namespace("user")

			// user relations
			p.Relationship(views, user.Subject(), document.Subject())
			p.Relationship(edits, user.Subject(), document.Subject())
			p.Relationship(owns, user.Subject(), document.Subject())
			p.Relationship(views, p.SubjectSet(document, edits), document.Subject())
			p.Relationship(edits, p.SubjectSet(document, owns), document.Subject())

			// folder relations
			p.Relationship(views, user.Subject(), folder.Subject())
			p.Relationship(edits, user.Subject(), folder.Subject())
			p.Relationship(owns, user.Subject(), folder.Subject())
			p.Relationship(views, p.SubjectSet(folder, edits), folder.Subject())
			p.Relationship(edits, p.SubjectSet(folder, owns), folder.Subject())

			// define folder relations
			p.Relationship(parents, folder.Subject(), folder.Subject())
			p.Relationship(parents, folder.Subject(), document.Subject())

			p.Relationship(views, p.SubjectSet(folder, views), document.Subject())
			p.Relationship(edits, p.SubjectSet(folder, edits), document.Subject())
			p.Relationship(owns, p.SubjectSet(folder, owns), document.Subject())
		})
	})

	bs, _ := json.Marshal(app)

	println(string(bs))
}

func (Folder) Create() any {
	return &Folder{}
}

func (Folder) CreateArray() any {
	return make([]*Folder, 0)
}

func (Folder) Name() string {
	return "folders"
}
