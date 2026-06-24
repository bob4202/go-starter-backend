package middleware

import (
	"go-starter-backend/internal/config"
	jwtPkg "go-starter-backend/pkg/jwt"
	"go-starter-backend/pkg/response"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			response.Error(c, http.StatusUnauthorized, "missing auth header")
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			response.Error(c, http.StatusUnauthorized, "invalid authorization header")
			c.Abort()
			return
		}

		tokenString := parts[1]

		claims, err := jwtPkg.ValidateToken(tokenString, cfg.JWTSecret)
		if err != nil {
			response.Error(c, http.StatusUnauthorized, "invalid or token expires")
			c.Abort()
			return
		}

		c.Set("user_id", claims.ID)
		c.Next()
	}
}
