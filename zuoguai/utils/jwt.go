package utils

import (
	"strconv"
	"time"

	jwt "github.com/golang-jwt/jwt"
)

var DefaultSecretKey = []byte("zuoguai_key")
var DefaultExpireDuration = 7 * 60 * 60

type MyClaims struct {
	UserId string `json:"user_id"`
	jwt.StandardClaims
}

func GenerateJwtToken(secret string, issuer string, userId int, expireDuration int) (string, error) {
	if secret == "" {
		secret = string(DefaultSecretKey)
	}
	if expireDuration == -1 {
		expireDuration = DefaultExpireDuration
	}
	token := jwt.New(jwt.SigningMethodHS256)
	nowTime := time.Now().Unix()
	token.Claims = MyClaims{
		UserId: strconv.Itoa(userId),
		StandardClaims: jwt.StandardClaims{
			NotBefore: nowTime,
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(expireDuration)).Unix(),
			Issuer:    issuer,
		},
	}
	// fmt.Println(time.Now().Add(time.Millisecond * time.Duration(expireDuration)))
	tokenString, err := token.SignedString([]byte(secret))
	return tokenString, err
}

// ParseJwtToken 解析token
func ParseJwtToken(tokenString string, secret string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	myClaims := token.Claims.(*MyClaims)

	return myClaims, nil
}
