package encrypt_test

import (
	"go-session/encrypt"
	"testing"
)

func TestSetPassword(t *testing.T) {
	hashedPass := encrypt.SetPassword("abcd")
	if len(hashedPass) == 0 {
		t.Error("Hashed password is nil")
	}
}

func TestCheckPassword(t *testing.T) {
	UserPass := "abcd"
	DBPass := encrypt.SetPassword(UserPass)
	if !encrypt.CheckPassword(DBPass, UserPass) {
		t.Error("Password do not match!!")
	}
}
