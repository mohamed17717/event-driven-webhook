package apis

import (
	"event-driven-webhook/config"
	"event-driven-webhook/models"
	"event-driven-webhook/viewsets"
	"github.com/gin-gonic/gin"
	"time"
)

type CreateChangeInput struct {
	ActionID   uint      `json:"action_id" binding:"required"`
	Data       string    `form:"data" json:"data" binding:"required"`
	Identifier string    `form:"identifier" json:"identifier" binding:"required"`
	When       time.Time `form:"when" json:"when" binding:"required"`
}

func CustomChangeCreate(c *gin.Context, obj *models.Change) error {
	user := c.MustGet("user").(models.User)
	obj.UserID = user.ID
	return nil
}

func InputToChange(data *CreateChangeInput) models.Change {
	return models.Change{
		ActionID:   data.ActionID,
		Data:       data.Data,
		Identifier: data.Identifier,
		When:       data.When,
	}
}

func ChangeRoutes() {
	protected := ProtectedRoute()

	changeViewSet := viewsets.ViewSet[models.Change, CreateChangeInput, Empty]{
		DB:                   config.DB,
		PerformCreateFunc:    CustomChangeCreate,
		InputOfCreateToModel: InputToChange,
	}

	protected.POST("/actions", changeViewSet.Create)
	protected.GET("/actions", changeViewSet.List)
}
