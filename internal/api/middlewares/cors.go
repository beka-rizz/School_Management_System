package middlewares

import (
	"net/http"
	"slices"
)

// api is hosted at www.myapi.com
// frontend server is at www.myfrontend.com

// then my frontend should include the "Origin" header with the accepted values with that origin key.

// Allowed origins
var allowedOrigins = []string{
	"https://my-origin-url.com",
	"https://localhost:3000",
	"https://www.myfrontend.com",
}

func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")

		if isOriginAllowed(origin) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		} else {
			http.Error(w, "Not allowed by CORS", http.StatusForbidden)
			return
		}
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Expose-Headers", "Authorization")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Max-Age", "3600")

		if r.Method == http.MethodOptions {
			return
		}

		next.ServeHTTP(w, r)
	})
}

func isOriginAllowed(origin string) bool {
	// for _, allowedOrigin := range allowedOrigins {
	// 	if origin == allowedOrigin {
	// 		return true
	// 	}
	// }
	// return false
	return slices.Contains(allowedOrigins, origin)
}
