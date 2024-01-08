package db

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"filip.filipovic/polling-app/config"
	"filip.filipovic/polling-app/model/ent"
	"filip.filipovic/polling-app/model/ent/poll"
	"filip.filipovic/polling-app/model/ent/polloption"
	"filip.filipovic/polling-app/model/ent/user"
)

func SavePoll(ctx context.Context, newPoll *ent.Poll) (*ent.Poll, error) {
	client := config.AppConfig.Client

	dbPoll, err := client.Poll.Create().
		SetName(newPoll.Name).
		SetDescription(newPoll.Description).
		SetCreatedBy(newPoll.Edges.CreatedBy).
		AddPollOptions(newPoll.Edges.PollOptions...).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return client.Poll.Query().
		Where(poll.ID(dbPoll.ID)).
		WithPollOptions().
		Only(ctx)
}

func FindPollById(ctx context.Context, id int) (*ent.Poll, error) {
	client := config.AppConfig.Client

	return client.Poll.Query().
		Where(poll.ID(id)).
		WithPollOptions(
			
		).
		Only(ctx)
}

func FindPolls(ctx context.Context, name string) ([]*ent.Poll, error) {
	client := config.AppConfig.Client

	return client.Poll.Query().
		Where(poll.NameHasPrefix(name)).
		WithPollOptions().
		Order(
			poll.ByPollOptions(
				sql.OrderBySum(
					polloption.FieldVoteCount,
					sql.OrderDesc(),
				),
			),
		).All(ctx)
}

func FindPollsByCreatedBy(ctx context.Context, createdBy *ent.User) ([]*ent.Poll, error) {
	client := config.AppConfig.Client

	return client.Poll.Query().
		WithPollOptions().
		Where(
			poll.HasCreatedByWith(
				user.ID(createdBy.ID),
			),
		).Order(
		poll.ByPollOptions(
			sql.OrderBySum(
				polloption.FieldVoteCount,
				sql.OrderDesc(),
			),
		),
	).All(ctx)
}

func ModifyPoll(ctx context.Context, oldPoll *ent.Poll, addOptions []*ent.PollOption) (*ent.Poll, error) {
	client := config.AppConfig.Client

	dbPoll, err := oldPoll.Update().
		SetName(oldPoll.Name).
		SetDescription(oldPoll.Description).
		AddPollOptions(addOptions...).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return client.Poll.Query().
		Where(poll.ID(dbPoll.ID)).
		WithPollOptions().
		Only(ctx)
}

func CheckUserAlreadyVotedOnPoll(ctx context.Context, pl *ent.Poll, usr *ent.User) (bool, error) {
	client := config.AppConfig.Client

	return client.Poll.Query().
		Where(poll.ID(pl.ID)).
		QueryPollOptions().
		QueryUsersVoted().
		Where(user.ID(usr.ID)).
		Exist(ctx)
}
