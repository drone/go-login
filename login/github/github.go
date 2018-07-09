// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package github

import (
	"net/http"
	"strings"

	"github.com/drone/go-login/login/internal/oauth2"
)

// Authorizer configures a GitHub authorization
// provider.
type Authorizer struct {
	Client       *http.Client
	ClientID     string
	ClientSecret string
	Server       string
	Scope        []string
}

// Authorize returns a http.Handler that runs h at the
// completion of the GitHub authorization flow. The GitHub
// authorization details are available to h in the
// http.Request context.
func (a *Authorizer) Authorize(h http.Handler) http.Handler {
	server := normalizeAddress(a.Server)
	return oauth2.Handler(h, &oauth2.Config{
		BasicAuthOff:     true,
		Client:           a.Client,
		ClientID:         a.ClientID,
		ClientSecret:     a.ClientSecret,
		AccessTokenURL:   server + "/login/oauth/access_token",
		AuthorizationURL: server + "/login/oauth/authorize",
		Scope:            a.Scope,
	})
}

func normalizeAddress(address string) string {
	if address == "" {
		return "https://github.com"
	}
	return strings.TrimSuffix(address, "/")
}
