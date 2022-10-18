package middlewares

import (
	"context"
	"strings"

	"github.com/Napigo/go-finance-service/pkg/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/kataras/jwt"
)

type userSubKey int

var (
	UserSubKey userSubKey
	JWTSecret  = []byte("NAPIGO-JWT-SECRET")
)

// This Middleware is to extract the user_id from
// the jwt token without verifying, as verification only be done in the api gateway level.
func UserSessionMiddleware(app *fiber.App) {
	app.Use(func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")

		if !strings.HasPrefix(authHeader, "Bearer") {
			fiber.NewError(fiber.StatusUnauthorized)
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		sub, err := getSubFromToken(tokenString)
		if err != nil {
			fiber.NewError(fiber.StatusUnauthorized)
		}

		a := context.WithValue(context.Background(), UserSubKey, sub)
		c.SetUserContext(a)

		return c.Next()
	})
}

func getSubFromToken(tokenString string) (string, error) {
	logger.Info("getSubFromToken" + tokenString)
	byteToken := []byte(tokenString)

	verifiedToken, err := jwt.VerifyWithHeaderValidator(jwt.HS256, JWTSecret, byteToken, func(alg string, headerDecoded []byte) (jwt.Alg, jwt.PublicKey, jwt.InjectFunc, error) {
		return jwt.HS256, JWTSecret, nil, nil
	})
	if err != nil {
		return "", nil
	}

	sub := verifiedToken.StandardClaims.Subject
	logger.Info(verifiedToken.StandardClaims.Subject)
	logger.Info("Why not running here")
	logger.Info(sub)
	return sub, nil
}
