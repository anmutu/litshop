package auth

import (
	"errors"
	"litshop/src/lvm/types"
	"litshop/src/model"
	"litshop/src/module/customer"
)

func SignWithEmailOrPhoneAndPassword(field string, password string, authType types.CustomerAuthType) (*model.Customer, error) {
	exists := false
	var member *model.Customer
	var err error

	if authType == types.CustomerAuthTypePhone {
		exists = customer.IsPhoneExists(field)
		if exists {
			member, err = customer.GetByEmail(field)
		}
	}

	if authType == types.CustomerAuthTypeEmail {
		exists = customer.IsEmailExists(field)
		if exists {
			member, err = customer.GetByEmail(field)
		}
	}

	if !exists {
		return nil, errors.New("customer not found")
	}

	if err != nil {
		return nil, err
	}

	if member == nil {
		return nil, errors.New("customer not found")
	}

	if !customer.CompareHashAndPassword([]byte(member.Password), []byte(password)) {
		return nil, errors.New("customer not found")
	}

	return member, nil
}

func SignInByPhoneAndOtp() {

}

func SignInByEmailAndOtp() {

}

func Logout() {

}
