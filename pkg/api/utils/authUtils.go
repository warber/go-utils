package api

import (
	"context"
	"net/http"
	"strings"

	"github.com/keptn/go-utils/pkg/api/models"
	v2 "github.com/keptn/go-utils/pkg/api/utils/v2"
)

type AuthV1Interface interface {
	// Authenticate authenticates the client request against the server.
	Authenticate() (*models.EventContext, *models.Error)
}

// AuthHandler handles projects
type AuthHandler struct {
	authHandler v2.AuthHandler
	BaseURL     string
	AuthToken   string
	AuthHeader  string
	HTTPClient  *http.Client
	Scheme      string
}

// NewAuthHandler returns a new AuthHandler
func NewAuthHandler(baseURL string) *AuthHandler {
	if strings.Contains(baseURL, "https://") {
		baseURL = strings.TrimPrefix(baseURL, "https://")
	} else if strings.Contains(baseURL, "http://") {
		baseURL = strings.TrimPrefix(baseURL, "http://")
	}

	httpClient := &http.Client{Transport: wrapOtelTransport(getClientTransport(nil))}

	return &AuthHandler{
		BaseURL:    baseURL,
		AuthHeader: "",
		AuthToken:  "",
		HTTPClient: httpClient,
		Scheme:     "http",

		authHandler: v2.AuthHandler{
			BaseURL:    baseURL,
			AuthHeader: "",
			AuthToken:  "",
			HTTPClient: httpClient,
			Scheme:     "http",
		},
	}
}

// NewAuthenticatedAuthHandler returns a new AuthHandler that authenticates at the endpoint via the provided token
// Deprecated: use APISet instead
func NewAuthenticatedAuthHandler(baseURL string, authToken string, authHeader string, httpClient *http.Client, scheme string) *AuthHandler {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	httpClient.Transport = wrapOtelTransport(getClientTransport(httpClient.Transport))

	return createAuthenticatedAuthHandler(baseURL, authToken, authHeader, httpClient, scheme)
}

func createAuthenticatedAuthHandler(baseURL string, authToken string, authHeader string, httpClient *http.Client, scheme string) *AuthHandler {
	baseURL = strings.TrimPrefix(baseURL, "http://")
	baseURL = strings.TrimPrefix(baseURL, "https://")
	return &AuthHandler{
		BaseURL:    baseURL,
		AuthHeader: authHeader,
		AuthToken:  authToken,
		HTTPClient: httpClient,
		Scheme:     scheme,

		authHandler: v2.AuthHandler{
			BaseURL:    baseURL,
			AuthHeader: authHeader,
			AuthToken:  authToken,
			HTTPClient: httpClient,
			Scheme:     scheme,
		},
	}
}

func (a *AuthHandler) getBaseURL() string {
	return a.BaseURL
}

func (a *AuthHandler) getAuthToken() string {
	return a.AuthToken
}

func (a *AuthHandler) getAuthHeader() string {
	return a.AuthHeader
}

func (a *AuthHandler) getHTTPClient() *http.Client {
	return a.HTTPClient
}

// Authenticate authenticates the client request against the server.
func (a *AuthHandler) Authenticate() (*models.EventContext, *models.Error) {
	return a.authHandler.Authenticate(context.TODO(), v2.AuthAuthenticateOptions{})
}
