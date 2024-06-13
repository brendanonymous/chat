package auth

import (
	"crypto/rand"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var SecretKey = generateRandomKey(10)

type UserRegistration struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
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

// GenerateJwtToken generates a JWT token for the user
func GenerateJwtToken(user_id int, username string) string {
	claims := &Claims{
		UserID:   user_id,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(SecretKey)
	if err != nil {
		panic(err)
	}

	return signedToken
}

// HashPassword hashes the plain text password
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// CheckPasswordAgainstHash checks the plain text password against a password hash
func CheckPasswordAgainstHash(password, password_hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(password_hash), []byte(password))
	if err != nil {
		return err
	}

	return nil
}
