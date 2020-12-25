package util

import (
	"encoding/json"
	"fmt"
)

func IsValidatePhone(s string) bool {
	return false
}

func IsValidateEmail(s string) bool {
	return false
}

func Struct2Struct(a, b interface{}) bool {
	s, err := json.Marshal(a)
	if err == nil {
		err := json.Unmarshal(s, b)
		if err == nil {
			return true
		}
	}

	return false
}

func Interface2String(inter interface{}) string {
	switch inter.(type) {
	case struct{}:
		b, err := json.Marshal(inter)
		if err == nil {
			return string(b)
		}
	case string:
		return fmt.Sprintf("%s", inter.(string))
	case int:
		return fmt.Sprintf("%d", inter.(int))
	case float64:
		return fmt.Sprintf("%f", inter.(float64))
	}

	return ""
}
