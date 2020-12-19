package request

import (
	"litshop/src/lvm/types"
)

type SignInRequest struct {
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Otp      string `json:"otp"`
}

type SignUpRequest struct {
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	RePassword string `json:"re_password"`
	Otp        string `json:"otp"`
}

type SendOtpRequest struct {
	Phone  string          `json:"phone"`
	Email  string          `json:"email"`
	Action types.OtpAction `json:"action"`
}
