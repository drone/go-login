// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package github

import (
	"net/http"
	"strings"

	"github.com/drone/go-login/login"
	"github.com/drone/go-login/login/internal/oauth2"
)

// Authorizer configures a GitHub authorization
// provider.
type Authorizer struct {
	scope        []string
	clientID     string
	clientSecret string
	server       string
	client       *http.Client
}

func newDefault() *Authorizer {
	return &Authorizer{
		server: "https://github.com",
		client: http.DefaultClient,
	}
}

// Option configures an authorization handler option.
type Option func(a *Authorizer)

// WithClient configures the authorization handler with a
// custom http.Client.
func WithClient(client *http.Client) Option {
	return func(o *Authorizer) {
		o.client = client
	}
}

// WithClientID configures the authorization handler with
// the client_id.
func WithClientID(clientID string) Option {
	return func(o *Authorizer) {
		o.clientID = clientID
	}
}

// WithClientSecret configures the authorization handler
// with the client_secret.
func WithClientSecret(clientSecret string) Option {
	return func(o *Authorizer) {
		o.clientSecret = clientSecret
	}
}

// WithScope configures the authorization handler with the
// these scopes.
func WithScope(scope ...string) Option {
	return func(o *Authorizer) {
		o.scope = scope
	}
}

// WithAddress configures the authorization handler with
// a GitHub enterprise server address.
func WithAddress(server string) Option {
	return func(o *Authorizer) {
		o.server = strings.TrimSuffix(server, "/")
	}
}

// New returns a GitHub authorization provider.
func New(opts ...Option) login.Authorizer {
	v := newDefault()
	for _, opt := range opts {
		opt(v)
	}
	return v
}

// Authorize returns a http.Handler that runs h at the
// completion of the GitHub authorization flow. The GitHub
// authorization details are available to h in the
// http.Request context.
func (a *Authorizer) Authorize(h http.Handler) http.Handler {
	return oauth2.Handler(h, &oauth2.Config{
		BasicAuthOff:     true,
		Client:           a.client,
		ClientID:         a.clientID,
		ClientSecret:     a.clientSecret,
		AccessTokenURL:   a.server + "/login/oauth/access_token",
		AuthorizationURL: a.server + "/login/oauth/authorize",
		Scope:            a.scope,
	})
}
