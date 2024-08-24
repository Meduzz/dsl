package app_test

import (
	"encoding/json"
	"testing"

	"github.com/Meduzz/dsl/app"
	"github.com/Meduzz/dsl/policy"
	"github.com/Meduzz/dsl/service"
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
)

var (
	MyKind = service.ServiceKind("MyKind")
)

func TestApp(t *testing.T) {
	app := app.NewApp("sheets")
	app.Description = "A very simple sheet app"

	documentsService := app.AddService("documents", MyKind)
	documentsService.TCP(8080)
	dbUrl := documentsService.Env("DB_URL")
	dbUrl.Description = "The DSN to connect to the DB."

	documentsApi := documentsService.API()
	listDocuments := documentsApi.GET("/")
	listDocuments.Description = "List documents"
	listDocuments.QueryVariable("skip")
	listDocuments.QueryVariable("take")
	listDocsResp := listDocuments.SetResponse("text/html")
	listDocsResp.ArrayOf(&Document{})

	createDocument := documentsApi.POST("/")
	createDocument.Description = "Crate a document"
	createDocPayload := createDocument.BodyVariable("body", "application/json")
	createDocPayload.SetType(&Document{})
	createDocResp := createDocument.SetResponse("application/json")
	createDocResp.SetType(&Document{})

	fetchDocument := documentsApi.GET("/:id")
	fetchDocument.Description = "Fetch a document by id"
	fetchDocument.PathVariable("id")
	fetchDocResp := fetchDocument.SetResponse("application/json")
	fetchDocResp.SetType(&Document{})

	topic := documentsApi.Event("document.created")
	body := topic.Event("application/json")
	body.SetType(&DocumentEvent{})

	folderService := app.AddService("folders", MyKind)
	folderService.TCP(8080)
	dbConn := folderService.Env("DB_URL")
	dbConn.Description = "The DSN to connect to the DB."

	p := app.GetPolicy()

	// define our relations
	owns := p.Relationship("owner")
	edits := p.Relationship("edit")
	views := p.Relationship("view")
	parents := p.Relationship("parent")

	// define our namespaces
	document := p.Namespace("document")
	folder := p.Namespace("folder")
	user := p.Namespace("user")

	// user relations
	p.Relation(views, user.Subject(), document.Subject())
	p.Relation(edits, user.Subject(), document.Subject())
	p.Relation(owns, user.Subject(), document.Subject())
	p.Relation(views, policy.SubjectSet(document, edits), document.Subject())
	p.Relation(edits, policy.SubjectSet(document, owns), document.Subject())

	// folder relations
	p.Relation(views, user.Subject(), folder.Subject())
	p.Relation(edits, user.Subject(), folder.Subject())
	p.Relation(owns, user.Subject(), folder.Subject())
	p.Relation(views, policy.SubjectSet(folder, edits), folder.Subject())
	p.Relation(edits, policy.SubjectSet(folder, owns), folder.Subject())

	// define folder relations
	p.Relation(parents, folder.Subject(), folder.Subject())   // folder/folder
	p.Relation(parents, folder.Subject(), document.Subject()) // folder/document

	p.Relation(views, policy.SubjectSet(folder, views), document.Subject())
	p.Relation(edits, policy.SubjectSet(folder, edits), document.Subject())
	p.Relation(owns, policy.SubjectSet(folder, owns), document.Subject())

	bs, _ := json.Marshal(app)

	println(string(bs))
}
