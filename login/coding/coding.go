package coding

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/sennotech/go-login/login"
	"github.com/sennotech/go-login/login/internal/oauth2"
	"github.com/sennotech/go-login/login/logger"
)

var _ login.Middleware = (*Config)(nil)

// Config configures a GitHub authorization provider.
type Config struct {
	Client       *http.Client
	ClientID     string
	ClientSecret string
	Server       string
	Scope        []string
	Logger       logger.Logger
	Dumper       logger.Dumper
}

// Handler returns a http.Handler that runs h at the
// completion of the GitHub authorization flow. The GitHub
// authorization details are available to h in the
// http.Request context.
func (c *Config) Handler(h http.Handler) http.Handler {
	server := normalizeAddress(c.Server)
	return oauth2.Handler(h, &oauth2.Config{
		BasicAuthOff:     true,
		Client:           c.Client,
		ClientID:         c.ClientID,
		ClientSecret:     c.ClientSecret,
		AccessTokenURL:   fmt.Sprintf("%s/api/oauth/access_token_v2", server),
		AuthorizationURL: fmt.Sprintf("%s/oauth_authorize.html", server),
		Scope:            c.Scope,
		Logger:           c.Logger,
		Dumper:           c.Dumper,
	})
}

func normalizeAddress(address string) string {
	if address == "" {
		return "https://coding.net"
	}
	return strings.TrimSuffix(address, "/")
}
