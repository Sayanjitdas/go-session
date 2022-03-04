package user

import (
	"errors"
	"fmt"
	"go-session/encrypt"

	"github.com/google/uuid"
)

//User struct
type user struct {
	id        string
	FirstName string
	LastName  string
	Email     string
	password  string
}

//DBUser Map
var dbUser = make(map[string]user)

//HasUser function checks for already existing user
func HasUser(email string) bool {
	for _, usr := range dbUser {
		if usr.Email == email {
			return true
		}
	}
	return false
}

//CreateUser function creates a new user and returns the updated user
func (u *user) CreateUser() error {
	if HasUser(u.Email) {
		return fmt.Errorf("user with %s already exists", u.Email)
	}
	uid := generateUUID()
	u.id = uid
	dbUser[uid] = *u
	return nil
}

//GetUserId is a getter method for returning user id of a specific user
func (u user) GetUserId() string {
	return u.id
}

//SetPassword set the encrytion of the stringed password into hashed bytes password
func (u *user) SetPassword(password string) {
	u.password = encrypt.SetPassword(password)
}

//CheckPassword validates the passwords between user and db
func (u *user) CheckPassword(password string) bool {
	return encrypt.CheckPassword(u.password, password)

}

//New function creates a new user instance
func New() user {
	var u user
	return u
}

//GetAllUser function returns all user in dbUser
func GetAllUser() map[string]user {
	db := dbUser
	return db
}

//GetSpecificUser function returns a specific user to uid
func GetSpecificUser(uid string) (*user, error) {
	if val, ok := dbUser[uid]; ok {
		return &val, nil
	}
	return nil, errors.New("key does not exists")
}

//GetUserByEmail function returns a specific user to email
func GetUserByEmail(email string) (*user, error) {

	for _, usr := range dbUser {
		if usr.Email == email {
			return &usr, nil
		}
	}
	return nil, fmt.Errorf("user with %s not found", email)
}

//checkUUID function check for unique uuid generation
func checkUUID(UUid string, data map[string]user) bool {
	if _, ok := data[UUid]; ok {
		return true
	}
	return false
}

//generateUUID function generates unique id for each user
func generateUUID() string {
	uid := uuid.NewString()
	if !checkUUID(uid, dbUser) {
		return uid
	}
	return generateUUID()
}
