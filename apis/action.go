package apis

import (
	"event-driven-webhook/config"
	"event-driven-webhook/models"
	"event-driven-webhook/viewsets"
	"github.com/gin-gonic/gin"
)

func CustomActionCreate(c *gin.Context, obj *models.Action) error {
	user := c.MustGet("user").(models.User)
	obj.UserID = user.ID
	return nil
}

type CreateActionInput struct {
	EventName string `json:"event_name" binding:"required"`
}

// Input Alisa name
type Input = CreateActionInput

func InputToAction(data *Input) models.Action {
	return models.Action{
		EventName: data.EventName,
	}
}

func ActionRoutes() {
	protected := ProtectedRoute()

	actionViewSet := viewsets.ViewSet[models.Action, Input, Input]{
		DB:                   config.DB,
		PerformCreateFunc:    CustomActionCreate,
		InputOfCreateToModel: InputToAction,
		InputOfUpdateToModel: InputToAction,
	}

	protected.POST("/actions", actionViewSet.Create)
	protected.GET("/actions", actionViewSet.List)
	protected.GET("/actions/:id", actionViewSet.Retrieve)
	protected.PUT("/actions/:id", actionViewSet.Update)
	protected.DELETE("/actions/:id", actionViewSet.Delete)
}
