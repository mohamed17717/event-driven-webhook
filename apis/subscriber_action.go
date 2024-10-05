package apis

import (
	"event-driven-webhook/config"
	"event-driven-webhook/models"
	"event-driven-webhook/viewsets"
)

type CreateSubscriptionActionInput struct {
	SubscriberID uint `json:"subscriber_id" binding:"required"`
	ActionID     uint `json:"action_id" binding:"required"`
}

func InputToSubscriptionAction(data *CreateSubscriptionActionInput) models.SubscriberAction {
	return models.SubscriberAction{
		SubscriberID: data.SubscriberID,
		ActionID:     data.ActionID,
	}
}

func SubscriberActionRoutes() {
	protected := ProtectedRoute()

	subscriberActionViewSet := viewsets.ViewSet[models.SubscriberAction, CreateSubscriptionActionInput, Empty]{
		DB:                   config.DB,
		InputOfCreateToModel: InputToSubscriptionAction,
	}

	protected.POST("/subscriber-action", subscriberActionViewSet.Create)
	protected.GET("/subscriber-action", subscriberActionViewSet.List)
	protected.DELETE("/subscriber-action/:id", subscriberActionViewSet.Delete)
}
