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
	o := &Options{}
	WithClient(c)(o)
	if got, want := o.client, c; got != want {
		t.Errorf("Expect custom client option applied")
	}
}

func TestWithClientID(t *testing.T) {
	opts := &Options{}
	WithClientID("3da54155991")(opts)
	if got, want := opts.clientID, "3da54155991"; got != want {
		t.Errorf("Expect client_id option applied")
	}
}

func TestWithClientSecret(t *testing.T) {
	opts := &Options{}
	WithClientSecret("5012f6c60b2")(opts)
	if got, want := opts.clientSecret, "5012f6c60b2"; got != want {
		t.Errorf("Expect client_secret option applied")
	}
}

func TestWithAddress(t *testing.T) {
	opts := &Options{}
	WithAddress("https://company.github.com/")(opts)
	if strings.HasSuffix(opts.server, "/") {
		t.Errorf("Expect trailing slash removed from server address")
	}
	if got, want := opts.server, "https://company.github.com"; got != want {
		t.Errorf("Expect server address option applied")
	}
}

func TestWithScope(t *testing.T) {
	opts := &Options{}
	WithScope("user", "repo")(opts)
	if got, want := opts.scope, []string{"user", "repo"}; !reflect.DeepEqual(got, want) {
		t.Errorf("Expect scope option applied")
	}
}

func TestDefaultOptions(t *testing.T) {
	opts := defaultOptions()
	if got, want := opts.client, http.DefaultClient; got != want {
		t.Errorf("Expect default client is http.DefaultClient")
	}
	if got, want := opts.server, "https://github.com"; got != want {
		t.Errorf("Expect default server is https://github.com")
	}
}
