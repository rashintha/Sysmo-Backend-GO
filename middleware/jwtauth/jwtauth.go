package jwtauth

import (
	"fmt"
	"net/http"

	"sysmo/lib/env"
	"sysmo/lib/fireauth"
	jwtUtil "sysmo/lib/jwt"
	"sysmo/lib/log"
	"sysmo/lib/routers"
	"sysmo/lib/util"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authEnabled := true

		switch c.Request.Method {
		case "GET":
			authEnabled = !util.StrContains(routers.NoAuthGetPaths, c.Request.URL.Path)
		case "POST":
			authEnabled = !util.StrContains(routers.NoAuthPostPaths, c.Request.URL.Path)
		case "PUT":
			authEnabled = !util.StrContains(routers.NoAuthPutPaths, c.Request.URL.Path)
		case "DELETE":
			authEnabled = !util.StrContains(routers.NoAuthDeletePaths, c.Request.URL.Path)
		}

		if authEnabled {
			authHeader := c.GetHeader("Authorization")

			if authHeader == "" {
				log.Errorln("Authorization header is not present")
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}

			var token *jwt.Token
			err := error(nil)

			if c.Request.URL.Path == fmt.Sprintf("%s/api/auth/refresh", env.CONF["API_PREFIX"]) {
				token, err = jwtUtil.ValidateRefreshToken(c.GetHeader("Authorization"))
			} else {
				token, err = jwtUtil.ValidateAccessToken(c.GetHeader("Authorization"))
			}

			if c.Request.URL.Path == fmt.Sprintf("%s/api/auth/exchange", env.CONF["API_PREFIX"]) {
				email, provider := fireauth.ValidateFirebase(c.GetHeader("Authorization"))

				if len(email) == 0 || len(provider) == 0 {
					c.AbortWithStatus(http.StatusUnauthorized)
					return
				}

				c.Set("email", email)
				c.Set("authProvider", provider)
				c.Next()
				return
			}

			if err != nil {
				log.Errorln(err.Error())
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)

			if ok && token.Valid {
				data := claims["data"].(map[string]interface{})

				c.Set("id", data["id"])
				c.Set("email", data["email"])
				c.Set("name", data["name"])
				c.Next()
			} else {
				log.Warningln("Invalid token. " + err.Error())
				c.AbortWithStatus(http.StatusUnauthorized)
			}
		}
	}
}
