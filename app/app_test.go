package app_test

import (
	"encoding/json"
	"testing"

	"github.com/Meduzz/dsl/app"
	"github.com/Meduzz/dsl/service"
)

type (
	Document struct {
		Name    string `json:"name"`
		Content string `json:"content,omitempty"`
		Created int64  `json:"created,omitempty"`
		Updated int64  `json:"updated,omitempty"`
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
	documentViewer := p.Relation(views, user.Subject(), document.Subject())
	documentEditor := p.Relation(edits, user.Subject(), document.Subject())
	documentEditor.Inherit(documentViewer)
	documentOwner := p.Relation(owns, user.Subject(), document.Subject())
	documentOwner.Inherit(documentEditor, documentViewer)

	folderViewer := p.Relation(views, user.Subject(), document.Subject())
	folderEditor := p.Relation(edits, user.Subject(), folder.Subject())
	folderEditor.Inherit(folderViewer)
	folderOwner := p.Relation(owns, user.Subject(), folder.Subject())
	folderOwner.Inherit(folderEditor, folderViewer)

	// define folder relations
	folderParent := p.Relation(parents, folder.Subject(), folder.Subject())
	folderParent.Inherit(folderOwner, folderEditor, folderViewer)

	folderDocumentViewer := p.Relation(views, folder.Subject(), document.Subject())
	folderDocumentEditor := p.Relation(edits, folder.Subject(), document.Subject())
	folderDocumentEditor.Inherit(folderDocumentViewer)
	folderDocumentOwner := p.Relation(owns, folder.Subject(), document.Subject())
	folderDocumentOwner.Inherit(folderDocumentEditor, folderDocumentViewer)

	bs, _ := json.Marshal(app)

	println(string(bs))

	t.Run("figure out stuff", func(t *testing.T) {
		t.Run("as a document owner, what can I else do?", func(t *testing.T) {
			for _, it := range documentOwner.Inherits {
				println("I can also", it.Relation)
			}
		})
	})
}
