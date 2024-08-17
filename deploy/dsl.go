package deploy

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

func NewPortMap(protocol string, container, host int) *PortMap {
	p := &PortMap{}

	p.Protocol = protocol
	p.Container = container
	p.Host = host

	return p
}

func NewVolume(host, container string) *Volume {
	v := &Volume{}

	v.Container = container
	v.Host = host

	return v
}

func NewConfigData(name, value string, kind ConfigKind) *ConfigData {
	d := &ConfigData{}

	d.Kind = kind
	d.Name = name
	d.Value = value

	return d
}
