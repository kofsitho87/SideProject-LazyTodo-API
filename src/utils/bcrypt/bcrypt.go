package bcrypt

import "golang.org/x/crypto/bcrypt"

func Generate(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func Compare(hash, password string) (bool, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		//err types
		// bcrypt.ErrMismatchedHashAndPassword: wrong password
		// bcrypt.ErrHashTooShort: ?
		// TODO: err를 wrap 하여 상세를 전달하면 좋다
		return false, err
	}

	return true, nil
}
