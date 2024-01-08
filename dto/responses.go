package dto

import "filip.filipovic/polling-app/model/ent"

type UserResponse struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
}

type PollResponse struct {
	Poll         ent.Poll `json:"poll"`
	HasUserVoted bool     `json:"has_user_voted"`
}
