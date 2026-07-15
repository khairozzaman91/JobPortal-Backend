package middlewares

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/khairozzaman91/JobPortal-Backend/config"
	"github.com/khairozzaman91/JobPortal-Backend/utils"
)

type Claims struct {
	Sub       int    `json:"sub"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Iat       int64  `json:"iat"`
	Exp       int64  `json:"exp"`
}

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		header := r.Header.Get("Authorization")

		if header == "" {
			http.Error(w, "Unauthorized: No token provided", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(header, " ")

		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Unauthorized: Invalid authorization header", http.StatusUnauthorized)
			return
		}

		token := parts[1]

		tokenParts := strings.Split(token, ".")

		if len(tokenParts) != 3 {
			http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
			return
		}

		headerPart := tokenParts[0]
		payloadPart := tokenParts[1]
		signaturePart := tokenParts[2]

		cfg := config.GetConfig()

		signingInput := headerPart + "." + payloadPart

		h := hmac.New(sha256.New, []byte(cfg.JWTSecret))
		_, err := h.Write([]byte(signingInput))
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		expectedSignature := utils.Base64UrlEncode(h.Sum(nil))

		if !hmac.Equal([]byte(expectedSignature), []byte(signaturePart)) {
			http.Error(w, "Unauthorized: Invalid signature", http.StatusUnauthorized)
			return
		}

		payloadBytes, err := utils.Base64UrlDecode(payloadPart)
		if err != nil {
			http.Error(w, "Unauthorized: Invalid payload", http.StatusUnauthorized)
			return
		}

		var claims Claims

		err = json.Unmarshal(payloadBytes, &claims)
		if err != nil {
			http.Error(w, "Unauthorized: Invalid payload", http.StatusUnauthorized)
			return
		}

		if time.Now().Unix() > claims.Exp {
			http.Error(w, "Unauthorized: Token expired", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
