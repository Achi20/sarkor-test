package main

import (
	"log"

	auth_controller "sarkor-test/internal/controller/auth"
	user_controller "sarkor-test/internal/controller/user"
	"sarkor-test/internal/middleware/auth"
	"sarkor-test/internal/pkg/db/sqlite"
	"sarkor-test/internal/repository/phone"
	"sarkor-test/internal/repository/user"
	auth_usecase "sarkor-test/internal/usecase/auth"
	user_usecase "sarkor-test/internal/usecase/user"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

// The project's architecture based on Uncle Bob's Clean Architecture
func main() {
	database, err := sqlite.New()
	if err != nil {
		log.Fatal(err)
	}

	userRepo := user.New(database)
	phoneRepo := phone.New(database)

	authMW := auth.New(userRepo)

	userUseCase := user_usecase.New(userRepo, phoneRepo)
	authUseCase := auth_usecase.New(userRepo)

	userController := user_controller.New(userUseCase)
	authController := auth_controller.New(authUseCase)

	r := gin.Default()

	v1 := r.Group("/api/v1")

	userG := v1.Group("/user")

	userG.POST("/user/register", userController.CreateUser)
	userG.POST("/user/auth", authController.Auth)

	userG.POST("/phone", authMW.AuthCookie, userController.CreatePhone)
	userG.GET("/phone", authMW.AuthCookie, userController.GetPhoneList)
	userG.PUT("/phone", authMW.AuthCookie, userController.UpdatePhone)
	userG.GET("/:name", authMW.AuthCookie, userController.GetUserDetail)
	userG.DELETE("/phone/:phone_id", authMW.AuthCookie, userController.DeletePhone)

	log.Fatal(r.Run())
}
