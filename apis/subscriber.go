package apis

import (
	"event-driven-webhook/config"
	"event-driven-webhook/models"
	"event-driven-webhook/viewsets"
	"github.com/gin-gonic/gin"
)

func CustomSubscriberCreate(c *gin.Context, obj *models.Subscriber) error {
	user := c.MustGet("user").(models.User)
	obj.UserID = user.ID
	return nil
}

type CreateSubscriberInput struct {
	WebhookLink string `json:"webhook_link" binding:"required"`
	SecretToken string `json:"secret_token" binding:"required"`
	IsVerified  bool   `json:"is_verified" binding:"required"`
	IsActive    bool   `json:"is_active" binding:"required"`
}

type CSInput = CreateSubscriberInput

type UpdateSubscriberInput struct {
	WebhookLink *string `json:"webhook_link"`
	SecretToken *string `json:"secret_token"`
	IsVerified  *bool   `json:"is_verified"`
	IsActive    *bool   `json:"is_active"`
}
type USInput = UpdateSubscriberInput

func CreateInputToSubscriber(data *CSInput) models.Subscriber {
	return models.Subscriber{
		WebhookLink: data.WebhookLink,
		SecretToken: data.SecretToken,
		IsVerified:  data.IsVerified,
		IsActive:    data.IsActive,
	}
}

func UpdateInputToSubscriber(data *USInput) models.Subscriber {
	subscriber := models.Subscriber{}
	if data.WebhookLink != nil {
		subscriber.WebhookLink = *data.WebhookLink
	}
	if data.SecretToken != nil {
		subscriber.SecretToken = *data.SecretToken
	}
	if data.IsVerified != nil {
		subscriber.IsVerified = *data.IsVerified
	}
	if data.IsActive != nil {
		subscriber.IsActive = *data.IsActive
	}

	return subscriber
}

func SubscriberRoutes() {
	protected := ProtectedRoute()

	subscriberViewSet := viewsets.ViewSet[models.Subscriber, CSInput, USInput]{
		DB:                   config.DB,
		PerformCreateFunc:    CustomSubscriberCreate,
		InputOfCreateToModel: CreateInputToSubscriber,
		InputOfUpdateToModel: UpdateInputToSubscriber,
	}

	protected.POST("/subscribers", subscriberViewSet.Create)
	protected.GET("/subscribers", subscriberViewSet.List)
	protected.GET("/subscribers/:id", subscriberViewSet.Retrieve)
	protected.PUT("/subscribers/:id", subscriberViewSet.Update)
	protected.DELETE("/subscribers/:id", subscriberViewSet.Delete)

}
