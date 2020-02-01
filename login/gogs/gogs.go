// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gogs

import (
	"net/http"
	"strings"

	"github.com/drone/go-login/login"
)

var _ login.Middleware = (*Config)(nil)

// Config configures the Gogs auth provider.
type Config struct {
	Label  string
	Login  string
	Server string
	Client *http.Client
}

// Handler returns a http.Handler that runs h at the
// completion of the GitLab authorization flow. The GitLab
// authorization details are available to h in the
// http.Request context.
func (c *Config) Handler(h http.Handler) http.Handler {
	v := &handler{
		next:   h,
		label:  c.Label,
		login:  c.Login,
		server: strings.TrimSuffix(c.Server, "/"),
		client: c.Client,
	}
	if v.client == nil {
		v.client = http.DefaultClient
	}
	if v.label == "" {
		v.label = "default"
	}
	return v
}
