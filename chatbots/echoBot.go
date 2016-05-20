package bots

type EchoBot struct {
	initialMessage string
	knownUsers     map[string]bool
}

func NewEchoBot() *EchoBot {
	return &EchoBot{
		"Hallo Danke für die Kontaktaufnahme mit uns, wie können wir Ihnen helfen?",
		make(map[string]bool)}
}

func (e EchoBot) ReceivedMessage(from string, message string) (error, string) {
	msg := message
	if e.knownUsers[from] == false {
		msg = e.initialMessage
		e.knownUsers[from] = true
	}

	return nil, msg
}
