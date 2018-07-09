// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bitbucket

import (
	"net/http"

	"github.com/drone/go-login/login"
	"github.com/drone/go-login/login/internal/oauth2"
)

const (
	accessTokenURL   = "https://bitbucket.org/site/oauth2/access_token"
	authorizationURL = "https://bitbucket.org/site/oauth2/authorize"
)

// Option configures an authorization handler option.
type Option func(a *Authorizer)

// WithClient configures the authorization handler with a
// custom http.Client.
func WithClient(client *http.Client) Option {
	return func(a *Authorizer) {
		a.client = client
	}
}

// WithClientID configures the authorization handler with
// the client_id.
func WithClientID(clientID string) Option {
	return func(a *Authorizer) {
		a.clientID = clientID
	}
}

// WithClientSecret configures the authorization handler
// with the client_secret.
func WithClientSecret(clientSecret string) Option {
	return func(a *Authorizer) {
		a.clientSecret = clientSecret
	}
}

// WithRedirectURL configures the authorization handler
// with the redirect_url
func WithRedirectURL(redirectURL string) Option {
	return func(a *Authorizer) {
		a.redirectURL = redirectURL
	}
}

// Authorizer configures a Bitbucket auth provider.
type Authorizer struct {
	redirectURL  string
	clientID     string
	clientSecret string
	client       *http.Client
}

func newDefault() *Authorizer {
	return &Authorizer{
		client: http.DefaultClient,
	}
}

// New returns a Bitbucket authorization provider.
func New(opts ...Option) login.Authorizer {
	auther := newDefault()
	for _, opt := range opts {
		opt(auther)
	}
	return auther
}

// Authorize returns a http.Handler that runs h at the
// completion of the GitHub authorization flow. The GitHub
// authorization details are available to h in the
// http.Request context.
func (a *Authorizer) Authorize(h http.Handler) http.Handler {
	return oauth2.Handler(h, &oauth2.Config{
		Client:           a.client,
		ClientID:         a.clientID,
		ClientSecret:     a.clientSecret,
		RedirectURL:      a.redirectURL,
		AccessTokenURL:   accessTokenURL,
		AuthorizationURL: authorizationURL,
	})
}
