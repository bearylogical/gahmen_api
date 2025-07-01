package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"gahmen-api/config"
	"gahmen-api/helpers"

	jose "github.com/go-jose/go-jose/v3"
	"github.com/go-jose/go-jose/v3/jwt"
)

// contextKey is a type for context keys to avoid collisions
type contextKey string

const (
	// ContextKeyClaims is the context key for JWT claims
	ContextKeyClaims contextKey = "claims"
)

// AuthMiddleware creates a middleware for JWT authentication
func AuthMiddleware(cfg *config.Config) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				helpers.WriteJSON(w, http.StatusUnauthorized, map[string]string{"error": "Authorization header missing"})
				return
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
				helpers.WriteJSON(w, http.StatusUnauthorized, map[string]string{"error": "Invalid Authorization header format"})
				return
			}

			tokenString := parts[1]

			var publicKey *jose.JSONWebKey
			if cfg.SupabaseJWKSURL != "" {
				// Fetch JWKS from URL
				jwks, err := getJWKS(cfg.SupabaseJWKSURL)
				if err != nil {
					helpers.WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": fmt.Sprintf("Failed to fetch JWKS: %v", err)})
					return
				}
				// Find the key that matches the token's kid (key ID)
				parsedToken, err := jwt.ParseSigned(tokenString)
				if err != nil {
					helpers.WriteJSON(w, http.StatusUnauthorized, map[string]string{"error": "Invalid token format"})
					return
				}
				kid := parsedToken.Headers[0].KeyID
				for _, key := range jwks.Keys {
					if key.KeyID == kid {
						publicKey = &key
						break
					}
				}
				if publicKey == nil {
					helpers.WriteJSON(w, http.StatusUnauthorized, map[string]string{"error": "No matching public key found for token"})
					return
				}
			} else if cfg.SupabaseJWTSecret != "" {
				// Use shared secret
				publicKey = &jose.JSONWebKey{Key: []byte(cfg.SupabaseJWTSecret), Algorithm: string(jose.HS256)}
			} else {
				helpers.WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": "JWT secret or JWKS URL not configured"})
				return
			}

			token, err := jwt.ParseSigned(tokenString)
			if err != nil {
				helpers.WriteJSON(w, http.StatusUnauthorized, map[string]string{"error": "Invalid token"})
				return
			}

			claims := jwt.Claims{}
			if err := token.Claims(publicKey, &claims); err != nil {
				helpers.WriteJSON(w, http.StatusUnauthorized, map[string]string{"error": "Failed to parse token claims or verify signature"})
				return
			}

			// Add claims to context
			ctx := context.WithValue(r.Context(), ContextKeyClaims, claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// getJWKS fetches the JSON Web Key Set from the given URL
func getJWKS(jwksURL string) (*jose.JSONWebKeySet, error) {
	resp, err := http.Get(jwksURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch JWKS: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch JWKS: received status code %d", resp.StatusCode)
	}

	var jwks jose.JSONWebKeySet
	if err := json.NewDecoder(resp.Body).Decode(&jwks); err != nil {
		return nil, fmt.Errorf("failed to decode JWKS: %w", err)
	}

	return &jwks, nil
}
