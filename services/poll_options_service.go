package services

import (
	"context"
	"errors"

	"filip.filipovic/polling-app/db"
	"filip.filipovic/polling-app/dto"
	"filip.filipovic/polling-app/model/ent"
)

func GetUsersForPollOption(ctx context.Context, id int) ([]*dto.UserResponse, error) {
	users, err := db.GetUsersForPollOption(ctx, id)
	if err != nil {
		return nil, err
	}

	return makeUserResponses(users), nil
}

func VoteForPollOption(ctx context.Context, req *dto.VoteForPollOptionRequest) (*ent.PollOption, error) {
	voter, err := db.FindUserByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}

	votedAlready, err := db.CheckUserAlreadyVotedForPollByPollOption(ctx, voter, req.OptionId)
	if err != nil {
		return nil, err
	}
	if votedAlready {
		return nil, errors.New("User already voted on this poll")
	}

	return db.AddUserToVotedListForPollOptionTx(ctx, voter, req.OptionId)
}