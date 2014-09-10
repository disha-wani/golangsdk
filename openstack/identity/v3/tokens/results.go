package tokens

import (
	"time"

	"github.com/mitchellh/mapstructure"
)

// RFC3339Milli describes the time format used by identity API responses.
const RFC3339Milli = "2006-01-02T15:04:05.999999Z"

// TokenCreateResult contains the document structure returned from a Create call.
type TokenCreateResult struct {
	response map[string]interface{}
	tokenID  string
}

// TokenID retrieves a token generated by a Create call from an token creation response.
func (r *TokenCreateResult) TokenID() (string, error) {
	return r.tokenID, nil
}

// ExpiresAt retrieves the token expiration time.
func (r *TokenCreateResult) ExpiresAt() (time.Time, error) {
	type tokenResp struct {
		ExpiresAt string `mapstructure:"expires_at"`
	}

	type response struct {
		Token tokenResp `mapstructure:"token"`
	}

	var resp response
	err := mapstructure.Decode(r.response, &resp)
	if err != nil {
		return time.Time{}, err
	}

	// Attempt to parse the timestamp.
	ts, err := time.Parse(RFC3339Milli, resp.Token.ExpiresAt)
	if err != nil {
		return time.Time{}, err
	}

	return ts, nil
}
