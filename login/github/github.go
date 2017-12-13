// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package github

import (
	"net/http"
	"strings"

	"github.com/drone/go-login/login/internal/oauth2"
)

// Options provides the GitHub authentication options.
type Options struct {
	scope        []string
	clientID     string
	clientSecret string
	server       string
	client       *http.Client
}

func defaultOptions() *Options {
	return &Options{
		server: "https://github.com",
		client: http.DefaultClient,
	}
}

// Option configures an authorization handler option.
type Option func(o *Options)

// WithClient configures the authorization handler with a
// custom http.Client.
func WithClient(client *http.Client) Option {
	return func(o *Options) {
		o.client = client
	}
}

// WithClientID configures the authorization handler with
// the client_id.
func WithClientID(clientID string) Option {
	return func(o *Options) {
		o.clientID = clientID
	}
}

// WithClientSecret configures the authorization handler
// with the client_secret.
func WithClientSecret(clientSecret string) Option {
	return func(o *Options) {
		o.clientSecret = clientSecret
	}
}

// WithScope configures the authorization handler with the
// these scopes.
func WithScope(scope ...string) Option {
	return func(o *Options) {
		o.scope = scope
	}
}

// WithAddress configures the authorization handler with
// a GitHub enterprise server address.
func WithAddress(server string) Option {
	return func(o *Options) {
		o.server = strings.TrimSuffix(server, "/")
	}
}

// New returns a http.Handler that runs h at the completion
// of the GitHub authorization flow. The GitHub authorization
// is availabe to h in the http.Request context.
func New(h http.Handler, opt ...Option) http.Handler {
	opts := defaultOptions()
	for _, fn := range opt {
		fn(opts)
	}
	return oauth2.Handler(h, &oauth2.Config{
		BasicAuthOff:     true,
		Client:           opts.client,
		ClientID:         opts.clientID,
		ClientSecret:     opts.clientSecret,
		AccessTokenURL:   opts.server + "/login/oauth/access_token",
		AuthorizationURL: opts.server + "/login/oauth/authorize",
		Scope:            opts.scope,
	})
}
