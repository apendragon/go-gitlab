//
// Copyright 2020, Thomas Cazali
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package gitlab

import (
	"fmt"
	"time"
)

// PersonalAccessTokensService handles communication with the
// personal_access_tokens related methods of the GitLab API.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/personal_access_tokens.html
type PersonalAccessTokensService struct {
	client *Client
}

// PersonalAccessToken represents a GitLab user's Personal Access Token.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/personal_access_tokens.html
type PersonalAccessToken struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	Revoked   bool       `json:"key"`
	CreatedAt *time.Time `json:"created_at"`
	Scopes    []string   `json:"scopes"`
	Active    bool       `json:"active"`
	UserID    int        `json:"user_id"`
	ExpiresAt *time.Time `json:"expires_at"`
}

func (k PersonalAccessToken) String() string {
	return Stringify(k)
}

// ListPersonalAccessTokensOptions represents the available ListPersonalAccessTokens() options.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/personal_access_tokens.html#list-personal-access-tokens
type ListPersonalAccessTokensOptions struct {
	ListOptions
	UserID *int `url:"user_id,omitempty" json:"user_id,omitempty"`
}

// ListPersonalAccessTokens gets a list of personal access tokens.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/personal_access_tokens.html#list-personal-access-tokens
func (s *PersonalAccessTokensService) ListPersonalAccessTokens(opt *ListPersonalAccessTokensOptions, options ...RequestOptionFunc) ([]*PersonalAccessToken, *Response, error) {

	req, err := s.client.NewRequest("GET", "personal_access_tokens", opt, options)
	if err != nil {
		return nil, nil, err
	}

	var pats []*PersonalAccessToken
	resp, err := s.client.Do(req, &pats)
	if err != nil {
		return nil, resp, err
	}

	return pats, resp, err
}

// RevokePersonalAccessToken revokes a personal access token.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/personal_access_tokens.html#revoke-a-personal-access-token
func (s *PersonalAccessTokensService) RevokePersonalAccessToken(ID int, options ...RequestOptionFunc) (*Response, error) {
	p := fmt.Sprintf("personal_access_tokens/%d", ID)

	req, err := s.client.NewRequest("DELETE", p, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
