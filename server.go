package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ngti/echo_bot/swagger"

	"github.com/go-martini/martini"
)

// {
// "id": "9e5cf2779ffd42f88803477b734eb415"
// "name": "echo bot"
// "picture": "http://www.linnrecords.com/img/gallery/6281_LR_Echo_logo_BLACK%20small.jpg"
// "owner": "ngti"
// "webhooks": {
// "message": "http://localhost:3000/message"
// }-
// }

const (
	extensionId = "9e5cf2779ffd42f88803477b734eb415"
)

type IncomingMessage struct {
	From string `json:"from,omitempty"`
	Body string `json:"body,omitempty"`
	Type string `json:"type,omitempty"`
	Ts   int64  `json:"ts,omitempty"`
}

// Server is the main PBMock application that can be started to listen to HTTP requests.
type Server struct {
	m *martini.ClassicMartini
}

func handleMessage(w http.ResponseWriter, r *http.Request, api *swagger.DefaultApi) {
	decoder := json.NewDecoder(r.Body)
	var t IncomingMessage
	err := decoder.Decode(&t)

	if err != nil {
		log.Print("Error decoding message:", r.Body)
	}
	log.Print("Got message: ", t)

	res, err := api.ApiV1MessagesPost(swagger.Message{
		ExtensionId: extensionId,
		Body:        t.Body,
		To:          []string{t.From},
		Type_:       "chat",
	})
	// if err != nil {
	log.Print("Error sending message request ", err)
	// } else {
	log.Print("Sent message", res)
	// }

}

// NewServer creates a new instance that will use the given dependencies.
func NewServer() *Server {
	m := martini.Classic()
	api := swagger.NewDefaultApi()

	m.Map(api)

	m.Post("/message", handleMessage)

	return &Server{m}
}

// Run starts listening to HTTP requests to port 3000 or to the port defined in the environment variable PORT.
func (server *Server) Run() {
	server.m.Run()
}
