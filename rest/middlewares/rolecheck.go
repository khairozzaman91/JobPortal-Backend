package middlewares

import "net/http"

func RequireRole(roles ...string) Middleware {
<<<<<<< Updated upstream

=======
>>>>>>> Stashed changes
	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			claims, ok := r.Context().Value("user").(Claims)

			if !ok {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

<<<<<<< Updated upstream
			allowed := false

			for _, role := range roles {

				if claims.Role == role {
					allowed = true
					break
				}
			}

			if !allowed {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
=======
			for _, role := range roles {
				if claims.Role == role {
					next.ServeHTTP(w, r)
					return
				}
			}

			http.Error(w, "Forbidden", http.StatusForbidden)
>>>>>>> Stashed changes
		})
	}
}