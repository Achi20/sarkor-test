package user

import (
	"context"
	"golang.org/x/crypto/bcrypt"

	"sarkor-test/internal/repository/phone"
	"sarkor-test/internal/repository/user"
)

type UseCase struct {
	user  User
	phone Phone
}

func New(user User, phone Phone) *UseCase {
	return &UseCase{user, phone}
}

// user

func (uu UseCase) CreateUser(ctx context.Context, data user.Create) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(data.Password), 9)
	if err != nil {
		return err
	}

	data.Password = string(bytes)

	return uu.user.Create(ctx, data)
}

func (uu UseCase) GetUserDetail(ctx context.Context, name string) (user.Detail, error) {
	data, err := uu.user.GetByName(ctx, name)
	if err != nil {
		return user.Detail{}, err
	}

	detail := user.Detail{
		UserID: data.UserID,
		Name:   data.Name,
		Age:    data.Age,
	}

	return detail, nil
}

// phone

func (uu UseCase) GetPhoneList(ctx context.Context, filter phone.Filter) ([]phone.List, error) {
	return uu.phone.GetAll(ctx, filter)
}

func (uu UseCase) CreatePhone(ctx context.Context, data phone.Create) error {
	return uu.phone.Create(ctx, data)
}

func (uu UseCase) UpdatePhone(ctx context.Context, data phone.Update) error {
	return uu.phone.Update(ctx, data)
}

func (uu UseCase) DeletePhone(ctx context.Context, id int) error {
	return uu.phone.Delete(ctx, id)
}
