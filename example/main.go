// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/drone/go-login/login"
	"github.com/drone/go-login/login/bitbucket"
	"github.com/drone/go-login/login/github"
	"github.com/drone/go-login/login/gitlab"
	"github.com/drone/go-login/login/gogs"
)

var (
	provider     = flag.String("provider", "github", "")
	providerURL  = flag.String("provider-url", "", "")
	clientID     = flag.String("client-id", "", "")
	clientSecret = flag.String("client-secret", "", "")
	redirectURL  = flag.String("redirect-url", "http://localhost:8080/login", "")
	address      = flag.String("address", ":8080", "")
	help         = flag.Bool("help", false, "")
)

func main() {
	flag.Usage = usage
	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	var auther http.Handler
	switch *provider {
	case "gogs", "gitea":
		auther = gogs.New(
			http.HandlerFunc(details),
			gogs.WithAddress(*providerURL),
			gogs.WithLoginRedirect("/login/form"),
		)
	case "gitlab":
		auther = gitlab.New(
			http.HandlerFunc(details),
			gitlab.WithClientID(*clientID),
			gitlab.WithClientSecret(*clientSecret),
			gitlab.WithRedirectURL(*redirectURL),
			gitlab.WithScope("read_user", "api"),
		)
	case "github":
		auther = github.New(
			http.HandlerFunc(details),
			github.WithClientID(*clientID),
			github.WithClientSecret(*clientSecret),
			github.WithScope("repo", "user", "read:org"),
		)
	case "bitbucket":
		auther = bitbucket.New(
			http.HandlerFunc(details),
			bitbucket.WithClientID(*clientID),
			bitbucket.WithClientSecret(*clientSecret),
			bitbucket.WithRedirectURL(*redirectURL),
		)
	}

	// handles the authorization flow and displays the
	// authorization results at completion.
	http.Handle("/login", auther)
	http.Handle("/login/form", http.HandlerFunc(form))

	// redirects the user to the login handler.
	http.Handle("/", http.RedirectHandler("/login", http.StatusSeeOther))
	http.ListenAndServe(*address, nil)
}

// returns the login credentials.
func details(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	err := login.ErrorFrom(ctx)
	if err != nil {
		fmt.Fprintf(w, failure, err)
		return
	}
	token := login.TokenFrom(ctx)
	fmt.Fprintf(w, success, token.Access)
}

// display the login form.
func form(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, loginForm)
}

// html page displayed to collect credentials.
var loginForm = `
<form method="POST" action="/login">
<input type="text" name="username" />
<input type="password" name="password" />
<input type="submit" />
</form>
`

// html page displayed on success.
var success = `
<html>
<body>
<h1>Token</h1>
<h2>%s</h2>
</body>
</html>
`

// html page displayed on failure.
var failure = `
<html>
<body>
<h1>Error</h1>
<h2>%s</h2>
</body>
</html>
`

func usage() {
	fmt.Println(`Usage: go run main.go [OPTION]...
  --provider           oauth provider (github, gitlab, gogs, gitea, bitbucket)
  --provider-url       oauth provider url (gitea, gogs only)
  --client-id          oauth client id
  --client-secret      oauth client secret
  --redirect-url       oauth redirect url
  --address            http server address (:8080)
  --help               display this help and exit`)
}
