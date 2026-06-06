package middleware

import (
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		header :=
			c.GetHeader(
				"Authorization",
			)

		if header == "" {

			c.AbortWithStatus(401)
			return
		}

		tokenString :=
			strings.TrimPrefix(
				header,
				"Bearer ",
			)

		token, err := jwt.Parse(
			tokenString,
			func(
				token *jwt.Token,
			) (
				interface{},
				error,
			) {

				return []byte(
					os.Getenv(
						"JWT_SECRET",
					),
				), nil
			},
		)

		if err != nil ||
			!token.Valid {

			c.AbortWithStatus(401)
			return
		}

		claims :=
			token.Claims.
				(jwt.MapClaims)

		c.Set(
			"user_id",
			uint(
				claims["user_id"].
					(float64),
			),
		)

		c.Next()
	}
}