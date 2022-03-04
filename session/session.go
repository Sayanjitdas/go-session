package session

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

//Session DB
var sessionDB = make(map[string]sessionUser)

//session user struct
type sessionUser struct {
	sessionId string
	userId    string
}

//HasSession
func HasSession(r *http.Request) bool {
	sid, _ := r.Cookie("session")
	if _, ok := sessionDB[sid.Value]; ok {
		return true
	}
	return false
}

//SetSession
func SetSession(w *http.ResponseWriter, userid string) {
	sid := uuid.NewString()
	sessionDB[sid] = sessionUser{
		sessionId: sid,
		userId:    userid,
	}
	http.SetCookie(*w, &http.Cookie{
		Name:  "session",
		Value: sid,
	})
}

//GetSessionUser returns the userId tagged to active provided sessionId
func GetSessionUser(sid string) (string, error) {
	if su, ok := sessionDB[sid]; ok {
		return su.userId, nil
	}
	return "", fmt.Errorf("no user with the session id %v", sid)
}

//DeleteSession
func DeleteSession(w *http.ResponseWriter, r *http.Request) {
	sid, _ := r.Cookie("session")
	delete(sessionDB, sid.Value)
	http.SetCookie(*w, &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1, //this means session cookie will expire immediately
	})
}
