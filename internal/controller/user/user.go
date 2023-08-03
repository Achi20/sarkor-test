package user

import (
	"context"
	"net/http"
	"strconv"

	"sarkor-test/internal/pkg/util/request"
	"sarkor-test/internal/repository/phone"
	"sarkor-test/internal/repository/user"
	user_usecase "sarkor-test/internal/usecase/user"

	"github.com/gin-gonic/gin"
)

// Phone is included in user controller and usecase since there is a logical connection between them
type Controller struct {
	user *user_usecase.UseCase
}

func New(user *user_usecase.UseCase) *Controller {
	return &Controller{user}
}

// user

func (uc Controller) GetUserDetail(c *gin.Context) {
	name := c.Param("name")

	ctx := context.Background()

	detail, err := uc.user.GetUserDetail(ctx, name)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"status":  false,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"status":  true,
		"result":  detail,
	})
}

func (uc Controller) CreateUser(c *gin.Context) {
	var data user.Create

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

	if err := uc.user.CreateUser(ctx, data); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"status":  false,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"status":  true,
	})
}

// phone

func (uc Controller) GetPhoneList(c *gin.Context) {
	var filter phone.Filter

	phoneParam, ok := c.GetQuery("phone")
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "phone value is required",
			"status":  false,
		})

		return
	}

	filter.Phone = &phoneParam

	ctx := context.Background()

	list, err := uc.user.GetPhoneList(ctx, filter)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"status":  false,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"status":  true,
		"result":  list,
	})
}

func (uc Controller) CreatePhone(c *gin.Context) {
	var data phone.Create

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

	userID, ok := c.Get("user_id")
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "invalid userID from cookie",
			"status":  false,
		})

		return
	}

	val, ok := userID.(int)
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "userID from cookie must be int",
			"status":  false,
		})

		return
	}

	data.UserID = val

	ctx := context.Background()

	if err := uc.user.CreatePhone(ctx, data); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"status":  false,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"status":  true,
	})
}

func (uc Controller) UpdatePhone(c *gin.Context) {
	var data phone.Update

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

	userID, ok := c.Get("user_id")
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "invalid userID from cookie",
			"status":  false,
		})

		return
	}

	val, ok := userID.(int)
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "userID from cookie must be int",
			"status":  false,
		})

		return
	}

	data.UserID = val

	ctx := context.Background()

	if err := uc.user.UpdatePhone(ctx, data); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"status":  false,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"status":  true,
	})
}

func (uc Controller) DeletePhone(c *gin.Context) {
	idParam := c.Param("phone_id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"message": "phone_id must be int",
			"status":  false,
		})

		return
	}

	ctx := context.Background()

	err = uc.user.DeletePhone(ctx, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"status":  false,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"status":  true,
	})
}
