package middleware

import (
	"filip.filipovic/polling-app/handlers"
	"github.com/julienschmidt/httprouter"
)

const (
	USER_PATH         = "/user"
	POLL_PATH         = "/poll"
	POLL_OPTIONS_PATH = "/poll-options"
)

// Route() Makes a new router and routes all handlers to their paths
func RoutePaths() *httprouter.Router {
	router := httprouter.New()

	routeUsers(router)
	routePolls(router)
	routePollOptions(router)

	return router
}

// routeUsers() Routes paths related to users
func routeUsers(router *httprouter.Router) {
	router.POST(USER_PATH+"/register", handlers.RegisterUser)
	router.POST(USER_PATH+"/login", handlers.LoginUser)
}

// routePolls() Routes paths related to polls
func routePolls(router *httprouter.Router) {
	router.GET(POLL_PATH+"/id/:id", handlers.GetPoll)
	router.GET(POLL_PATH, handlers.GetPolls)
	router.GET(POLL_PATH+"/created-by/:createdByUsername", handlers.GetPollsCreatedByUser)

	router.POST(POLL_PATH+"/create", handlers.CreatePoll)

	router.POST(POLL_PATH+"/modify", handlers.ModifyPoll)
}

// routePollOptions() Routes paths related to poll options
func routePollOptions(router *httprouter.Router) {
	router.GET(POLL_OPTIONS_PATH+"/:id/users", handlers.GetUsersForPollOption)

	router.POST(POLL_OPTIONS_PATH+"/vote", handlers.VoteForPollOption)
}
