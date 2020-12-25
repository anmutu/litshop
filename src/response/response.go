package response

import "litshop/src/pkg/d"

type SignUpResponse struct {
}

type SignInResponse struct {
	Customer    d.H    `json:"customer"`
	AccessToken string `json:"access_token"`
}
