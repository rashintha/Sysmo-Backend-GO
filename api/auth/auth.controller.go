package apiAuth

import (
	"net/http"
	jwtUtil "sysmo/lib/jwt"
	authService "sysmo/services/auth"
	usersService "sysmo/services/users"

	"github.com/gin-gonic/gin"
)

func login(c *gin.Context) {
	var user usersService.User
	c.Bind(&user)

	retUser, message, _ := usersService.VerifyUser(user)

	if retUser != nil {
		authService.PostUserSignInProcess(retUser.ID.Hex())
		user.Email = retUser.Email

		authUser := authService.AuthUser{
			ID:    retUser.ID.Hex(),
			Name:  user.Name,
			Email: retUser.Email,
		}

		c.JSON(http.StatusOK, gin.H{
			"access_token":  jwtUtil.GenerateAccessToken(retUser),
			"refresh_token": jwtUtil.GenerateRefreshToken(retUser),
			"user":          authUser,
			"message":       message,
		})
		return
	}

	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": message})
}
