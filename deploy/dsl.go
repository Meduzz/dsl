package deploy

func DeployConfig(binary string) *Deploy {
	return &Deploy{
		Binary: binary,
	}
}

func Env(name string) *Property {
	return &Property{
		Name: name,
		Kind: "env",
	}
}

func Arg(name string) *Property {
	return &Property{
		Name: name,
		Kind: "arg",
	}
}

func WithDB(dialect *Dialect, strategy *Property) *Option {
	return &Option{
		Kind:     "dependency",
		Metadata: map[string]any{"kind": "db", "strategy": strategy, "dialect": dialect},
	}
}

func WithDialect(dialect, dbName string, dsn bool) *Dialect {
	return &Dialect{
		Name: dbName,
		Kind: dialect,
		DSN:  dsn,
	}
}

func WithRedis(strategy *Property) *Option {
	return &Option{
		Kind:     "dependency",
		Metadata: map[string]any{"kind": "redis", "strategy": strategy},
	}
}

func WithNats(strategy *Property) *Option {
	return &Option{
		Kind:     "dependency",
		Metadata: map[string]any{"kind": "nats", "strategy": strategy},
	}
}

func WithSecret(strategy *Property) *Option {
	return &Option{
		Kind:     "dependency",
		Metadata: map[string]any{"kind": "secret", "strategy": strategy},
	}
}

func WithConfig(value string, strategy *Property) *Option {
	return &Option{
		Kind:     "dependency",
		Metadata: map[string]any{"value": value, "strategy": strategy},
	}
}

func WithTcpPort(port int, name string) *Option {
	return &Option{
		Kind:     "port",
		Metadata: map[string]any{"name": name, "port": port, "protocol": "tcp"},
	}
}

func WithUdpPort(port int, name string) *Option {
	return &Option{
		Kind:     "port",
		Metadata: map[string]any{"name": name, "port": port, "protocol": "udp"},
	}
}

func WithVolume(name, containerPath string) *Option {
	return &Option{
		Kind:     "volume",
		Metadata: map[string]any{"path": containerPath, "name": name},
	}
}

func WithCommand(cmd string) *Option {
	return &Option{
		Kind:     "command",
		Metadata: map[string]any{"kind": "cmd", "cmd": cmd},
	}
}

func WithArgument(name, value string) *Option {
	return &Option{
		Kind:     "command",
		Metadata: map[string]any{"kind": "arg", "value": value, "strategy": Arg(name)},
	}
}

func (d *Deploy) WithOptions(options ...*Option) *Deploy {
	d.Options = append(d.Options, options...)
	return d
}
