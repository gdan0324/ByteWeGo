package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GnerateToken(userId string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Hour * 24 * 7)
	issuer := "ByteWeGo"

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Id":         userId,
		"ExpitresAt": expireTime,
		"Issuer":     issuer,
	})

	return token.SignedString([]byte("golang"))
}

func ParseToken(token string) (jwt.MapClaims, error) {
	tokenClaims, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte("golang"), nil
	})

	if err != nil {
		return nil, err
	}
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(jwt.MapClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
