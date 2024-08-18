package api

import (
	"reflect"
	"strings"
)

func (a *Api) GET(path string) *Endpoint {
	e := &Endpoint{}

	e.Path = path
	e.Method = "GET"

	a.Endpoints = append(a.Endpoints, e)

	return e
}

func (a *Api) POST(path string) *Endpoint {
	e := &Endpoint{}

	e.Path = path
	e.Method = "POST"

	a.Endpoints = append(a.Endpoints, e)

	return e
}

func (a *Api) PUT(path string) *Endpoint {
	e := &Endpoint{}

	e.Path = path
	e.Method = "PUT"

	return e
}

func (a *Api) DELETE(path string) *Endpoint {
	e := &Endpoint{}

	e.Path = path
	e.Method = "DELETE"

	a.Endpoints = append(a.Endpoints, e)

	return e
}

func (a *Api) PATCH(path string) *Endpoint {
	e := &Endpoint{}

	e.Path = path
	e.Method = "PATCH"

	a.Endpoints = append(a.Endpoints, e)

	return e
}

func (a *Api) OPTION(path string) *Endpoint {
	e := &Endpoint{}

	e.Path = path
	e.Method = "OPTION"

	a.Endpoints = append(a.Endpoints, e)

	return e
}

func (e *Endpoint) PathVariable(name string) *Param {
	p := &Param{}

	p.Name = name
	p.Kind = PathKind

	e.Request = append(e.Request, p)

	return p
}

func (e *Endpoint) QueryVariable(name string) *Param {
	p := &Param{}

	p.Name = name
	p.Kind = QueryKind

	e.Request = append(e.Request, p)

	return p
}

func (e *Endpoint) BodyVariable(name, contentType string) *Param {
	p := &Param{}

	p.Name = name
	p.Kind = BodyKind
	p.Format = contentType

	e.Request = append(e.Request, p)

	return p
}

func (e *Endpoint) HeaderVariable(name string) *Param {
	p := &Param{}

	p.Name = name
	p.Kind = HeaderKind

	e.Request = append(e.Request, p)

	return p
}

func (e *Endpoint) SetResponse(contentType string) *Param {
	p := &Param{}

	p.Kind = BodyKind
	p.Format = contentType

	e.Response = p

	return p
}

func (p *Param) SetType(it any) {
	v := reflect.ValueOf(it)

	if v.Kind() == reflect.Pointer {
		p.Pointer = true
		v = v.Elem() // drop pointer
	}

	t := v.Type()
	p.Type = t.String()

	if v.Kind() == reflect.Struct {
		p.Payload = parseStruct(v)
	}
}

func (p *Param) ArrayOf(it any) {
	p.Array = true
	p.SetType(it)
}

func (p *Param) MapOf(it any) {
	p.Map = true
	p.SetType(it)
}

func parseStruct(v reflect.Value) *Payload {
	p := &Payload{}
	p.Name = v.Type().Name()

	fieldCount := v.NumField()

	for i := 0; i < fieldCount; i++ {
		rf := v.Type().Field(i)
		f := &Field{}

		f.Name = rf.Name
		f.Type = rf.Type.Name()
		jsonTag, ok := rf.Tag.Lookup("json")

		if ok {
			jsonTag = strings.Replace(jsonTag, "omitempty", "", -1)
			tagContent := strings.Split(jsonTag, ",")

			for _, it := range tagContent {
				if it != "" && it != "-" {
					f.Name = it
					break
				}
			}
		}

		p.Fields = append(p.Fields, f)
	}

	return p
}
