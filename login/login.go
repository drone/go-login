// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package login

import (
	"context"
	"net/http"
	"time"
)

// Authorizer returns a http.Handler that runs h at the
// completion of the authorization flow. The authorization
// results are available to h in the http.Request context.
type Authorizer interface {
	Authorize(h http.Handler) http.Handler
}

// Middleware defines a Login middleware. The middleware
// wraps the http.Handler and intercepts the http.Request
// to perform authentication. The http.Handler is invoked
// when authentication is complete, with authentication
// details (oauth token, etc) passed to the handler via
// the http.Request context.
type Middleware interface {
	Wrap(h http.Handler) http.Handler
}

// Token represents an authorization token.
type Token struct {
	Access  string
	Refresh string
	Expires time.Time
}

type key int

const (
	tokenKey key = iota
	errorKey
)

// WithToken returns a parent context with the token.
func WithToken(parent context.Context, token *Token) context.Context {
	return context.WithValue(parent, tokenKey, token)
}

// WithError returns a parent context with the error.
func WithError(parent context.Context, err error) context.Context {
	return context.WithValue(parent, errorKey, err)
}

// TokenFrom returns the login token rom the context.
func TokenFrom(ctx context.Context) *Token {
	token, _ := ctx.Value(tokenKey).(*Token)
	return token
}

// ErrorFrom returns the login error from the context.
func ErrorFrom(ctx context.Context) error {
	err, _ := ctx.Value(errorKey).(error)
	return err
}
