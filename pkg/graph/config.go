package graph

// GraphConfig represents the graph configuration
type GraphConfig struct {
	Nodes []NodeConfig `json:"nodes"`
	Edges []EdgeConfig `json:"edges"`
}

// NodeConfig represents a node configuration
type NodeConfig struct {
	ID       string                 `json:"id"`
	Settings map[string]interface{} `json:"settings"`
	Inputs   []interface{}          `json:"inputs,omitempty"`
}

// EdgeConfig represents an edge configuration
type EdgeConfig struct {
	Source string `json:"source"`
	Target string `json:"target"`
}
