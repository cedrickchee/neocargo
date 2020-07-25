// Service layer housing services like authentication and more.

package main

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	pb "github.com/cedrickchee/neocargo/neocargo-service-user/proto/user"
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
func (srv *TokenService) Decode(tokenString string) (*CustomClaims, error) {

	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	// Validate the token and return the custom claims
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}

// Encode a claim into a JWT
func (srv *TokenService) Encode(user *pb.User) (string, error) {
	expireToken := time.Now().Add(time.Hour * 72).Unix()

	// Create the Claims
	claims := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    "neocargo.service.user",
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token and return
	return token.SignedString(key)
}
