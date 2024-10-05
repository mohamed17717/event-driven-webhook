package apis

import (
	"event-driven-webhook/config"
	"event-driven-webhook/models"
	"event-driven-webhook/viewsets"
)

func WebhookLogRoutes() {
	protected := ProtectedRoute()

	webhookLogViewSet := viewsets.ViewSet[models.WebhookLog, Empty, Empty]{
		DB: config.DB,
	}

	protected.GET("/webhook-logs", webhookLogViewSet.List)
	protected.GET("/webhook-logs/:id", webhookLogViewSet.Retrieve)
}
