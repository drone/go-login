// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package login

import (
	"context"
	"net/http"
	"time"
)

// Middleware provides login middleware.
type Middleware interface {
	// Handler returns a http.Handler that runs h at the
	// completion of the authorization flow. The authorization
	// results are available to h in the http.Request context.
	Handler(h http.Handler) http.Handler
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
