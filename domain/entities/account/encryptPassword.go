package account

import "golang.org/x/crypto/bcrypt"

func GenerateHashPassword(secret string) (string, error) {

	hashSecret, err := bcrypt.GenerateFromPassword([]byte(secret), 14)

	if err != nil {
		return "", err
	}

	return string(hashSecret), nil
}

func VerifyPasswordHash(secret string, secretSaved string) error {
	err := bcrypt.CompareHashAndPassword([]byte(secret), []byte(secretSaved))

	return err
}
