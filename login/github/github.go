// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package github

import (
	"net/http"
	"strings"

	"github.com/drone/go-login/login"
	"github.com/drone/go-login/login/internal/oauth2"
	"github.com/drone/go-login/login/logger"
)

var _ login.Middleware = (*Config)(nil)

// Config configures a GitHub authorization provider.
type Config struct {
	Client       *http.Client
	ClientID     string
	ClientSecret string
	Server       string
	Scope        []string
	Logger       logger.Logger
	Dumper       logger.Dumper
}

// Handler returns a http.Handler that runs h at the
// completion of the GitHub authorization flow. The GitHub
// authorization details are available to h in the
// http.Request context.
func (c *Config) Handler(h http.Handler) http.Handler {
	server := normalizeAddress(c.Server)
	return oauth2.Handler(h, &oauth2.Config{
		BasicAuthOff:     true,
		Client:           c.Client,
		ClientID:         c.ClientID,
		ClientSecret:     c.ClientSecret,
		AccessTokenURL:   server + "/login/oauth/access_token",
		AuthorizationURL: server + "/login/oauth/authorize",
		Scope:            c.Scope,
		Logger:           c.Logger,
		Dumper:           c.Dumper,
	})
}

func normalizeAddress(address string) string {
	if address == "" {
		return "https://github.com"
	}
	return strings.TrimSuffix(address, "/")
}
