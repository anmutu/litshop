package response

import "litshop/src/pkg/d"

type SignUpResponse struct {
}

type SignInResponse struct {
	Customer    d.H    `json:"customer"`
	IdToken     string `json:"id_token"`
	AccessToken string `json:"access_token"`
}
