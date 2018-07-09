// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gogs

import (
	"net/http"
	"strings"

	"github.com/drone/go-login/login"
)

// Authorizer configures the Gogs auth provider.
type Authorizer struct {
	label  string
	login  string
	server string
	client *http.Client
}

func newDefault() *Authorizer {
	return &Authorizer{
		label:  "default",
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

// WithTokenName configures the authorization handler to
// use the specificed token name when finding and creating
// authorization tokens.
func WithTokenName(name string) Option {
	return func(a *Authorizer) {
		a.label = name
	}
}

// WithLoginRedirect configures the authorization handler
// to redirect the http.Request to the login form when the
// username or password are missing from the Form data.
func WithLoginRedirect(path string) Option {
	return func(a *Authorizer) {
		a.login = path
	}
}

// New returns a Gogs authorization provider.
func New(address string, opts ...Option) login.Authorizer {
	auther := newDefault()
	auther.server = strings.TrimSuffix(address, "/")
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
	return &handler{
		next:   h,
		label:  a.label,
		login:  a.login,
		server: a.server,
		client: a.client,
	}
}
