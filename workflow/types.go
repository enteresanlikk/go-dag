package workflow

type Payload struct {
	Nodes []Node `json:"nodes"`
	Edges []Edge `json:"edges"`
}

type Node struct {
	ID       string                 `json:"id"`
	Data     []interface{}          `json:"data"`
	Settings map[string]interface{} `json:"settings"`
}

type Edge struct {
	Source string `json:"source"`
	Target string `json:"target"`
}
