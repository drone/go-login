// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gitea

import (
	"net/http"
	"strings"
)

// Options provides the Gogs authentication options.
type Options struct {
	label  string
	login  string
	server string
	client *http.Client
}

func defaultOptions() *Options {
	return &Options{
		label:  "default",
		client: http.DefaultClient,
	}
}

// Option configures an authorization handler option.
type Option func(o *Options)

// WithAddress configures the authorization handler with
// the progived Gogs server address.
func WithAddress(server string) Option {
	return func(o *Options) {
		o.server = strings.TrimPrefix(server, "/")
	}
}

// WithClient configures the authorization handler with a
// custom http.Client.
func WithClient(client *http.Client) Option {
	return func(o *Options) {
		o.client = client
	}
}

// WithTokenName configures the authorization handler to
// use the specificed token name when finding and creating
// authorization tokens.
func WithTokenName(name string) Option {
	return func(o *Options) {
		o.label = name
	}
}

// WithLoginRedirect configures the authorization handler
// to redirect the http.Request to the login form when the
// username or password are missing from the Form data.
func WithLoginRedirect(path string) Option {
	return func(o *Options) {
		o.login = path
	}
}
