package policy

type (
	Subject      string
	Relationship string

	Policy struct {
		Relationships []Relationship `json:"relationships"`
		Relations     []*Relation    `json:"relations"`
		Namespaces    []*Namespace   `json:"namespaces"`
	}

	Namespace struct {
		Name string `json:"name"`
	}

	Relation struct {
		From     Subject      `json:"from"`
		Relation Relationship `json:"relation"`
		To       Subject      `json:"to"`
		Inherits []*Relation  `json:"inherits,omitempty"` // This makes the document a lot larger and more complex
	}
)
