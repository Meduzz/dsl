package service

import "reflect"

func NewService(name string, kind ServiceKind) *Service {
	s := &Service{}

	s.Name = name
	s.Kind = kind

	return s
}

func (s *Service) AddEndpoint(endpoint *Endpoint) *Endpoint {
	s.Endpoints = append(s.Endpoints, endpoint)
	return endpoint
}

func (s *Service) AddPort(port *Port) *Port {
	s.Ports = append(s.Ports, port)
	return port
}

func (s *Service) AddParam(param *Config) *Config {
	s.Params = append(s.Params, param)
	return param
}

func (s *Service) SetDeploy(deploy *Deploy) *Deploy {
	s.Deploy = deploy

	return deploy
}

func GET(path string) *Endpoint {
	e := &Endpoint{}

	e.Path = path
	e.Method = "GET"

	return e
}

func POST(path string) *Endpoint {
	e := &Endpoint{}

	e.Path = path
	e.Method = "POST"

	return e
}

func PUT(path string) *Endpoint {
	e := &Endpoint{}

	e.Path = path
	e.Method = "PUT"

	return e
}

func DELETE(path string) *Endpoint {
	e := &Endpoint{}

	e.Path = path
	e.Method = "DELETE"

	return e
}

func PATCH(path string) *Endpoint {
	e := &Endpoint{}

	e.Path = path
	e.Method = "PATCH"

	return e
}

func OPTION(path string) *Endpoint {
	e := &Endpoint{}

	e.Path = path
	e.Method = "OPTION"

	return e
}

func (e *Endpoint) AddArgument(param *Param) *Param {
	e.Arguments = append(e.Arguments, param)
	return param
}

func TCP(port int) *Port {
	p := &Port{}

	p.Port = port
	p.Protocol = "tcp"

	return p
}

func UDP(port int) *Port {
	p := &Port{}

	p.Port = port
	p.Protocol = "udp"

	return p
}

func Argv(name string) *Config {
	p := &Config{}

	p.Name = name
	p.Kind = Argument

	return p
}

func Env(name string) *Config {
	p := &Config{}

	p.Name = name
	p.Kind = Environment

	return p
}

func PathVariable(name string) *Param {
	p := &Param{}

	p.Name = name
	p.Kind = PathKind

	return p
}

func QueryVariable(name string) *Param {
	p := &Param{}

	p.Name = name
	p.Kind = QueryKind

	return p
}

func BodyVariable(name, contentType string) *Param {
	p := &Param{}

	p.Name = name
	p.Kind = BodyKind
	p.Format = contentType

	return p
}

func HeaderVariable(name string) *Param {
	p := &Param{}

	p.Name = name
	p.Kind = HeaderKind

	return p
}

func (p *Param) SetType(it any) {
	v := reflect.ValueOf(it)

	if v.Kind() == reflect.Pointer {
		println("it's a pointer")
		p.Pointer = true
		v = v.Elem() // drop pointer
	}

	t := v.Type()
	p.Type = t.String()
}

func (p *Param) ArrayOf(it any) {
	p.Array = true
	p.SetType(it)
}

func (p *Param) MapOf(it any) {
	p.Map = true
	p.SetType(it)
}

func NewDeploy(image, command string, network ...string) *Deploy {
	d := &Deploy{}

	d.Image = image
	d.Command = command
	d.Networks = network

	return d
}

func (d *Deploy) AddPortMap(portMap *PortMap) *PortMap {
	d.PortMaps = append(d.PortMaps, portMap)

	return portMap
}

func (d *Deploy) AddVolume(volume *Volume) *Volume {
	d.Volumes = append(d.Volumes, volume)

	return volume
}

func (d *Deploy) AddConfigData(data *ConfigData) *ConfigData {
	d.ConfigData = append(d.ConfigData, data)
	return data
}

func NewPortMap(from *Port, host int) *PortMap {
	p := &PortMap{}

	p.Protocol = from.Protocol
	p.Container = from.Port
	p.Host = host

	return p
}

func NewVolume(host, container string) *Volume {
	v := &Volume{}

	v.Container = container
	v.Host = host

	return v
}

func NewConfigData(config *Config, value string) *ConfigData {
	d := &ConfigData{}

	d.Kind = config.Kind
	d.Name = config.Name
	d.Value = value

	return d
}
