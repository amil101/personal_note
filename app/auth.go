package app

import (
	"context"
	"encoding/base64"
	"net/http"
	"strings"
)

var BasicAuthentication = func(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// TODO : filter List of endpoints that doesn't require auth

		auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)

		if len(auth) != 2 || auth[0] != "Basic" {
			http.Error(w, "Authorization failed", http.StatusUnauthorized)
			return
		}

		payload, _ := base64.StdEncoding.DecodeString(auth[1])
		pair := strings.SplitN(string(payload), ":", 2)

		if len(pair) != 2 || !validate(pair[0], pair[1]) {
			http.Error(w, "Authorization failed", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "user", pair[0])
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r) //proceed in the middleware chain!
	})
}

func validate(username, password string) bool {
	if username == "appId1" || username == "appId2" && password == "appPassword" {
		return true
	}
	return false
}
