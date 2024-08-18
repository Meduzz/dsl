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
	s1.TCP(8080)
	s1pa1 := s1.Env("DB_URL")
	s1pa1.Description = "The DSN to connect to the DB."
	s1api := s1.API()
	root := s1api.GET("/")
	root.Name = "root"
	root.Description = "The root of the app, the first thing the visitor sees"
	rootRes := root.SetResponse("text/html")
	rootRes.SetType("")
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
