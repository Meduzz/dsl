package policy

/*
 * TODO
 * - (Subject, Relationship & Namespace) Invent better names. I would like to use this for permission mapping too.
 * - (Policy) It's not really anything about policy any more.
 * - (Relation) This might be the most poorly named one of the bunch.
 */
type (
	Subject      string
	Relationship string
	Namespace    string

	Policy struct {
		Relations []*Relation `json:"relations"`
	}

	Relation struct {
		From     Subject      `json:"from"`
		Relation Relationship `json:"relation"`
		To       Subject      `json:"to"`
	}
)
