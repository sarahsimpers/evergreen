// Code generated by rest/model/codegen.go. DO NOT EDIT.

package model

import "github.com/mongodb/grip/send"

type APIGithubOptions struct {
	Account string `json:"account"`
	Repo    string `json:"repo"`
}

func (m *APIGithubOptions) BuildFromService(t send.GithubOptions) error {
	m.Account = StringString(t.Account)
	m.Repo = StringString(t.Repo)
	return nil
}

func (m *APIGithubOptions) ToService() (send.GithubOptions, error) {
	out := send.GithubOptions{}
	out.Account = StringString(m.Account)
	out.Repo = StringString(m.Repo)
	return out, nil
}
