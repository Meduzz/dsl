package serviceref

import (
	"fmt"
	"strings"
)

func NewServiceRef(app, service string) ServiceRef {
	if app == "" && service == "" {
		return ""
	}

	if app == "" && service != "" {
		app = service
	}

	if app != "" && service == "" {
		service = app
	}

	return ServiceRef(fmt.Sprintf("%s/%s", app, service))
}

func Parse(ref string) (ServiceRef, error) {
	splitterCount := strings.Count(ref, "/")

	if splitterCount == 0 || splitterCount > 1 {
		return "", fmt.Errorf("invalid serviceRef")
	}

	return ServiceRef(ref), nil
}

func (s ServiceRef) Valid() bool {
	return s != ""
}

func (s ServiceRef) split() []string {
	return strings.Split(string(s), "/")
}

func (s ServiceRef) App() (string, bool) {
	split := s.split()

	if len(split) == 2 {
		return split[0], true
	}

	return "", false
}

func (s ServiceRef) Service() (string, bool) {
	split := s.split()

	if len(split) == 2 {
		return split[1], true
	}

	return "", false
}
