package policy

import "fmt"

func (p *policyBuilder) Relationship(relation Relation, from, to Subject) {
	r := &Relationship{}

	r.Relation = relation
	r.From = from
	r.To = to

	p.policy.Relations = append(p.policy.Relations, r)
}

func (p *policyBuilder) Relation(name string) Relation {
	return Relation(name)
}

func (p *policyBuilder) Namespace(name string) Namespace {
	return Namespace(name)
}

func (n Namespace) Subject() Subject {
	return Subject(n)
}

func (p *policyBuilder) SubjectSet(namespace Namespace, relation Relation) Subject {
	return Subject(fmt.Sprintf("%s#%s", namespace.Subject(), relation))
}

func NewPolicyBuilder(policy *Policy) PolicyBuilder {
	return &policyBuilder{policy}
}
