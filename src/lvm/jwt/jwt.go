package jwt

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/uniplaces/carbon"
	"io/ioutil"
	"litshop/src/config"
	"litshop/src/config/redis"
	"litshop/src/lvm/literr"
	"litshop/src/lvm/types"
	"litshop/src/pkg/logger"
	"litshop/src/pkg/path"
	"litshop/src/pkg/sn"
	"litshop/src/pkg/util"
	"time"
)

var secret []byte
var rds = redis.Get()

const (
	cachePrefix = "litshop:jwt:"
)

func init() {
	content, err := ioutil.ReadFile(path.RootPathWithPostfix(config.GetString("jwt.pubkey")))
	if err != nil {
		logger.Fatal("读取jwt pubkey 失败")
	}

	secret = content
}

func parseToken(token string) (*jwt.StandardClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return secret, nil
	})

	if err == nil && jwtToken != nil {
		if claim, ok := jwtToken.Claims.(*jwt.StandardClaims); ok && jwtToken.Valid {
			return claim, nil
		}
	}

	return nil, err
}

func ParseToken(token string) (*jwt.StandardClaims, error) {
	claim, err := parseToken(token)
	if err != nil {
		return nil, err
	}

	exists, err := rds.Exists(context.TODO(), cachePrefix+claim.Id).Result()
	if exists == 0 {
		return nil, literr.NewWithCode(literr.ErrCodeAuthError)
	}

	return claim, nil
}

func GenJwtToken(c types.HasTokener) (string, error) {
	expiresTime := carbon.Now().AddDays(7).Unix()
	tokenId := sn.UuidString()

	claims := &jwt.StandardClaims{
		Audience:  util.Interface2String(c.GetId()), // 受众
		ExpiresAt: expiresTime,                      // 失效时间
		Id:        util.Interface2String(tokenId),   // 编号
		IssuedAt:  time.Now().Unix(),                // 签发时间
		Issuer:    "litshop api",                    // 签发人
		NotBefore: time.Now().Unix(),                // 生效时间
		Subject:   "customer jwt token",             // 主题
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secret)

	if err != nil {
		return "", err
	}

	go CacheToken(tokenId, tokenString, time.Hour*time.Duration(24*7)-time.Hour)
	return tokenString, nil
}

func CacheToken(tokenId, token string, ttl time.Duration) {
	rds.SetXX(context.TODO(), cachePrefix+tokenId, token, ttl)
}
