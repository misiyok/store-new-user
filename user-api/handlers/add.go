package handlers

import (
	"net/http"

	"github.com/miracle-org/store-new-user/user-api/data"
)

func (u *User) AddUser(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("Handle POST User")

	user := r.Context().Value(KeyUser{}).(data.User)
	data.AddUser(&user, u.mc)
}
