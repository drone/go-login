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

	"github.com/drone/go-login/login/internal/oauth1"
)

const (
	requestTokenURL   = "%s/plugins/servlet/oauth/request-token"
	authorizeTokenURL = "%s/plugins/servlet/oauth/authorize"
	accessTokenURL    = "%s/plugins/servlet/oauth/access-token"
)

// Options provides the Bitbucket Server (Stash)
// authentication options.
type Options struct {
	callbackURL    string
	address        string
	consumerKey    string
	consumerSecret string
	signer         oauth1.Signer
	client         *http.Client
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

// WithConsumerKey configures the authorization handler with
// the oauth_consumer_key.
func WithConsumerKey(consumerKey string) Option {
	return func(o *Options) {
		o.consumerKey = consumerKey
	}
}

// WithConsumerSecret configures the authorization handler
// with the oauth_consumer_secret.
func WithConsumerSecret(consumerSecret string) Option {
	return func(o *Options) {
		o.consumerSecret = consumerSecret
	}
}

// WithCallbackURL configures the authorization handler
// with the oauth_callback_url
func WithCallbackURL(callbackURL string) Option {
	return func(o *Options) {
		o.callbackURL = callbackURL
	}
}

// WithAddress configures the authorization handler with
// the Bitbucket Server address.
func WithAddress(address string) Option {
	return func(o *Options) {
		o.address = strings.TrimPrefix(address, "/")
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
	return func(o *Options) {
		p, _ := pem.Decode(data)
		k, err := x509.ParsePKCS1PrivateKey(p.Bytes)
		if err != nil {
			panic(err)
		}
		o.signer = &oauth1.RSASigner{PrivateKey: k}
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
	return oauth1.Handler(h, &oauth1.Config{
		Signer:           opts.signer,
		Client:           opts.client,
		ConsumerKey:      opts.consumerKey,
		ConsumerSecret:   opts.consumerSecret,
		CallbackURL:      opts.callbackURL,
		AccessTokenURL:   fmt.Sprintf(accessTokenURL, opts.address),
		AuthorizationURL: fmt.Sprintf(authorizeTokenURL, opts.address),
		RequestTokenURL:  fmt.Sprintf(requestTokenURL, opts.address),
	})
}
