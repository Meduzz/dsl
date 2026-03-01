package serviceref_test

import (
	"testing"

	. "github.com/Meduzz/dsl/serviceRef"
)

func TestServiceRef(t *testing.T) {
	invalid, err := Parse("invalid")

	if err == nil {
		t.Error("Expected an error")
	}

	if invalid != "" {
		t.Errorf("invalid serviceRef was not blank but: %s", invalid)
	}

	if invalid.Valid() {
		t.Error("invalid reported it self as valid")
	}

	valid, err := Parse("im/valid")

	if err != nil {
		t.Errorf("did not expect an error: %v", err)
	}

	if !valid.Valid() {
		t.Error("valid did not report it self as valid")
	}

	app, ok := valid.App()

	if !ok {
		t.Error("could not extract app from valid ServiceRef")
	}

	if app != "im" {
		t.Errorf("app was not im but: %s", app)
	}

	service, ok := valid.Service()

	if !ok {
		t.Error("could not extract service from valid ServiceRef")
	}

	if service != "valid" {
		t.Errorf("service was not valid but: %s", service)
	}

	invalid = NewServiceRef("", "")

	if invalid != "" {
		t.Errorf("invalid serviceRef was not blank but: %s", invalid)
	}

	if invalid.Valid() {
		t.Error("invalid reported it self as valid")
	}

	valid = NewServiceRef("im", "valid")

	if !valid.Valid() {
		t.Error("valid did not report it self as valid")
	}

	app, ok = valid.App()

	if !ok {
		t.Error("could not extract app from valid ServiceRef")
	}

	if app != "im" {
		t.Errorf("app was not im but: %s", app)
	}

	service, ok = valid.Service()

	if !ok {
		t.Error("could not extract service from valid ServiceRef")
	}

	if service != "valid" {
		t.Errorf("service was not valid but: %s", service)
	}
}
