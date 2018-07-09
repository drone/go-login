// Copyright 2018 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stash

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/drone/go-login/login"
	"github.com/drone/go-login/login/internal/oauth1"
)

const (
	requestTokenURL   = "%s/plugins/servlet/oauth/request-token"
	authorizeTokenURL = "%s/plugins/servlet/oauth/authorize"
	accessTokenURL    = "%s/plugins/servlet/oauth/access-token"
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

// WithConsumerKey configures the authorization handler with
// the oauth_consumer_key.
func WithConsumerKey(consumerKey string) Option {
	return func(a *Authorizer) {
		a.consumerKey = consumerKey
	}
}

// WithConsumerSecret configures the authorization handler
// with the oauth_consumer_secret.
func WithConsumerSecret(consumerSecret string) Option {
	return func(a *Authorizer) {
		a.consumerSecret = consumerSecret
	}
}

// WithCallbackURL configures the authorization handler
// with the oauth_callback_url
func WithCallbackURL(callbackURL string) Option {
	return func(a *Authorizer) {
		a.callbackURL = callbackURL
	}
}

// WithPrivateKeyFile configures the authorization handler
// with the oauth private rsa key for signing requests.
func WithPrivateKeyFile(path string) Option {
	d, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return WithPrivateKey(d)
}

// WithPrivateKey configures the authorization handler
// with the oauth private rsa key for signing requests.
func WithPrivateKey(data []byte) Option {
	return func(a *Authorizer) {
		p, _ := pem.Decode(data)
		k, err := x509.ParsePKCS1PrivateKey(p.Bytes)
		if err != nil {
			panic(err)
		}
		a.signer = &oauth1.RSASigner{PrivateKey: k}
	}
}

// Authorizer configures the Bitbucket Server (Stash)
// authorization provider.
type Authorizer struct {
	callbackURL    string
	address        string
	consumerKey    string
	consumerSecret string
	signer         oauth1.Signer
	client         *http.Client
}

func newDefault() *Authorizer {
	return &Authorizer{
		client: http.DefaultClient,
	}
}

// New returns a Bitbucket Server authorization provider.
func New(address string, opts ...Option) login.Authorizer {
	auther := newDefault()
	auther.address = strings.TrimPrefix(address, "/")
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
	return oauth1.Handler(h, &oauth1.Config{
		Signer:           a.signer,
		Client:           a.client,
		ConsumerKey:      a.consumerKey,
		ConsumerSecret:   a.consumerSecret,
		CallbackURL:      a.callbackURL,
		AccessTokenURL:   fmt.Sprintf(accessTokenURL, a.address),
		AuthorizationURL: fmt.Sprintf(authorizeTokenURL, a.address),
		RequestTokenURL:  fmt.Sprintf(requestTokenURL, a.address),
	})
}
