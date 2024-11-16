package middleware

import "net/http"

// Chain structure, the next handler is determined by what Next is set to
// AuthMiddleware ..
type AuthMiddleware struct {
	Next http.Handler
}

func (am *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// If there is only one middleware and its Next field is nil, 
	// pass the request to the default router for handling
	if am.Next == nil {
		am.Next = http.DefaultServeMux
	}

	// Check the Authorization header
	auth := r.Header.Get("Authorization")
	if auth != "" {
		// Before passing to the router
		am.Next.ServeHTTP(w, r)
		// After passing to the router
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
