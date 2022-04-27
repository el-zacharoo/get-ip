package auth

import (
	"context"
	"net/http"
)

func Authz(p string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if allow := authorise(r.Context(), p); !allow {
				http.Error(w, "insufficient permission", http.StatusForbidden)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

func authorise(ctx context.Context, permission string) bool {
	_, claims, _ := FromContext(ctx)
	ps, ok := claims["permissions"].([]interface{})
	if !ok {
		return false
	}

	allow := false
	for _, p := range ps {
		if p.(string) == permission {
			allow = true
			break
		}
	}

	return allow
}
