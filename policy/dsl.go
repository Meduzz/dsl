package policy

import "fmt"

func (p *Policy) Relation(relationship Relationship, from, to Subject) *Relation {
	r := &Relation{}

	r.Relation = relationship
	r.From = from
	r.To = to

	p.Relations = append(p.Relations, r)
	return r
}

func (p *Policy) AddRelation(rel *Relation) *Policy {
	p.Relations = append(p.Relations, rel)

	return p
}

func (p *Policy) Relationship(name string) Relationship {
	return Relationship(name)
}

func (p *Policy) Namespace(name string) Namespace {
	return Namespace(name)
}

func (n Namespace) Subject() Subject {
	return Subject(n)
}

func SubjectSet(namespace Namespace, relation Relationship) Subject {
	return Subject(fmt.Sprintf("%s#%s", namespace.Subject(), relation))
}

func (r Relationship) Between(from, to Subject) *Relation {
	rel := &Relation{}

	rel.Relation = r
	rel.From = from
	rel.To = to

	return rel
}
