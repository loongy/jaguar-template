package actions

import (
	"errors"

	"github.com/loongy/jaguar/models"
)

func CreateUser(ctx Context, user *models.User) (*models.User, error) {
	userID, err := ctx.Store.InsertUser(user)
	if err != nil {
		return nil, err
	}
	return GetUser(ctx, userID)
}

func GetUsers(ctx Context, offset, limit int64) (models.Users, error) {
	return ctx.Store.SelectUsers(offset, limit)
}

func GetUser(ctx Context, userID int64) (*models.User, error) {
	return ctx.Store.GetUser(userID)
}

func UpdateUser(ctx Context, user *models.User) (*models.User, error) {
	if !user.ID.Valid {
		return nil, errors.New("Unexpected invalid value 'id'")
	}
	if err := ctx.Store.UpdateUser(user); err != nil {
		return nil, err
	}
	return GetUser(ctx, user.ID.Int64)
}

func DeleteUser(ctx Context, userID int64) error {
	return ctx.Store.DeleteUser(userID)
}
