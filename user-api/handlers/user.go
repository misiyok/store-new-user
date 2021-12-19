package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/miracle-org/store-new-user/user-api/data"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	l  *log.Logger
	mc *mongo.Client
}

func NewUser(l *log.Logger, mc *mongo.Client) *User {
	return &User{l: l, mc: mc}
}

type KeyUser struct{}

func (u User) MiddlewareUserValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		user := data.User{}

		err := user.FromJSON(r.Body)
		if err != nil {
			u.l.Println("[ERROR] deserializing user", err)
			http.Error(rw, "Errror reading user", http.StatusBadRequest)
			return
		}

		// validate the user
		err = user.Validate()
		if err != nil {
			u.l.Println("[ERROR] validating user", err)
			http.Error(
				rw,
				fmt.Sprintf("Error validating user: %s", err),
				http.StatusBadRequest,
			)
			return
		}

		ctx := context.WithValue(r.Context(), KeyUser{}, user)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})
}
