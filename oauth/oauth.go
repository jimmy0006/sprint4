package oauth

import (
    "context"
    "golang.org/x/oauth2"
    "golang.org/x/oauth2/google"
    "google.golang.org/api/urlshortener/v1"
)

svc, err:=urlshortener.New(httpClient)

var config = &oauth2.Config{
	ClientID:     "", // from https://console.developers.google.com/project/<your-project-id>/apiui/credential
	ClientSecret: "", // from https://console.developers.google.com/project/<your-project-id>/apiui/credential
	Endpoint:     google.Endpoint,
	Scopes:       []string{urlshortener.UrlshortenerScope},
}