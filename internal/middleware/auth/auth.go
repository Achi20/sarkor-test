package auth

import (
	"database/sql"
	"net/http"

	"sarkor-test/internal/pkg/config"
	"sarkor-test/internal/repository/user"
	
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type Auth struct {
	user User
}

func New(user User) *Auth {
	return &Auth{user}
}

type Claims struct {
	Login  string `json:"login"`
	UserID int    `json:"user_id"`
	jwt.StandardClaims
}

func (am Auth) AuthCookie(c *gin.Context) {
	token, err := c.Cookie("SESSTOKEN")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"message": "forbidden with no cookie",
			"status":  false,
		})

		return
	}

	claims := new(Claims)
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetConf().JWTKey), nil
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"status":  false,
		})

		return
	}

	if !parsedToken.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "invalid token",
			"status":  false,
		})

		return
	}

	err = am.user.GetAuthDetail(c, user.CookieAuth{
		UserID: claims.UserID,
		Login:  claims.Login,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "invalid cookie",
				"status":  false,
			})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
				"status":  false,
			})
		}

		return
	}

	c.Set("user_id", claims.UserID)

	c.Next()
}
