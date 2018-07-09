// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gitea

import (
	"net/http"
	"testing"
)

func TestWithAddress(t *testing.T) {
	h := New(nil, WithAddress("https://try.gitea.io"))
	if got, want := h.(*handler).server, "https://try.gitea.io"; got != want {
		t.Errorf("Expect server address %q, got %q", want, got)
	}
}

func TestWithClient(t *testing.T) {
	c := &http.Client{}
	h := New(nil, WithClient(c))
	if got, want := h.(*handler).client, c; got != want {
		t.Errorf("Expect custom client")
	}
}

func TestWithTokenName(t *testing.T) {
	h := New(nil, WithTokenName("some_token"))
	if got, want := h.(*handler).label, "some_token"; got != want {
		t.Errorf("Expect token name url %q, got %q", want, got)
	}
}

func TestWithLoginRedirect(t *testing.T) {
	h := New(nil, WithLoginRedirect("/path/to/login"))
	if got, want := h.(*handler).login, "/path/to/login"; got != want {
		t.Errorf("Expect login redirect url %q, got %q", want, got)
	}
}
