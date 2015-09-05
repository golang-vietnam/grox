package domain

type AccountType string

const (
	AccountTypeGithub = AccountType("github")
)

type Account struct {
	Type      AccountType `json:"type" rethink:"type"`
	AuthToken string      `json:"auth_token" rethink:"auth_token"`
}
