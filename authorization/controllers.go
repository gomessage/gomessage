package authorization

import (
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gomessage/pkg/log/loggers"
)

var JwtKey = []byte("syx8F6FtdWKS9yj0UGcnKGaTq2MsWOilqwB8uELRCGOMlYJgl4t1zxSsOwiF7sAS")

type Account struct {
	*Users
	Password string `json:"password"` //Users表没有设计Password字段，是为了更方便的处理密码序列化与反序列化
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
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
