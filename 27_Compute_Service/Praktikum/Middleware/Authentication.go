package middleware

// authentication jwt
import (
	"github.com/NaufalFarros/golang-fiber/config"
	"github.com/NaufalFarros/golang-fiber/helper/responseJSON"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Authentication(c *fiber.Ctx) error {

	// ambil token dari header
	token := c.Cookies("token")

	// jika token tidak ada
	if token == "" {
		return responseJSON.ResponseJSON(c, 401, "Unauthorized", nil)
	}

	// cek token
	_, err := jwt.ParseWithClaims(token, &config.JTWClaim{}, func(token *jwt.Token) (interface{}, error) {
		return config.JWT_KEY, nil
	})

	// jika token tidak valid
	if err != nil {
		return responseJSON.ResponseJSON(c, 401, "Unauthorized", nil)
	}

	// jika token valid
	return c.Next()

}
