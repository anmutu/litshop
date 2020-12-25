package customer

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"litshop/src/lvm/types"
	"litshop/src/model"
)

var (
	repo = model.NewCustomerRepository()
)

func IsPhoneExists(phone string) bool {
	return checkLoginIsExists(types.CustomerAuthTypePhone, "auth_id", phone)
}

func IsWechatOpenidExists(openid string) bool {
	return checkLoginIsExists(types.CustomerAuthTypeWechat, "auth_id", openid)
}

func IsWechatUnionIdExists(unionId string) bool {
	return checkLoginIsExists(types.CustomerAuthTypeWechat, "auth_value", unionId)
}

func IsEmailExists(email string) bool {
	return checkLoginIsExists(types.CustomerAuthTypeEmail, "auth_id", email)
}

func checkLoginIsExists(authType types.CustomerAuthType, field string, value string) bool {
	login := &model.CustomerLogin{}
	result := repo.DB.Where(fmt.Sprintf("auth_type = %s and %s = ?", authType.String(), field), value).First(login)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false
	}

	return true
}

func Create(customer *model.Customer, customLogin *model.CustomerLogin) error {
	err := repo.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(customer).Error; err != nil {
			return err
		}

		customLogin.CustomerId = customer.ID
		if err := tx.Create(customLogin).Error; err != nil {
			return err
		}

		return nil
	})

	return err
}

func CreateFromWxUser() {

}

func GetByField() {

}

func GetById() {

}

func GetByPhone(phone string) (*model.Customer, error) {
	return nil, nil
}

func GetByEmail(email string) (*model.Customer, error) {
	return nil, nil
}

func HashAndSaltPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func CompareHashAndPassword(hashedPassword, password []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)
	if err != nil {
		return false
	}

	return true
}
