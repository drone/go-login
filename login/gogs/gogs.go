// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gogs

import (
	"net/http"
	"strings"
)

// Authorizer configures the Gogs auth provider.
type Authorizer struct {
	Label  string
	Login  string
	Server string
	Client *http.Client
}

// Authorize returns a http.Handler that runs h at the
// completion of the GitLab authorization flow. The GitLab
// authorization details are available to h in the
// http.Request context.
func (a *Authorizer) Authorize(h http.Handler) http.Handler {
	v := &handler{
		next:   h,
		label:  a.Label,
		login:  a.Login,
		server: strings.TrimSuffix(a.Server, "/"),
		client: a.Client,
	}
	if v.client == nil {
		v.client = http.DefaultClient
	}
	if v.label == "" {
		v.label = "default"
	}
	return v
}
