package encrypt

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

/* SetPassword function uses bcrypt GenerateFromPassword to generate a encrypted version of the provided password <string> in []byte
the encrypted password is returned as hashed string */
func SetPassword(password string) string {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Fatalln(err.Error())
	}
	// fmt.Println(hashedPass)
	return string(hashedPass)
}

/* CheckPassword function uses bcrypt CompareHashAndPassword to compare the password values between user password and password in Db
return true on match or else false */
func CheckPassword(passwdFromUserDB, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwdFromUserDB), []byte(password))
	return err == nil
}
