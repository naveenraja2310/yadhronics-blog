package middleware

import (
	"net/http"
	"time"
	"yadhronics-blog/response"
	"yadhronics-blog/security"
	"yadhronics-blog/utils"

	"github.com/gofiber/fiber/v2"
)

// UserClaims represents the user claims extracted from a JWT token
type UserClaims struct {
	Email string
}

// VerifyToken verifies a JWT token and returns the user claims
func VerifyToken(token string) (*UserClaims, error) {
	claims, err := security.Extractclaims(token)
	if err != nil {
		return nil, err
	}

	return &UserClaims{
		Email: claims.Email,
	}, nil
}

func JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token, err := utils.ExtractTokenFromHeader(c)
		if err != nil {
			return c.Status(http.StatusUnauthorized).JSON(response.ErrorResponse{
				ApiPath:      c.OriginalURL(),
				ErrorCode:    http.StatusUnauthorized,
				ErrorMessage: err.Error(),
				ErrorTime:    time.Now(),
			})
		}

		claims, err := security.Extractclaims(token)
		if err != nil {
			return c.Status(http.StatusUnauthorized).JSON(response.ErrorResponse{
				ApiPath:      c.OriginalURL(),
				ErrorCode:    http.StatusUnauthorized,
				ErrorMessage: err.Error(),
				ErrorTime:    time.Now(),
			})
		}

		c.Locals("admin_email", claims.Email)

		return c.Next()
	}
}
