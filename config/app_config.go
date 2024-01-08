package config

import (
	"context"

	"filip.filipovic/polling-app/model/ent"
)

// Global AppConfig that is used to store information about the application that should be accessible from anywhere
type ApplicationConfiguration struct {
	Client         *ent.Client
	DefaultContext context.Context
}

var AppConfig *ApplicationConfiguration

func init() {
	AppConfig = &ApplicationConfiguration{}
}
