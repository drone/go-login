// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gogs

import (
	"net/http"
	"testing"
)

func TestWithClient(t *testing.T) {
	c := &http.Client{}
	h := New("https://try.gogs.io", WithClient(c))
	if got, want := h.(*Authorizer).client, c; got != want {
		t.Errorf("Expect custom client")
	}
}

func TestWithTokenName(t *testing.T) {
	h := New("https://try.gogs.io", WithTokenName("some_token"))
	if got, want := h.(*Authorizer).label, "some_token"; got != want {
		t.Errorf("Expect token name url %q, got %q", want, got)
	}
}

func TestWithLoginRedirect(t *testing.T) {
	h := New("https://try.gogs.io", WithLoginRedirect("/path/to/login"))
	if got, want := h.(*Authorizer).login, "/path/to/login"; got != want {
		t.Errorf("Expect login redirect url %q, got %q", want, got)
	}
}
