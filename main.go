package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/go-github/v29/github"
	"github.com/slack-go/slack"
	"golang.org/x/oauth2"
	"log"
	"net/http"
	"os"
)

var client *github.Client
var ctx = context.Background()

func main() {
	token := os.Getenv("GITHUB_AUTH_TOKEN")
	initGithubClient(token)

	bytes, _ := json.Marshal(struct { Id string } { Id: ""})
	payload := json.RawMessage(bytes)
	client.Repositories.Dispatch(ctx, "moreal", "slack-chatops-practice", github.DispatchRequestOptions{
		EventType: "go",
		ClientPayload: &payload,
	})

	http.HandleFunc("/commands/deploy", deployCommandHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func initGithubClient(token string) {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	client = github.NewClient(tc)
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
