package apis

import (
	"event-driven-webhook/config"
	"event-driven-webhook/models"
	"event-driven-webhook/viewsets"
)

type UserConfInput struct {
	LatestChangeOnly *bool `json:"latest_change_only"`
	RetryFailure     *bool `json:"retry_failure"`
	MaxRetries       *int  `json:"max_retries"`
}

func InputToUserConf(data *UserConfInput) models.UserConfiguration {
	userConf := models.UserConfiguration{}
	if data.LatestChangeOnly != nil {
		userConf.LatestChangeOnly = *data.LatestChangeOnly
	}

	if data.RetryFailure != nil {
		userConf.RetryFailure = *data.RetryFailure
	}

	if data.MaxRetries != nil {
		userConf.MaxRetries = *data.MaxRetries
	}

	return userConf
}

func UserConfigurationRoutes() {
	protected := ProtectedRoute()

	userConfigurationViewSet := viewsets.ViewSet[models.UserConfiguration, UserConfInput, UserConfInput]{
		DB:                   config.DB,
		InputOfUpdateToModel: InputToUserConf,
	}

	protected.GET("/actions/:id", userConfigurationViewSet.Retrieve)
	protected.PUT("/actions/:id", userConfigurationViewSet.Update)
}
