package policy

import "github.com/Meduzz/dsl/api"

type (
	Subject      string
	Relationship string
	Namespace    string

	Policy struct {
		Relationships []Relationship `json:"relationships"`
		Relations     []*Relation    `json:"relations"`
		Namespaces    []Namespace    `json:"namespaces"`
		Rules         []*Rule        `json:"rules,omitempty"`
	}

	Relation struct {
		From     Subject      `json:"from"`
		Relation Relationship `json:"relation"`
		To       Subject      `json:"to"`
	}

	Rule struct {
		Name       string     `json:"rule"`
		Conditions []string   `json:"conditons"`
		Target     *api.Param `json:"target,omitempty"`
		Extends    []string   `json:"inherits,omitempty"`
	}
)
