// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gogs

import (
	"net/http"
	"testing"
)

func TestAuthorizer(t *testing.T) {
	h := http.RedirectHandler("/", 302)
	c := new(http.Client)
	a := Config{
		Label:  "drone",
		Login:  "/path/to/login",
		Server: "https://try.gogs.io/",
		Client: c,
	}
	v := a.Handler(h).(*handler)
	if got, want := v.login, "/path/to/login"; got != want {
		t.Errorf("Expect login redirect url %q, got %q", want, got)
	}
	if got, want := v.server, "https://try.gogs.io"; got != want {
		t.Errorf("Expect server address %q, got %q", want, got)
	}
	if got, want := v.label, "drone"; got != want {
		t.Errorf("Expect label %q, got %q", want, got)
	}
	if got, want := v.client, c; got != want {
		t.Errorf("Expect custom client")
	}
	if got, want := v.next, h; got != want {
		t.Errorf("Expect handler wrapped")
	}
}

func TestAuthorizerDefault(t *testing.T) {
	a := Config{
		Login:  "/path/to/login",
		Server: "https://try.gogs.io",
	}
	v := a.Handler(
		http.NotFoundHandler(),
	).(*handler)
	if got, want := v.label, "default"; got != want {
		t.Errorf("Expect label %q, got %q", want, got)
	}
	if got, want := v.client, http.DefaultClient; got != want {
		t.Errorf("Expect custom client")
	}
}
