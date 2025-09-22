package security

import (
	"fmt"
	"time"
	"yadhronics-blog/settings"

	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
)

var JwtSigningMethod = jwt.SigningMethodHS256

type CustomClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

// GenerateJWT generates a new JWT token with custom claims
func GenerateJWT(claims CustomClaims) (string, time.Time, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claim := CustomClaims{
		Email: claims.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(JwtSigningMethod, claim)
	tokenString, err := token.SignedString([]byte(settings.Config.JWTSecret))
	if err != nil {
		return "", time.Time{}, err
	}

	return tokenString, expirationTime, nil
}

func Extractclaims(tokenStr string) (*CustomClaims, error) {
	claims := &CustomClaims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(settings.Config.JWTSecret), nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, fmt.Errorf("malformed token")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, fmt.Errorf("token is expired")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, fmt.Errorf("token not valid yet")
			} else {
				return nil, fmt.Errorf("couldn't handle this token: %v", err)
			}
		}
		return nil, fmt.Errorf("couldn't handle this token: %v", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

func GenerateJWTCookie(email string) (fiber.Cookie, error) {
	customClaims := CustomClaims{
		Email: email,
	}

	// generate token
	generatetoken, expirationTime, jwterr := GenerateJWT(customClaims)
	if jwterr != nil {
		return fiber.Cookie{}, fmt.Errorf("failed to generate token: %v", jwterr)
	}

	cookie := fiber.Cookie{
		Name:     "admintoken",
		Value:    generatetoken,
		Expires:  time.Unix(expirationTime.Unix(), 0),
		MaxAge:   int(time.Until(expirationTime).Seconds()),
		Path:     "/",
		HTTPOnly: true,
		Secure:   true,
		SameSite: "None",
	}

	return cookie, nil
}

func ClearJWTCookie(c *fiber.Ctx) fiber.Cookie {
	cookie := fiber.Cookie{
		Name:     "admintoken",
		Value:    "",
		Expires:  time.Unix(0, 0),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "None",
	}

	return cookie
}
