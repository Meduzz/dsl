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

func (p *Policy) Relationship(name string) Relationship {
	r := Relationship(name)
	p.Relationships = append(p.Relationships, r)
	return r
}

func (p *Policy) Namespace(name string) Namespace {
	n := NewNamespace(name)
	p.Namespaces = append(p.Namespaces, n)
	return n
}

func NewNamespace(name string) Namespace {
	return Namespace(name)
}

func (n Namespace) Subject() Subject {
	return Subject(n)
}

func SubjectSet(namespace Namespace, relation Relationship) Subject {
	return Subject(fmt.Sprintf("%s#%s", namespace.Subject(), relation))
}

func (p *Policy) Rule(name string) *Rule {
	r := &Rule{}

	r.Name = name
	p.Rules = append(p.Rules, r)

	return r
}

func (r *Rule) Condition(conditions ...string) {
	r.Conditions = conditions
}

func (r *Rule) And(condition string) {
	r.Conditions = append(r.Conditions, condition)
}

func (r *Rule) Inherits(rules ...string) {
	r.Extends = append(r.Extends, rules...)
}
