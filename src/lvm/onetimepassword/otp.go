package onetimepassword

import "litshop/src/lvm/types"

func GetCode(action types.OtpAction) string {
	return ""
}

func VerifyCode(Code string, action types.OtpAction) bool {
	return false
}

type OtpDriver interface {
	Store()
	Load()
	Delete()
}
