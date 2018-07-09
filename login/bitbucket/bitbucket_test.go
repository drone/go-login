// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bitbucket

import (
	"net/http"
	"testing"
)

func TestAuthorizer(t *testing.T) {
	c := &http.Client{}
	v := New(
		WithClient(c),
		WithClientID("3da54155991"),
		WithClientSecret("5012f6c60b2"),
		WithRedirectURL("http://company.com/login"),
	).(*Authorizer)

	if got, want := v.client, c; got != want {
		t.Errorf("Expect custom client")
	}

	if got, want := v.clientID, "3da54155991"; got != want {
		t.Errorf("Expect custom client_id")
	}

	if got, want := v.clientSecret, "5012f6c60b2"; got != want {
		t.Errorf("Expect custom client_secret")
	}

	if got, want := v.redirectURL, "http://company.com/login"; got != want {
		t.Errorf("Expect custom redirect_uri")
	}
}

func TestAuthorizerDefault(t *testing.T) {
	v := New().(*Authorizer)
	if got, want := v.client, http.DefaultClient; got != want {
		t.Errorf("Expect custom client")
	}
}
