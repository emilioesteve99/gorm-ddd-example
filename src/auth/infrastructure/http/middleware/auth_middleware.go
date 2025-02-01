package authMiddlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	appErrors "gorm-ddd-example/src/common/application/model"
	"gorm-ddd-example/src/config"
	"net/http"
	"strings"
)

func AuthMiddleware(cfg config.Config) gin.HandlerFunc {
	protectedRoutes := map[string]bool{
		"GET /v1/users":     true,
		"GET /v1/users/:id": true,
	}

	return func(c *gin.Context) {
		if _, ok := protectedRoutes[c.Request.Method+" "+c.FullPath()]; ok {
			authHeader := c.GetHeader("Authorization")
			if authHeader == "" {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"message": appErrors.AppError{
						Code:    appErrors.UnauthorizedCode,
						Message: appErrors.UnauthorizedMsg,
					}.Error(),
				})
				return
			}
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString == authHeader {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"message": appErrors.AppError{
						Code:    appErrors.UnauthorizedCode,
						Message: appErrors.UnauthorizedMsg,
					}.Error(),
				})
				return
			}
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				return []byte(cfg.Secret), nil
			})
			if err != nil || !token.Valid {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"message": appErrors.AppError{
						Code:    appErrors.UnauthorizedCode,
						Message: appErrors.UnauthorizedMsg,
					}.Error(),
				})
				return
			}
		}

		c.Next()
	}
}
