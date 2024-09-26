package middleware

import (
	"net/http"

	"github.com/Rakhulsr/go-rest-api-pzn/helper"
	"github.com/Rakhulsr/go-rest-api-pzn/model/web"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewMiddleware(h http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: h}
}

func (middleware *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if "SECRET" == r.Header.Get("X-API-KEY") {
		middleware.Handler.ServeHTTP(w, r)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}
		helper.WriteJson(w, webResponse)
	}
}
