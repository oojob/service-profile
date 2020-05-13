package db

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/oojob/service-profile/src/model"
	"github.com/spf13/viper"
)

var (
	key = []byte(viper.GetString("jwtsecret"))
)

// CustomClaims :- claims
type CustomClaims struct {
	Profile *model.Profile
	jwt.StandardClaims
}

// Authable :- interface for auth
type Authable interface {
	Decode(token string) (*CustomClaims, error)
	Encode(profile *model.Profile) (string, error)
}

// Encode a claim into a JWT
func (db *Database) Encode(profile *model.Profile) (string, error) {

	expireToken := time.Now().Add(time.Hour * 72).Unix()

	// Create the Claims
	claims := CustomClaims{
		profile,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    "account.service",
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token and return
	return token.SignedString(key)
}

// Decode a token string into a token object
func (db *Database) Decode(tokenString string) (*CustomClaims, error) {

	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	// Validate the token and return the custom claims
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
