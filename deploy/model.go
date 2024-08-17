package deploy

type (
	ConfigKind string

	Deploy struct {
		Image      string        `json:"image"`
		Command    string        `json:"command"`
		PortMaps   []*PortMap    `json:"portMap,omitempty"`
		Volumes    []*Volume     `json:"volumes,omitempty"`
		Networks   []string      `json:"networks,omitempty"`
		ConfigData []*ConfigData `json:"configData,omitempty"`
	}

	PortMap struct {
		Protocol  string `json:"protocol"`
		Container int    `json:"container"`
		Host      int    `json:"host"`
	}

	Volume struct {
		Container string `json:"container"`
		Host      string `json:"host"`
	}

	ConfigData struct {
		Name  string     `json:"name"`
		Kind  ConfigKind `json:"kind"` // arg|env
		Value string     `json:"value"`
	}
)

var (
	Argument    = ConfigKind("arg")
	Environment = ConfigKind("env")
)
