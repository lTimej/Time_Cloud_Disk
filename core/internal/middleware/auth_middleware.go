package middleware

import (
	"liujun/Time_Cloud_Disk/core/helper"
	"net/http"
)

type AuthMiddleware struct {
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		token := r.Header.Get("Authorization")
		if token == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("未登录"))
			return
		}
		token_claim, err := helper.VerifyToken(token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(err.Error()))
			return
		}
		r.Header.Set("UserId", string(rune(token_claim.Id)))
		r.Header.Set("UserIdentity", token_claim.Identity)
		r.Header.Set("UserName", token_claim.Name)
		next(w, r)
		// Passthrough to next handler if need
		next(w, r)
	}
}
