package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/go-martini/martini"
	"github.com/ngti/echo_bot/bots"
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

type config struct {
	apiKey string
	bot    bots.ChatBot
}

func NewServer(apiKey string, api *swagger.DefaultApi) *Server {
	m := martini.Classic()

	bot := bots.NewEchoBot("Hi!")
	conf := &config{apiKey, *bot}
	m.Map(api)
	m.Map(conf)

	m.Post("/message", handleMessage)

	return &Server{m}
}

// Run starts listening to HTTP requests to port 3000 or to the port defined in the environment variable PORT.
func (server *Server) Run() {
	server.m.Run()
}

func handleMessage(w http.ResponseWriter, r *http.Request, api *swagger.DefaultApi, config *config) {
	err, jsonMsg := decode(r.Body)

	if err != nil {
		log.Panic("Error decoding message: ", r.Body, err)
	}
	log.Print("Got message: ", jsonMsg)

	err, response := config.bot.ReceivedMessage(jsonMsg.From, jsonMsg.Body)

	if err != nil {
		log.Panic("Error processing message: ", r.Body, err)
	}
	if response != "" {
		reply(jsonMsg.From, response, *api, *config)
	}
}

func reply(to string, msg string, api swagger.DefaultApi, config config) {
	log.Print("Replying: ", msg)

	res, err := api.ApiV1MessagesPost(config.apiKey, swagger.Message{
		Body: msg,
		To:   []string{to},
	})

	if err != nil {
		log.Print("Error sending message request ", err)
	} else {
		log.Print("Sent message id:", res)
	}
}

func decode(body io.Reader) (error, IncomingMessage) {
	decoder := json.NewDecoder(body)
	var jsonMsg IncomingMessage
	err := decoder.Decode(&jsonMsg)

	return err, jsonMsg
}
