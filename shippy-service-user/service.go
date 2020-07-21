// Service layer housing services like authentication and more.

package main

import (
	"github.com/dgrijalva/jwt-go"
	pb "github.com/haxorbit/shippy/shippy-service-user/proto/user"
)

var (
	// Define a secure key string used as a salt when hashing our tokens.
	// Please make your own way more secure than this, use a randomly generated
	// SHA hash or something.
	key = []byte("sup3rsecu9eKEY")
)

// CustomClaims is our custom metadata, which will be hashed and sent as the
// second segment in our JWT.
type CustomClaims struct {
	User *pb.User
	jwt.StandardClaims
}

// TokenService is a token encoding/decoding service.
type TokenService struct {
	repo Repository
}

// Decode a token string into a token object.
func (srv *TokenService) Decode(token string) (*CustomClaims, error) {

	// Parse the token
	tokenType, err := jwt.ParseWithClaims(string(key), &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	// Validate the token and return the custom claims
	if claims, ok := tokenType.Claims.(*CustomClaims); ok && tokenType.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

// Encode a claim into a JWT
func (srv *TokenService) Encode(user *pb.User) (string, error) {
	// Create the Claims
	claims := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "shippy.service.user",
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token and return
	return token.SignedString(key)
}
