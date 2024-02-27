package ApplicationLogic

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (hashPassword string) {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 5)

	return string(bytes)
}

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
