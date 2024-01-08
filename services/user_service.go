package services

import (
	"context"
	"errors"
	"fmt"

	"filip.filipovic/polling-app/db"
	"filip.filipovic/polling-app/dto"
	"filip.filipovic/polling-app/model/ent"
	"filip.filipovic/polling-app/utils"
)

func RegisterUser(ctx context.Context, userReq *dto.UserRequest) (*dto.UserResponse, error) {
	user := &ent.User{
		FirstName: userReq.FirstName,
		LastName:  userReq.LastName,
		Username:  userReq.Username,
		Password:  userReq.Password,
	}

	existingUser, _ := db.FindUserByUsername(ctx, user.Username)
	if existingUser != nil {
		return nil, errors.New(fmt.Sprintf("User with the username '%s' already exists", user.Username))
	}

	user, err := db.SaveUser(ctx, user)
	return makeUserResponse(user), err
}

func LoginUser(ctx context.Context, loginReq *dto.LoginRequest) (*dto.UserResponse, error) {
	user, err := db.FindUserByUsername(ctx, loginReq.Username)
	if err != nil {
		return nil, err
	}

	if !utils.CheckPasswordHash(loginReq.Password, user.Password) {
		return nil, errors.New("Incorrect password")
	}
	return makeUserResponse(user), nil
}



func makeUserResponse(user *ent.User) *dto.UserResponse {
	return &dto.UserResponse{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		ID:        user.ID,
		Username:  user.Username,
	}
}

func makeUserResponses(users []*ent.User) []*dto.UserResponse {
	var userResponses []*dto.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, makeUserResponse(user))
	}
	return userResponses
}
