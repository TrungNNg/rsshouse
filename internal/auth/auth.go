package auth

import "golang.org/x/crypto/bcrypt"

// Hash the password with bcrypt
func HashPassword(password string) (string, error) {
	data, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// check if the given password when login and the hashed in db is match
func CheckPasswordHash(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
