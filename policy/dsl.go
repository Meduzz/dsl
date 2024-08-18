package policy

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

func (p *Policy) Namespace(name string) *Namespace {
	n := NewNamespace(name)
	p.Namespaces = append(p.Namespaces, n)
	return n
}

func NewNamespace(name string) *Namespace {
	n := &Namespace{}
	n.Name = name
	return n
}

func (n *Namespace) Subject() Subject {
	return Subject(n.Name)
}

func (s *Relation) Inherit(other ...*Relation) {
	s.Inherits = append(s.Inherits, other...)
}
