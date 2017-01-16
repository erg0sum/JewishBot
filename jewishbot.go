package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"golang.org/x/oauth2"
)

func authenticateWithReddit() {
	ctx := context.Background()
	conf := &oauth2.Config{
		ClientID:     "",
		ClientSecret: "1nEYXm0wWMPnJE0wdM9crYwStug",
		Scopes:       []string{"privatemessages", "read"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://provider.com/o/oauth2/auth",
			TokenURL: "https://provider.com/o/oauth2/token",
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
	client.Get("...")
}
func main() {
	if glossary, err := ReadGlossary(os.Args[1]); err != nil {
		fmt.Printf("Error reading glossary: %s\n", err.Error())
	} else {
		for key, value := range glossary {
			fmt.Printf("%s:%s\n", key, value)
		}
	}
}
