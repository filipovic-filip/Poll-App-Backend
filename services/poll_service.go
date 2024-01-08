package services

import (
	"context"
	"fmt"

	"filip.filipovic/polling-app/db"
	"filip.filipovic/polling-app/dto"
	"filip.filipovic/polling-app/logging"
	"filip.filipovic/polling-app/model/ent"
)

func GetPollByID(ctx context.Context, id int, username string) (*dto.PollResponse, error) {
	poll, err := db.FindPollById(ctx, id)
	if err != nil {
		fmt.Println(id)
		return nil, err
	}

	user, err := db.FindUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	hasVoted, err := db.CheckUserAlreadyVotedOnPoll(ctx, poll, user)
	if err != nil {
		return nil, err
	}

	return &dto.PollResponse{
		Poll: *poll,
		HasUserVoted: hasVoted,
	}, nil
}

func GetPolls(ctx context.Context, namePrefix string) ([]*ent.Poll, error) {
	return db.FindPolls(ctx, namePrefix)
}

func CreatePoll(ctx context.Context, req *dto.PollRequest) (*ent.Poll, error) {
	user, err := db.FindUserByUsername(ctx, req.CreatedByUsername)
	if err != nil {
		return nil, err
	}

	var pollOptions []*ent.PollOption
	for _, por := range req.PollOptions {
		pollOption := &ent.PollOption{
			Name:      por.Name,
			VoteCount: 0,
		}
		pollOptions = append(pollOptions, pollOption)
	}

	pollOptions, err = db.SavePollOptionsBulk(ctx, pollOptions)
	if err != nil {
		return nil, err
	}

	poll := &ent.Poll{
		Name:        req.Name,
		Description: req.Description,
		Edges: ent.PollEdges{
			CreatedBy:   user,
			PollOptions: pollOptions,
		},
	}
	return db.SavePoll(ctx, poll)
}

func GetPollsByCreatedByUsername(ctx context.Context, createdByUsername string) ([]*ent.Poll, error) {
	user, err := db.FindUserByUsername(ctx, createdByUsername)
	if err != nil {
		return nil, err
	}

	return db.FindPollsByCreatedBy(ctx, user)
}

func ModifyPoll(ctx context.Context, req *dto.ModifyPollRequest) (*ent.Poll, error) {
	poll, err := db.FindPollById(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	if req.Name != "" {
		poll.Name = req.Name
	}
	if req.Description != "" {
		poll.Description = req.Description
	}

	var addPollOptions []*ent.PollOption
	if len(req.AddPollOptions) > 0 {
		for _, po := range req.AddPollOptions {
			addPollOptions = append(addPollOptions, &ent.PollOption{
				Name:      po.Name,
				VoteCount: 0,
			})
		}

		addPollOptions, err = db.SavePollOptionsBulk(ctx, addPollOptions)
		if err != nil {
			return nil, err
		}
	}
	logging.Info(fmt.Sprintf("Added Poll Options: %v", addPollOptions))

	if len(req.DeletePollOptionIds) > 0 {
		_, err = db.DeletePollOptionIdsForPoll(ctx, req.DeletePollOptionIds, poll)
		if err != nil {
			return nil, err
		}
	}
	logging.Info(fmt.Sprintf("Deleted Poll Options with IDs: %v", req.DeletePollOptionIds))

	return db.ModifyPoll(ctx, poll, addPollOptions)
}
