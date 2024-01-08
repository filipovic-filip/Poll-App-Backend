package db

import (
	"context"
	"errors"
	"fmt"

	"filip.filipovic/polling-app/config"
	"filip.filipovic/polling-app/logging"
	"filip.filipovic/polling-app/model/ent"
	"filip.filipovic/polling-app/model/ent/migrate"
	"filip.filipovic/polling-app/model/ent/user"
	"filip.filipovic/polling-app/utils"

	_ "github.com/lib/pq"
)

// SetupDb() sets up the database
func SetupDb() {
	connectToDb()
	dropAndCreateDbSchema()
	//createDbSchema()
	seedDb()
}

// ConnectToDb() Opens a connection to the database and saves it inside the global AppConfig
// so that it can be used from anywhere to communicate with the database
func connectToDb() {
	client, err := ent.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=postgres password=postgres sslmode=disable")
	if err != nil {
		logging.FatalMsg("Couldn't connect to the database", err)
	}

	config.AppConfig.Client = client
	config.AppConfig.DefaultContext = context.Background()
}

// createDbSchema() Created the database tables if they don't exist
func createDbSchema() {
	client := config.AppConfig.Client
	ctx := config.AppConfig.DefaultContext

	err := client.Schema.Create(ctx)
	if err != nil {
		logging.FatalMsg("Couldn't make the database schema", err)
	}
}

// dropAndCreateDbSchema() Drops the database schema and creates a fresh one
// Used for testing mostly
func dropAndCreateDbSchema() {
	client := config.AppConfig.Client
	ctx := config.AppConfig.DefaultContext

	err := client.Schema.Create(
		ctx,
		migrate.WithDropColumn(true),
		migrate.WithDropIndex(true),
	)

	if err != nil {
		logging.FatalMsg("Couldn't make the database schema", err)
	}
}

// seedDb() Seeds the database with entities if the database is empty
func seedDb() {
	client := config.AppConfig.Client
	ctx := config.AppConfig.DefaultContext
	checkUsername := "user"

	_, err := client.User.
		Query().
		Where(user.Username(checkUsername)).
		Only(ctx)

	if err == nil {
		logging.Err(errors.New(fmt.Sprintf("Couldn't seed the database, user with username '%s' already exists", checkUsername)))
		return
	}
	logging.Info("Seeding the database")

	users := []ent.User{
		{FirstName: "User", LastName: "User", Username: checkUsername, Password: utils.MustHashPassword("123")},
		{FirstName: "Zika", LastName: "Zikic", Username: "zikazikic", Password: utils.MustHashPassword("zikazikic")},
		{FirstName: "Pera", LastName: "Peric", Username: "peraperic", Password: utils.MustHashPassword("peraperic")},
	}

	polls := []ent.Poll{
		{Name: "Test Poll", Description: "A poll made for test purposes"},
		{Name: "Test Poll 2", Description: "A poll made for test purposes 2"},
	}

	pollOptions := []ent.PollOption{
		{Name: "Test Option 1"},
		{Name: "Test Option 2"},
		{Name: "Test Option 3"},
	}

	var entUsers []*ent.User
	for _, user := range users {
		entUser, err := client.User.
			Create().
			SetFirstName(user.FirstName).
			SetLastName(user.LastName).
			SetUsername(user.Username).
			SetPassword(user.Password).
			Save(ctx)

		if err == nil {
			entUsers = append(entUsers, entUser)
		}
	}

	var entPolls []*ent.Poll
	for _, poll := range polls {
		entPoll, err := client.Poll.
			Create().
			SetName(poll.Name).
			SetDescription(poll.Description).
			SetCreatedBy(entUsers[0]).
			Save(ctx)

		if err == nil {
			entPolls = append(entPolls, entPoll)
		}
	}

	for i, pollOption := range pollOptions {
		poll := entPolls[0]
		if i % 4 == 0 {
			poll = entPolls[1]
		}

		client.PollOption.
			Create().
			SetName(pollOption.Name).
			SetPoll(poll).
			AddUsersVoted(entUsers[i]).
			SetVoteCount(1).
			Save(ctx)
	}
}
