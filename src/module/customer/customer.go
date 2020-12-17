package customer

import (
	"golang.org/x/crypto/bcrypt"
	"litshop/src/model"
)

func IsPhoneExists(phone string) bool {
	return false
}

func IsWechatOpenidExists() bool {
	return false
}

func IsWechatUnionIdExists() bool {
	return false
}

func IsEmailExists(email string) bool {
	return false
}

func Create(customer *model.Customer) error {
	return nil
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
