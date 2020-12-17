package authentcation

import (
	"litshop/src/core/literr"
	"litshop/src/core/types"
	"litshop/src/module/auth"
	"litshop/src/pkg/context"
	"litshop/src/pkg/util"
	"litshop/src/request"
	"litshop/src/response"
	"litshop/src/response/resource"
)

// 常规登陆
// 手机+密码
// 邮箱+密码
// todo 手机+验证码
// todo 邮箱+验证码
func SignIn(ctx *context.Context, request *request.SignInRequest) (*response.SignInResponse, error) {
	var loginField string
	loginMethod := types.CustomerAuthTypeNull
	if util.IsValidatePhone(request.Phone) {
		loginMethod = types.CustomerAuthTypePhone
		loginField = request.Phone
	} else {
		if util.IsValidateEmail(request.Email) {
			loginMethod = types.CustomerAuthTypeEmail
			loginField = request.Email
		}
	}

	if !loginMethod.IsValid() {
		return nil, literr.NewWithCode(literr.ErrCodeInvalidRequestParams)
	}

	if !isValidatePassword(request.Password) {
		return nil, literr.NewWithCode(literr.ErrCodeInvalidRequestParams)
	}

	member, err := auth.SignWithEmailOrPhoneAndPassword(loginField, request.Password, loginMethod)
	if err != nil {
		return nil, literr.NewWithCode(literr.ErrCodeInvalidRequestParams)
	}

	res := &response.SignInResponse{
		Customer: resource.CustomerResItem(member),
	}

	return res, nil
}

func isValidatePassword(p string) bool {
	if len(p) < 8 {
		return false
	}

	return true
}
