package auth

import (
	"card2go_service/config"

	"github.com/golang-jwt/jwt/v5"
)

type IDClaim struct {
	ID uint `json:"id"`
	jwt.RegisteredClaims
}

func CreateSignedToken(id uint) string {
	str, _ := jwt.NewWithClaims(jwt.SigningMethodES256, IDClaim{
		ID: id,
	}).SignedString(config.TokenKey)
	return str
}

func GetIDFromToken(token string) (uint, error) {
	// i'm not sure how to use keyfunc
	// TODO: learn about the vulnerabilities of this function
	t, err := jwt.ParseWithClaims(token, &IDClaim{}, func(t *jwt.Token) (interface{}, error) { return config.TokenKey, nil })

	if claims, ok := t.Claims.(*IDClaim); ok && t.Valid {
		return claims.ID, nil
	}
	return 0, err
}
