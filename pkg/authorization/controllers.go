package authorization

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"gomessage/pkg/utils/log/loggers"
)

func JwtKey() []byte {
	return []byte(viper.GetString("auth.jwtKey"))
}

type Account struct {
	*Users
	Password string `json:"password"` //Users表没有设计Password字段，是为了更方便的处理密码序列化与反序列化
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func HashAndSalt(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		loggers.DefaultLogger.Error(err)
	}
	return string(hash)
}

func ComparePassword(pwdHash, pwd string) bool {
	byteHash := []byte(pwdHash)
	bytePwd := []byte(pwd)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePwd)
	if err != nil {
		return false
	}
	return true
}
