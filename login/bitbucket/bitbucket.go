// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bitbucket

import (
	"net/http"

	"github.com/drone/go-login/login/internal/oauth2"
)

// Options provides the Bitbucket authentication options.
type Options struct {
	redirectURL  string
	clientID     string
	clientSecret string
	client       *http.Client
}

func createOptions() *Options {
	return &Options{
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

// WithRedirectURL configures the authorization handler
// with the redirect_url
func WithRedirectURL(redirectURL string) Option {
	return func(o *Options) {
		o.redirectURL = redirectURL
	}
}

// New returns a http.Handler that runs h at the completion
// of the Bitbucket authorization flow. The Bitbucket
// authorization is passed to h in the http.Request context.
func New(h http.Handler, opt ...Option) http.Handler {
	opts := createOptions()
	for _, fn := range opt {
		fn(opts)
	}
	return oauth2.Handler(h, &oauth2.Config{
		Client:           opts.client,
		ClientID:         opts.clientID,
		ClientSecret:     opts.clientSecret,
		RedirectURL:      opts.redirectURL,
		AccessTokenURL:   "https://bitbucket.org/site/oauth2/access_token",
		AuthorizationURL: "https://bitbucket.org/site/oauth2/authorize",
	})
}
