// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package github

import (
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func TestWithClient(t *testing.T) {
	c := &http.Client{}
	v := &Authorizer{}
	WithClient(c)(v)
	if got, want := v.client, c; got != want {
		t.Errorf("Expect custom client option applied")
	}
}

func TestWithClientID(t *testing.T) {
	v := &Authorizer{}
	WithClientID("3da54155991")(v)
	if got, want := v.clientID, "3da54155991"; got != want {
		t.Errorf("Expect client_id option applied")
	}
}

func TestWithClientSecret(t *testing.T) {
	v := &Authorizer{}
	WithClientSecret("5012f6c60b2")(v)
	if got, want := v.clientSecret, "5012f6c60b2"; got != want {
		t.Errorf("Expect client_secret option applied")
	}
}

func TestWithAddress(t *testing.T) {
	v := &Authorizer{}
	WithAddress("https://company.github.com/")(v)
	if strings.HasSuffix(v.server, "/") {
		t.Errorf("Expect trailing slash removed from server address")
	}
	if got, want := v.server, "https://company.github.com"; got != want {
		t.Errorf("Expect server address option applied")
	}
}

func TestWithScope(t *testing.T) {
	v := &Authorizer{}
	WithScope("user", "repo")(v)
	if got, want := v.scope, []string{"user", "repo"}; !reflect.DeepEqual(got, want) {
		t.Errorf("Expect scope option applied")
	}
}

func TestDefaultAuthorizer(t *testing.T) {
	v := newDefault()
	if got, want := v.client, http.DefaultClient; got != want {
		t.Errorf("Expect default client is http.DefaultClient")
	}
	if got, want := v.server, "https://github.com"; got != want {
		t.Errorf("Expect default server is https://github.com")
	}
}
