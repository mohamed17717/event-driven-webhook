package apis

import (
	"bytes"
	"encoding/json"
	"event-driven-webhook/config"
	"event-driven-webhook/models"
	"event-driven-webhook/viewsets"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"time"
)

type CreateChangeInput struct {
	ActionID   uint      `json:"action_id" binding:"required"`
	Data       string    `form:"data" json:"data" binding:"required"`
	Identifier string    `form:"identifier" json:"identifier" binding:"required"`
	When       time.Time `form:"when" json:"when" binding:"required" time_format:"2006-01-02T15:04:05Z07:00"`
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

func readBody(c *gin.Context) ([]byte, error) {
	// Read the body and store it
	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return nil, err
	}

	// Restore the io.ReadCloser to the original state so it can be read again
	c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	return bodyBytes, nil
}

func ChangeRoutes() {
	protected := ProtectedRoute()

	var changeViewSet = viewsets.ViewSet[models.Change, CreateChangeInput, Empty]{
		DB:                   config.DB,
		PerformCreateFunc:    CustomChangeCreate,
		InputOfCreateToModel: InputToChange,
	}

	creatChange := func(c *gin.Context) {
		//take a copy of the request
		bodyBytes, err := readBody(c)
		//just normal create
		changeViewSet.Create(c)

		// load the data
		var data CreateChangeInput
		json.Unmarshal(bodyBytes, &data)

		bytesData, err := json.Marshal(data)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Publish change to rabbit mq
		config.MQPublishChange(string(bytesData))

		changeTime, _ := data.When.MarshalText()
		config.RedisSet(data.Identifier, string(changeTime))
	}

	protected.POST("/changes", creatChange)
	protected.GET("/changes", changeViewSet.List)
}
