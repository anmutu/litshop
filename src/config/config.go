package config

import (
	"fmt"
	"github.com/spf13/viper"
	"litshop/src/pkg/path"
	"os"
)

var staticConfig *viper.Viper

func init() {
	var err error
	v := viper.New()

	v.SetConfigType("yml")
	v.SetConfigName("env")
	v.AddConfigPath(path.RootPath())

	err = v.ReadInConfig()
	if err != nil {
		fmt.Printf("parse configuration file err %#v\n", err)
	}

	staticConfig = v
}

func Get() *viper.Viper {
	return staticConfig
}

func GetNormal(key string) interface{} {
	return staticConfig.Get(key)
}

func GetBool(key string) bool {
	return staticConfig.GetBool(key)
}

func GetString(key string) string {
	return staticConfig.GetString(key)
}

func GetStringSlice(key string) []string {
	return staticConfig.GetStringSlice(key)
}

func GetInt(key string) int {
	return staticConfig.GetInt(key)
}

func GetFloat64(key string) float64 {
	return staticConfig.GetFloat64(key)
}

func GetInt64(key string) int64 {
	return staticConfig.GetInt64(key)
}

func GetInt32(key string) int32 {
	return staticConfig.GetInt32(key)
}

func GetMap(key string) map[string]interface{} {
	return staticConfig.GetStringMap(key)
}

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
