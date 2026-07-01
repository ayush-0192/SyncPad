package operation

type OperationType string

const (
	Insert OperationType = "insert"
	Delete OperationType = "delete"
)

type Operation struct {
	Type OperationType `json:"type"`
	Position int `json:"position"`
	Text string `json:"text,omitempty"`
	Length int `json:"length,omitempty"`
	Version int `json:version`
}