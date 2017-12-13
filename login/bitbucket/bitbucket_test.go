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
	o := &Options{}
	WithClient(c)(o)
	if got, want := o.client, c; got != want {
		t.Errorf("Expect custom client")
	}
}

func TestWithClientID(t *testing.T) {
	opts := &Options{}
	WithClientID("3da54155991")(opts)
	if got, want := opts.clientID, "3da54155991"; got != want {
		t.Errorf("Expect custom client_id")
	}
}

func TestWithClientSecret(t *testing.T) {
	opts := &Options{}
	WithClientSecret("5012f6c60b2")(opts)
	if got, want := opts.clientSecret, "5012f6c60b2"; got != want {
		t.Errorf("Expect custom client_secret")
	}
}

func TestWithRedirectURL(t *testing.T) {
	opts := &Options{}
	WithRedirectURL("http://company.com/login")(opts)
	if got, want := opts.redirectURL, "http://company.com/login"; got != want {
		t.Errorf("Expect custom redirect_uri")
	}
}

func TestDefaultOptions(t *testing.T) {
	opts := createOptions()
	if got, want := opts.client, http.DefaultClient; got != want {
		t.Errorf("Expect default client is http.DefaultClient")
	}
}
