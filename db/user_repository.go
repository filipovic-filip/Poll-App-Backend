package db

import (
	"context"

	"filip.filipovic/polling-app/config"
	"filip.filipovic/polling-app/model/ent"
	"filip.filipovic/polling-app/model/ent/user"
	"filip.filipovic/polling-app/utils"
)

func SaveUser(ctx context.Context, user *ent.User) (*ent.User, error) {
	client := config.AppConfig.Client

	hash, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	return client.User.Create().
		SetFirstName(user.FirstName).
		SetLastName(user.LastName).
		SetUsername(user.Username).
		SetPassword(hash).
		Save(ctx)
}

func FindUserByUsername(ctx context.Context, username string) (*ent.User, error) {
	client := config.AppConfig.Client

	return client.User.Query().
		Where(user.Username(username)).
		Only(ctx)
}
