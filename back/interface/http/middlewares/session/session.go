// session/store.go
package session

import "github.com/gorilla/sessions"

var Store = sessions.NewCookieStore([]byte("troque-este-secret"))
