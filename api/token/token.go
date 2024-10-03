package token

import (
	"api_gateway/internal/config"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

func ValidateToken(tokenstr string) (bool, error) {
	_, err := ExtractClaims(tokenstr)
	if err != nil {
		return false, err
	}
	return true, nil
}

func ExtractClaims(tokenstr string) (jwt.MapClaims, error) {
	token, err := jwt.ParseWithClaims(tokenstr, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(config.Load().SIGNING_KEY), nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token: %s", tokenstr)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("failed to parse token claims")
	}
	return claims, nil
}

func GetUserInfoFromAccessToken(accessTokenString string) (string, string, string, string, error) {
	refreshToken, err := jwt.Parse(accessTokenString, func(token *jwt.Token) (interface{}, error) { return []byte(config.Load().SIGNING_KEY), nil })
	if err != nil || !refreshToken.Valid {
		return "", "", "", "", err
	}
	claims, ok := refreshToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", "", "", "", err
	}
	userID := claims["user_id"].(string)
	email := claims["email"].(string)
	password := claims["password"].(string)
	role := claims["role"].(string)

	return userID, email, password, role, nil
}
