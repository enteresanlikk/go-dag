package graph

type GraphConfig struct {
	Nodes []NodeConfig `json:"nodes"`
	Edges []EdgeConfig `json:"edges"`
}

type NodeConfig struct {
	ID       string                 `json:"id"`
	Settings map[string]interface{} `json:"settings"`
	Inputs   []interface{}          `json:"inputs,omitempty"`
}

type EdgeConfig struct {
	Source string `json:"source"`
	Target string `json:"target"`
}
