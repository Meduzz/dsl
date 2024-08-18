package policy_test

import (
	"encoding/json"
	"testing"

	"github.com/Meduzz/dsl/policy"
)

func TestPolicy(t *testing.T) {
	t.Run("dev look and feel", func(t *testing.T) {
		p := policy.NewPolicy()

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

		bs, _ := json.Marshal(p)

		println(string(bs))

		result := &policy.Policy{}
		err := json.Unmarshal(bs, result)

		if err != nil {
			t.Error(err)
		}

		t.Run("figure out stuff", func(t *testing.T) {
			t.Run("as a document owner, what can I else do?", func(t *testing.T) {
				for _, it := range documentOwner.Inherits {
					println("I can also", it.Relation)
				}
			})
		})
	})
}
