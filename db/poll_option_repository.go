package db

import (
	"context"

	"filip.filipovic/polling-app/config"
	"filip.filipovic/polling-app/model/ent"
	"filip.filipovic/polling-app/model/ent/poll"
	"filip.filipovic/polling-app/model/ent/polloption"
	"filip.filipovic/polling-app/model/ent/user"
)

func SavePollOptionsBulk(ctx context.Context, pollOptions []*ent.PollOption) ([]*ent.PollOption, error) {
	client := config.AppConfig.Client

	return client.PollOption.MapCreateBulk(pollOptions, func(c *ent.PollOptionCreate, i int) {
		c.SetName(pollOptions[i].Name).
			SetVoteCount(pollOptions[i].VoteCount)
	}).Save(ctx)
}

func GetUsersForPollOption(ctx context.Context, pollOptionId int) ([]*ent.User, error) {
	client := config.AppConfig.Client

	return client.PollOption.Query().
		Where(polloption.ID(pollOptionId)).
		QueryUsersVoted().
		All(ctx)
}

func CheckUserAlreadyVotedForPollByPollOption(ctx context.Context, voter *ent.User, optionId int) (bool, error) {
	client := config.AppConfig.Client

	return client.PollOption.Query().
		Where(polloption.ID(optionId)).
		QueryPoll().
		QueryPollOptions().
		QueryUsersVoted().
		Where(user.ID(voter.ID)).
		Exist(ctx)
}

func AddUserToVotedListForPollOptionTx(ctx context.Context, voter *ent.User, optionId int) (*ent.PollOption, error) {
	client := config.AppConfig.Client

	return client.PollOption.UpdateOneID(optionId).
		AddUsersVoted(voter).
		AddVoteCount(1).
		Save(ctx)
}

func DeletePollOptionIdsForPoll(ctx context.Context, ids []int, pl *ent.Poll) (int, error) {
	client := config.AppConfig.Client

	return client.PollOption.Delete().
		Where(polloption.And(
			polloption.IDIn(ids...),
			polloption.HasPollWith(poll.ID(pl.ID)),
		)).
		Exec(ctx)
}
