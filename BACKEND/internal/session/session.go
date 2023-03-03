package session

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("secret-key"))

func SetUserID(session *sessions.Session, userID int) {
	session.Values["userID"] = userID
}

func GetUserID(session *sessions.Session) int {
	userID, ok := session.Values["userID"].(int)
	if !ok {
		return 0
	}
	return userID
}

func ClearSession(session *sessions.Session) {
	session.Values["userID"] = nil
}

func GetUserSession(r *http.Request) *sessions.Session {
	session, _ := store.Get(r, "session-name")
	return session
}

func SaveSession(w http.ResponseWriter, r *http.Request, session *sessions.Session) {
	session.Save(r, w)
}
