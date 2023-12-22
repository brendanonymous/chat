package auth_handlers

import (
	"crypto/rand"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type UserRegistration struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthenticatedUser struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

// Claims represents the JWT claims
type Claims struct {
	UserID   int    `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func generateRandomKey(length int) []byte {
	// Create a byte slice to store the random key.
	key := make([]byte, length)

	// Generate random bytes for the key.
	rand.Read(key)

	// Convert the byte slice to a string.
	return key
}

// generateJwtToken generates a JWT token for the user
func generateJwtToken(user_id int, username string) string {
	claims := &Claims{
		UserID:   user_id,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		panic(err)
	}

	return signedToken
}

// hashPassword hashes the plain text password
func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
