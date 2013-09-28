package octokat

import (
	"time"
)

type Authorization struct {
	ID        int       `json:"id,omitempty"`
	URL       string    `json:"url,omitempty"`
	App       App       `json:"app,omitempty"`
	Token     string    `json:"token,omitempty"`
	Note      string    `json:"note,omitempty"`
	NoteURL   string    `json:"note_url,omitempty"`
	Scopes    []string  `json:"scopes,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type App struct {
	ClientID string `json:"client_id,omitempty"`
	URL      string `json:"url,omitempty"`
	Name     string `json:"name,omitempty"`
}

type AuthorizationParams struct {
	Scopes       []string `json:"scopes,omitempty"`
	Note         string   `json:"note,omitempty"`
	NoteUrl      string   `json:"note_url,omitempty"`
	ClientID     string   `json:"client_id,omitempty"`
	ClientSecret string   `json:"client_secret,omitempty"`
}

// List the authenticated user's authorizations
//
// API for users to manage their own tokens.
// You can only access your own tokens, and only through
// Basic Authentication.
//
// See http://developer.github.com/v3/oauth/#list-your-authorizations
func (c *Client) Authorizations(options *Options) ([]Authorization, error) {
	var auths []Authorization
	err := c.jsonGet("authorizations", options, &auths)
	if err != nil {
		return nil, err
	}

	return auths, nil
}

// Create an authorization for the authenticated user.
//
// You can create your own tokens, and only through
// Basic Authentication.
//
// See http://developer.github.com/v3/oauth/#scopes Available scopes
// See http://developer.github.com/v3/oauth/#create-a-new-authorization
func (c *Client) CreateAuthorization(options *Options) (*Authorization, error) {
	var auth Authorization
	err := c.jsonPost("authorizations", options, options.Params, &auth)
	if err != nil {
		return nil, err
	}

	return &auth, nil
}
