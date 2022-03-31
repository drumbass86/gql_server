package auth

import (
	"context"
	"gql_serv/db"
	"gql_serv/pkg/jwt"
	"net/http"
)

var ctxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")
			if header == "" {
				next.ServeHTTP(w, r)
				return
			}

			tokenStr := header
			username, err := jwt.ParseToken(tokenStr)
			if err != nil {
				http.Error(w, "Invalid token in header Authorization", http.StatusForbidden)
			}

			userId, err := db.GetUserIDsByName(username)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}
			//create user and put it in context
			user := db.User{
				ID:       userId,
				Username: username,
			}
			ctx := context.WithValue(r.Context(), ctxKey, &user)
			//call with new context
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *db.User {
	raw, _ := ctx.Value(ctxKey).(*db.User)
	return raw
}
