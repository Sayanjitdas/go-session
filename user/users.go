package user

import (
	"errors"

	"github.com/google/uuid"
)

//User struct
type user struct {
	id        string
	FirstName string
	LastName  string
	Email     string
	Password  string
}

//CreateUser function creates a new user and returns the updated user
func (u *user) CreateUser() {
	uid := generateUUID()
	u.id = uid
	dbUser[uid] = *u
}

//GetUserId is a getter method for returning user id of a specific user
func (u user) GetUserId() string {
	return u.id
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

//DBUser Map
var dbUser = make(map[string]user)

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
