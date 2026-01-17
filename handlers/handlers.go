package handlers

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var Keyfunc = func(token *jwt.Token) (any, error) {
	return []byte(""), nil
}

func GenerateToken(w http.ResponseWriter, userID string) string {
	key := ""
	t := jwt.New(jwt.SigningMethodHS256)

	t.Claims = jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	tokenString, err := t.SignedString([]byte(key))
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return ""
	}
	return tokenString
}

func VerifyToken(token string, w http.ResponseWriter) string {
	parsedtoken, err := jwt.Parse(token, Keyfunc)
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return ""
	}
	mapClaims := parsedtoken.Claims.(jwt.MapClaims)
	userID := mapClaims["user_id"].(string)
	return userID
}

func HashPassword(password string, w http.ResponseWriter) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return ""
	}
	return string(hashedPassword)
}

func VerifyPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
