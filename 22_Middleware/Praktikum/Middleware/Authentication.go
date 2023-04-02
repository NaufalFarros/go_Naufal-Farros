package Middleware

import (
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// JWT middleware struct
type JWTMiddleware struct {
	SecretKey string
}

// JWTClaims struct
type JWTClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// JWTAuth middleware
func (mw *JWTMiddleware) JWTAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, "Unauthorized")
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(mw.SecretKey), nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				return c.JSON(http.StatusUnauthorized, "Unauthorized")
			}
			return c.JSON(http.StatusBadRequest, "Bad Request")
		}

		if !token.Valid {
			return c.JSON(http.StatusUnauthorized, "Unauthorized")
		}

		claims, ok := token.Claims.(*JWTClaims)
		if !ok {
			return c.JSON(http.StatusUnauthorized, "Unauthorized")
		}

		c.Set("username", claims.Username)

		return next(c)
	}
}

// JWTConfig middleware
func JWTConfig() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &JWTClaims{},
		SigningKey: []byte("NaufalFarros"),
	}
}
