// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bitbucket

import (
	"net/http"
	"testing"
)

func TestWithClient(t *testing.T) {
	c := &http.Client{}
	a := &Authorizer{}
	WithClient(c)(a)
	if got, want := a.client, c; got != want {
		t.Errorf("Expect custom client")
	}
}

func TestWithClientID(t *testing.T) {
	a := &Authorizer{}
	WithClientID("3da54155991")(a)
	if got, want := a.clientID, "3da54155991"; got != want {
		t.Errorf("Expect custom client_id")
	}
}

func TestWithClientSecret(t *testing.T) {
	a := &Authorizer{}
	WithClientSecret("5012f6c60b2")(a)
	if got, want := a.clientSecret, "5012f6c60b2"; got != want {
		t.Errorf("Expect custom client_secret")
	}
}

func TestWithRedirectURL(t *testing.T) {
	a := &Authorizer{}
	WithRedirectURL("http://company.com/login")(a)
	if got, want := a.redirectURL, "http://company.com/login"; got != want {
		t.Errorf("Expect custom redirect_uri")
	}
}

func TestDefaultAuthorizer(t *testing.T) {
	a := newDefault()
	if got, want := a.client, http.DefaultClient; got != want {
		t.Errorf("Expect default client is http.DefaultClient")
	}
}
