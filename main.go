package main

import (
	"event-driven-webhook/config"
	"event-driven-webhook/models"
	"event-driven-webhook/utils"
)

func main() {
	utils.LoadEnv()
	config.ConnectDatabase()
	err := config.DB.AutoMigrate(
		&models.User{}, &models.UserConfiguration{},
		&models.Action{}, &models.Subscriber{}, &models.SubscriberAction{},
		&models.Change{}, &models.WebhookLog{},
	)

	utils.CheckErr(err, true)
}
