package gitlab

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestListPersonalAccessTokens(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/personal_access_tokens",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			fmt.Fprint(w, `
[
    {
        "id": 4,
        "name": "Test Token",
        "revoked": false,
        "created_at": "2020-07-23T14:31:47.000Z",
        "scopes": [
            "api"
        ],
        "active": true,
        "user_id": 24,
        "expires_at": null
    }
]
`)
		},
	)

	opt := &ListPersonalAccessTokensOptions{
		UserID: Int(4),
	}

	pats, _, err := client.PersonalAccessTokens.ListPersonalAccessTokens(opt)
	if err != nil {
		t.Errorf("PersonalAccessTokens.ListPersonalAccessTokens returned error: %v", err)
	}

	wantExpiresAt := time.Date(2020, 07, 23, 14, 31, 47, 0, time.UTC)

	want := []*PersonalAccessToken{
		{
			ID:        4,
			Name:      "Test Token",
			Revoked:   false,
			CreatedAt: &wantExpiresAt,
			Scopes:    []string{"api"},
			Active:    true,
			UserID:    24,
			ExpiresAt: nil,
		},
	}

	if !reflect.DeepEqual(want, pats) {
		t.Errorf("PersonalAccessTokens.ListPersonalAccessTokens returned %+v, want %+v", pats, want)
	}
}

func TestRevokePersonalAccessToken(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)
	mux.HandleFunc("/api/v4/personal_access_tokens/4",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "DELETE")
		},
	)
	_, err := client.PersonalAccessTokens.RevokePersonalAccessToken(4)
	if err != nil {
		t.Error(err)
	}
}
