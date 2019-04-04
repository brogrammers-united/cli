package main

import (
	"context"
	"fmt"
	"log"

	"golang.org/x/oauth2"
)

func GetAccessToken() {
	ctx := context.Background()
	conf := &oauth2.Config{
		ClientID:     "bgu-cli-client",
		ClientSecret: "bgu-cli-secret",
		Scopes:       []string{"user", "repo"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "http://localhost:8080/oauth2/authorize/github-api",
			TokenURL: "http://localhost:8080/oauth/token",
		},
	}

	// Redirect user to consent page to ask for permission
	// for the scopes specified above.
	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	fmt.Printf("Visit the URL for the auth dialog: %v", url)

	// Use the authorization code that is pushed to the redirect
	// URL. Exchange will do the handshake to retrieve the
	// initial access token. The HTTP Client returned by
	// conf.Client will refresh the token as necessary.
	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatal(err)
	}
	tok, err := conf.Exchange(ctx, code)
	if err != nil {
		log.Fatal(err)
	}

	client := conf.Client(ctx, tok)
	client.Get("https://api.github.com/user?")
}
