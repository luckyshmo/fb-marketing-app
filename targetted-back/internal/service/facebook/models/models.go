package models

type Credentials struct {
	Token      string
	ApiVersion string
	BusinessID string
	// Probably would choose from several ad accounts in the future
	AdAccountID string
}
