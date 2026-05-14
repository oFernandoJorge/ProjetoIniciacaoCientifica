package auth

import "golang.org/x/crypto/bcrypt"

// HashPassword gera hash seguro da senha
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	return string(bytes), err
}

// CheckPassword compara senha e hash
func CheckPassword(password string, hash string) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(hash),
		[]byte(password),
	)
}
