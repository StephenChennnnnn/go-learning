package util

import (
	"github.com/dgrijalva/jwt-go"
	"go-learning/gin/gin-example/pkg/setting"
)

var jwtSecret = []byte(setting.JwtSecret)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

//func GenerateToken(username, password string) (string, error) {
//	nowTime := time.Now()
//	expireTime := nowTime.Add(3 * time.Hour)
//}
