// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bitbucket

import (
	"net/http"

	"github.com/drone/go-login/login/internal/oauth2"
)

const (
	accessTokenURL   = "https://bitbucket.org/site/oauth2/access_token"
	authorizationURL = "https://bitbucket.org/site/oauth2/authorize"
)

// Authorizer configures a Bitbucket auth provider.
type Authorizer struct {
	Client       *http.Client
	ClientID     string
	ClientSecret string
	RedirectURL  string
}

// Authorize returns a http.Handler that runs h at the
// completion of the GitHub authorization flow. The GitHub
// authorization details are available to h in the
// http.Request context.
func (a *Authorizer) Authorize(h http.Handler) http.Handler {
	return oauth2.Handler(h, &oauth2.Config{
		Client:           a.Client,
		ClientID:         a.ClientID,
		ClientSecret:     a.ClientSecret,
		RedirectURL:      a.RedirectURL,
		AccessTokenURL:   accessTokenURL,
		AuthorizationURL: authorizationURL,
	})
}
