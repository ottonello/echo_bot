package swagger

type Message struct {
	To   []string `json:"to,omitempty"`
	Body string   `json:"body,omitempty"`
}
