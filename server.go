package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-martini/martini"
	"github.com/ngti/echo_bot/swagger"
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

func handleMessage(w http.ResponseWriter, r *http.Request, api *swagger.DefaultApi, apiKey *string) {
	decoder := json.NewDecoder(r.Body)
	var t IncomingMessage
	err := decoder.Decode(&t)

	if err != nil {
		log.Print("Error decoding message:", r.Body)
	}
	log.Print("Got message: ", t)

	res, err := api.ApiV1MessagesPost(*apiKey, swagger.Message{
		Body: t.Body,
		To:   []string{t.From},
	})
	if err != nil {
		log.Print("Error sending message request ", err)
	} else {
		log.Print("Sent message id:", res)
	}

}

// NewServer creates a new instance that will use the given dependencies.
func NewServer(apiKey *string, api *swagger.DefaultApi) *Server {
	m := martini.Classic()

	m.Map(api)
	m.Map(apiKey)

	m.Post("/message", handleMessage)

	return &Server{m}
}

// Run starts listening to HTTP requests to port 3000 or to the port defined in the environment variable PORT.
func (server *Server) Run() {
	server.m.Run()
}
