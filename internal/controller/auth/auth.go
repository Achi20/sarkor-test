package auth

import (
	"context"
	"database/sql"
	"net/http"

	"sarkor-test/internal/pkg/util/request"
	"sarkor-test/internal/repository/user"
	"sarkor-test/internal/usecase/auth"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	auth *auth.UseCase
}

func New(auth *auth.UseCase) *Controller {
	return &Controller{auth}
}

func (ac Controller) Auth(c *gin.Context) {
	var data user.Auth

	if err := c.ShouldBind(&data); err != nil {
		out, err := request.FieldErrorCheck(err)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"errors": err.Error(),
				"status": false,
			})
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errors": out,
			"status": false,
		})

		return
	}

	ctx := context.Background()

	token, err := ac.auth.Auth(ctx, data)
	if err != nil {
		if err == sql.ErrNoRows {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "user doesn't exist",
				"status":  false,
			})
		} else if err.Error() == "incorrect password" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
				"status":  false,
			})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
				"status":  false,
			})
		}
	}

	c.SetCookie("SESSTOKEN", token, 3600, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"status":  true,
	})
}
