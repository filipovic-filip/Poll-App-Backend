package dto

//-------User Requests-------

type UserRequest struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Username  string `json:"username" validate:"required"`
	Password  string `json:"password" validate:"required"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

//-------Poll and Poll Option Requests-------

type PollRequest struct {
	Name              string              `json:"name" validate:"required"`
	Description       string              `json:"description"`
	CreatedByUsername string              `json:"created_by_username" validate:"required"`
	PollOptions       []PollOptionRequest `json:"poll_options"`
}

type ModifyPollRequest struct {
	Id                  int                 `json:"id" validate:"required"`
	Name                string              `json:"name"`
	Description         string              `json:"description"`
	DeletePollOptionIds []int               `json:"delete_poll_option_ids"`
	AddPollOptions      []PollOptionRequest `json:"add_poll_options"`
}

type PollOptionRequest struct {
	Name string `json:"name" validate:"required"`
}

type VoteForPollOptionRequest struct {
	OptionId int    `json:"option_id" validate:"required"`
	Username string `json:"username" validate:"required"`
}
