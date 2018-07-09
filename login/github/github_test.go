// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package github

import (
	"net/http"
	"reflect"
	"testing"
)

func TestAuthorizer(t *testing.T) {
	c := &http.Client{}
	v := New(
		WithClient(c),
		WithClientID("3da54155991"),
		WithClientSecret("5012f6c60b2"),
		WithAddress("https://company.github.com"),
		WithScope("user", "repo"),
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

	if got, want := v.scope, []string{"user", "repo"}; !reflect.DeepEqual(got, want) {
		t.Errorf("Expect scope option applied")
	}

	if got, want := v.server, "https://company.github.com"; got != want {
		t.Errorf("Expect server address option applied")
	}
}

func TestAuthorizerDefault(t *testing.T) {
	v := New().(*Authorizer)
	if got, want := v.client, http.DefaultClient; got != want {
		t.Errorf("Expect custom client")
	}

	if got, want := v.server, "https://github.com"; got != want {
		t.Errorf("Expect server address option applied")
	}
}
