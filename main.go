package main

import (
	"fmt"
	"github.com/slack-go/slack"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/commands/deploy", deployCommandHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func deployCommandHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: check it was came from slack.
	payload, err := slack.SlashCommandParse(r)
	if err != nil {
		fmt.Printf("Could not parse action response JSON: %v", err)
	}
	fmt.Printf("Comand came = %s\n", payload.Command)
	fmt.Printf("Also text = %s\n", payload.Text)
}
