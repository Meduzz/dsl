package policy

/*
 * TODO
 * - (Subject, Relation & Namespace) Invent better names. I would like to use this for permission mapping too.
 * - (Policy) It's not really anything about policy any more.
 * - (Relationship) This might be the most poorly named one of the bunch.
 */
type (
	Relation  string // read
	Namespace string // document
	Subject   string

	Policy struct {
		Relations []*Relationship `json:"relations"`
	}

	Relationship struct {
		From     Subject  `json:"from"`
		Relation Relation `json:"relation"`
		To       Subject  `json:"to"`
	}

	PolicyBuilder interface {
		Namespace(string) Namespace
		Relation(string) Relation
		SubjectSet(ns Namespace, relation Relation) Subject
		Relationship(relation Relation, start, end Subject)
	}

	policyBuilder struct {
		policy *Policy
	}
)

var (
	_ PolicyBuilder = &policyBuilder{}
)
