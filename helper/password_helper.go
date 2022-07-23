package helper

import "golang.org/x/crypto/bcrypt"

func BcryptPassword(password string) (string, error) {
	bytePassword := []byte(password)
	byteHashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	hashedPassword := string(byteHashedPassword)
	return hashedPassword, nil
}
