package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	r.POST("/", HelloPubSub)
	log.Println("listening to port *:8080. press ctrl + c to cancel.")
	log.Fatal(http.ListenAndServe(":8080", r))
}

type PubSubMessage struct {
	Message struct {
		Data []byte `json:"data,omitempty"`
		ID   string `json:"id"`
	} `json:"message"`
	Subscription string `json:"subscription"`
}

func HelloPubSub(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Parse the Pub/Sub message.
	var m PubSubMessage

	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		log.Printf("json.NewDecoder: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	name := string(m.Message.Data)
	if name == "" {
		name = "World"
	}
	log.Printf("Hello %s!", name)
}
