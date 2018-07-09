// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gitlab

import (
	"net/http"
	"strings"

	"github.com/drone/go-login/login"
	"github.com/drone/go-login/login/internal/oauth2"
)

// Authorizer configures the GitLab auth provider.
type Authorizer struct {
	scope        []string
	clientID     string
	clientSecret string
	redirectURL  string
	server       string
	client       *http.Client
}

func newDefault() *Authorizer {
	return &Authorizer{
		server: "https://gitlab.com",
		client: http.DefaultClient,
	}
}

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

// WithScope configures the authorization handler with the
// these scopes.
func WithScope(scope ...string) Option {
	return func(a *Authorizer) {
		a.scope = scope
	}
}

// WithAddress configures the authorization handler with
// the the self-hosted GitLab server address.
func WithAddress(server string) Option {
	return func(a *Authorizer) {
		if server != "" {
			a.server = strings.TrimSuffix(server, "/")
		}
	}
}

// New returns a GitLab authorization provider.
func New(opts ...Option) login.Authorizer {
	auther := newDefault()
	for _, opt := range opts {
		opt(auther)
	}
	return auther
}

// Authorize returns a http.Handler that runs h at the
// completion of the GitLab authorization flow. The GitLab
// authorization details are available to h in the
// http.Request context.
func (a *Authorizer) Authorize(h http.Handler) http.Handler {
	return oauth2.Handler(h, &oauth2.Config{
		BasicAuthOff:     true,
		Client:           a.client,
		ClientID:         a.clientID,
		ClientSecret:     a.clientSecret,
		RedirectURL:      a.redirectURL,
		AccessTokenURL:   a.server + "/oauth/token",
		AuthorizationURL: a.server + "/oauth/authorize",
		Scope:            a.scope,
	})
}
