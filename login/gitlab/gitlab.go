// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gitlab

import (
	"net/http"
	"strings"

	"github.com/drone/go-login/login/internal/oauth2"
)

// Authorizer configures the GitLab auth provider.
type Authorizer struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
	Server       string
	Scope        []string
	Client       *http.Client
}

// Authorize returns a http.Handler that runs h at the
// completion of the GitLab authorization flow. The GitLab
// authorization details are available to h in the
// http.Request context.
func (a *Authorizer) Authorize(h http.Handler) http.Handler {
	server := normalizeAddress(a.Server)
	return oauth2.Handler(h, &oauth2.Config{
		BasicAuthOff:     true,
		Client:           a.Client,
		ClientID:         a.ClientID,
		ClientSecret:     a.ClientSecret,
		RedirectURL:      a.RedirectURL,
		AccessTokenURL:   server + "/oauth/token",
		AuthorizationURL: server + "/oauth/authorize",
		Scope:            a.Scope,
	})
}

func normalizeAddress(address string) string {
	if address == "" {
		return "https://gitlab.com"
	}
	return strings.TrimSuffix(address, "/")
}
