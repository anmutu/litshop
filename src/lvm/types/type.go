package types

import "fmt"

// 一次性密码
type OtpAction = int

const (
	OtpActionDefault OtpAction = iota
	OtpActionSignUp
	OtpActionSignIn
	OtpActionResetPassword
	OtpActionVerifyEmail
	OtpActionVerifyPhone
	OtpActionBindPhone
	OtpActionBindEmail
)

// 通用的状态
type ComStatus = int

const (
	ComStatusNormal   ComStatus = 1
	ComStatusDisabled ComStatus = 2
	ComStatusDeleted  ComStatus = -9
)

// 用户状态
type CustomerGender = int
type CustomerStatus = int
type CustomerLoginStatus = int
type CustomerAuthType int

func (at *CustomerAuthType) String() string {
	return fmt.Sprintf("%d", *at)
}

func (at CustomerAuthType) IsValid() bool {
	if at > 0 && at <= CustomerAuthTypeMax {
		return true
	}

	return false
}

const (
	CustomerGenderUnknown CustomerGender = iota
	CustomerGenderFemale
	CustomerGenderMale
)

const (
	CustomerStatusNormal   = 1
	CustomerStatusDisabled = 2
	CustomerStatusDeleted  = -9
)
const (
	CustomerAuthTypeNull CustomerAuthType = iota
	CustomerAuthTypeUsername
	CustomerAuthTypeEmail
	CustomerAuthTypePhone
	CustomerAuthTypeWechat
	CustomerAuthTypeGoogle
	CustomerAuthTypeFacebook
	CustomerAuthTypeAppleId
	CustomerAuthTypeWeibo
	CustomerAuthTypeQQ
	CustomerAuthTypeMax
)
const (
	CustomerLoginStatusNormal   = 1
	CustomerLoginStatusDisabled = 2
	CustomerLoginStatusDeleted  = -9
)

// 属性状态
type AttributeType = int

// product
type ProductType int

const (
	ProductItem        ProductType = 1
	ProductVirtualItem ProductType = 2
)
