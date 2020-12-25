package config

import (
	"fmt"
	"github.com/spf13/viper"
	"litshop/src/pkg/path"
	"os"
)

var (
	err error
	v   *viper.Viper
)

func init() {
	v = viper.New()
	v.SetConfigType("yml")
	v.SetConfigName("app")

	p := fmt.Sprintf("%s/%s", path.RootPath(), "configs")
	fmt.Println("conf path", p)
	v.AddConfigPath(p)

	err = v.ReadInConfig()
	if err != nil {
		fmt.Printf("parse configuration file err %#v\n", err)
		os.Exit(1)
	}

	if v.GetBool("watch") {
		v.WatchConfig()
	}
}

func GetNormal(key string) interface{} {
	return v.Get(key)
}

func GetBool(key string) bool {
	return v.GetBool(key)
}

func GetString(key string) string {
	return v.GetString(key)
}

func GetStringSlice(key string) []string {
	return v.GetStringSlice(key)
}

func GetInt(key string) int {
	return v.GetInt(key)
}

func GetFloat64(key string) float64 {
	return v.GetFloat64(key)
}

func GetInt64(key string) int64 {
	return v.GetInt64(key)
}

func GetInt32(key string) int32 {
	return v.GetInt32(key)
}

func GetMap(key string) map[string]interface{} {
	return v.GetStringMap(key)
}

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
