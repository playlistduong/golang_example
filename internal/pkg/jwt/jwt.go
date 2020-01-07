package jwt

import (
	"crypto/rsa"
	"errors"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

type (
	Signer interface {
		Sign(data Data) (string, error)
	}

	Verifier interface {
		Verify(tokenString string) (*Data, error)
	}

	Config struct {
		JWTSecret []byte `envconfig:"JWT_SECRET"`
	}

	Data struct {
		jwt.StandardClaims
		Username string
		Email    string
		Role     int16
		Avatar   string
	}
)

// Sign generator JWT token
func (conf *Config) Sign(data Data) (string, error) {
	signKey, _ = jwt.ParseRSAPrivateKeyFromPEM(conf.JWTSecret)
	token := jwt.New(jwt.GetSigningMethod("RS256"))
	token.Claims = data
	v, err := token.SignedString(signKey)
	return v, err
}

// Verify if token invalid
func (conf *Config) Verify(tokenString string) (*Data, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Data{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(conf.JWTSecret), nil
	})

	if err != nil {
		return nil, err
	}

	data, ok := token.Claims.(*Data)

	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return data, nil
}
