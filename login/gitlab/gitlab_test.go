// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gitlab

import (
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func TestWithClient(t *testing.T) {
	c := &http.Client{}
	a := &Authorizer{}
	WithClient(c)(a)
	if got, want := a.client, c; got != want {
		t.Errorf("Expect custom client option applied")
	}
}

func TestWithClientID(t *testing.T) {
	a := &Authorizer{}
	WithClientID("3da54155991")(a)
	if got, want := a.clientID, "3da54155991"; got != want {
		t.Errorf("Expect client_id option applied")
	}
}

func TestWithClientSecret(t *testing.T) {
	a := &Authorizer{}
	WithClientSecret("5012f6c60b2")(a)
	if got, want := a.clientSecret, "5012f6c60b2"; got != want {
		t.Errorf("Expect client_secret option applied")
	}
}

func TestWithRedirectURL(t *testing.T) {
	a := &Authorizer{}
	WithRedirectURL("http://company.com/login")(a)
	if got, want := a.redirectURL, "http://company.com/login"; got != want {
		t.Errorf("Expect redirect_uri option applied")
	}
}

func TestWithAddress(t *testing.T) {
	a := &Authorizer{}
	WithAddress("https://company.gitlab.com/")(a)
	if strings.HasSuffix(a.server, "/") {
		t.Errorf("Expect trailing slash removed from server address")
	}
	if got, want := a.server, "https://company.gitlab.com"; got != want {
		t.Errorf("Expect server address option applied")
	}
}

func TestWithScope(t *testing.T) {
	a := &Authorizer{}
	WithScope("read_user", "api")(a)
	if got, want := a.scope, []string{"read_user", "api"}; !reflect.DeepEqual(got, want) {
		t.Errorf("Expect scope option applied")
	}
}

func TestDefaultAuthorizer(t *testing.T) {
	a := newDefault()
	if got, want := a.client, http.DefaultClient; got != want {
		t.Errorf("Expect default client is http.DefaultClient")
	}
	if got, want := a.server, "https://gitlab.com"; got != want {
		t.Errorf("Expect default server is https://gitlab.com")
	}
}
