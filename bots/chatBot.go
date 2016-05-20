package chatbots

type ChatBot interface {
	ReceivedMessage(from string, message string) (err error, response string)
}
