package chatbots

type EchoBot struct {
	initialMessage string
	knownUsers     map[string]bool
}

func NewEchoBot(initialMessage string) *EchoBot {
	return &EchoBot{initialMessage, make(map[string]bool)}
}

func (e EchoBot) ReceivedMessage(from string, message string) (error, string) {
	msg := message
	if e.knownUsers[from] == false {
		msg = e.initialMessage
		e.knownUsers[from] = true
	}

	return nil, msg
}
