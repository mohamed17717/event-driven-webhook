package main

import (
	"event-driven-webhook/apis"
	"event-driven-webhook/config"
	"event-driven-webhook/models"
	"event-driven-webhook/utils"
)

func main() {
	utils.LoadEnv()
	config.ConnectDatabase()
	config.DB.AutoMigrate(
		&models.User{}, &models.UserConfiguration{},
		&models.Action{}, &models.Subscriber{}, &models.SubscriberAction{},
		&models.Change{}, &models.WebhookLog{},
	)

	apis.AuthRoutes()
	apis.ActionRoutes()

	// Start the Gin server
	config.Server.Run()

}
