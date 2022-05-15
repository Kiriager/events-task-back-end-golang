package handler

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func JwtAuthentication() gin.HandlerFunc {

	return func(c *gin.Context) {

		tokenHeader := c.GetHeader("Authorization") //Grab the token from the header

		if tokenHeader == "" { //Token is missing, returns with error code 403 Unauthorized
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"auth_message": "Missing auth token"})
			return
		}

		splitted := strings.Split(tokenHeader, " ") //The token normally comes in format `Bearer {token-body}`, we check if the retrieved token matched this requirement
		if len(splitted) != 2 {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"auth_message": "Invalid/Malformed auth token"})
			return
		}

		tokenPart := splitted[1] //Grab the token part, what we are truly interested in
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})

		if err != nil { //Malformed token, returns with http code 403 as usual
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"auth_message": "Malformed authentication token"})
			return
		}

		if !token.Valid { //Token is invalid, maybe not signed on this server
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"auth_message": "Token is not valid"})
			return
		}

		_, err = models.FetchAuth(tk)
		if err != nil { //Token was invalidated
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"auth_message": "Token was invalidated"})
			return
		}

		//Everything went well, proceed with the request and set the caller to the user retrieved from the parsed token
		c.Set("user", tk.UserId)
		c.Set("auth", tk.AuthUUID)
		c.Next()
	}
}
