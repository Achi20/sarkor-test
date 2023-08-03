package auth

import (
	"context"
	"errors"
	"time"

	"sarkor-test/internal/pkg/config"
	"sarkor-test/internal/repository/user"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type UseCase struct {
	user User
}

func New(user User) *UseCase {
	return &UseCase{user}
}

func (au UseCase) Auth(ctx context.Context, data user.Auth) (string, error) {
	detail, err := au.user.GetByLogin(ctx, data)
	if err != nil {
		return "", err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(detail.Password), []byte(data.Password)); err != nil {
		return "", errors.New("incorrect password")
	}

	token := jwt.New(jwt.SigningMethodHS512)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(1 * time.Hour).Unix()
	claims["login"] = data.Login
	claims["user_id"] = detail.UserID

	tokenString, err := token.SignedString([]byte(config.GetConf().JWTKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
