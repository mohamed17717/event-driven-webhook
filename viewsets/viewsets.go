package viewsets

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type ViewSet[T any, C any, U any] struct {
	DB                   *gorm.DB
	PerformCreateFunc    func(c *gin.Context, obj *T) error
	InputOfCreateToModel func(n *C) T
	InputOfUpdateToModel func(n *U) T
}

func (v ViewSet[T, C, U]) Retrieve(c *gin.Context) {
	var obj T
	id := c.Param("id")

	if err := v.DB.First(&obj, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Object not found"})
		return
	}

	c.JSON(http.StatusOK, obj)
}

func (v ViewSet[T, C, U]) List(c *gin.Context) {
	var objs []T
	if err := v.DB.Find(&objs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch objects"})
		return
	}

	c.JSON(http.StatusOK, objs)
}

func (v ViewSet[T, C, U]) Create(c *gin.Context) {
	var input C
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var obj T
	obj = v.InputOfCreateToModel(&input)

	// Call the injected custom create logic
	if v.PerformCreateFunc != nil {
		if err := v.PerformCreateFunc(c, &obj); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	// Save the object after performing custom logic
	if err := v.DB.Create(&obj).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create object"})
		return
	}

	c.JSON(http.StatusOK, obj)
}

func (v ViewSet[T, C, U]) Update(c *gin.Context) {
	var obj T
	id := c.Param("id")

	if err := v.DB.First(&obj, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Object not found"})
		return
	}

	var input U
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	obj = v.InputOfUpdateToModel(&input)

	if err := v.DB.Save(&obj).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update object"})
		return
	}

	c.JSON(http.StatusOK, obj)
}

func (v ViewSet[T, C, U]) Delete(c *gin.Context) {
	var obj T
	id := c.Param("id")

	if err := v.DB.First(&obj, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Object not found"})
		return
	}

	if err := v.DB.Delete(&obj).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete object"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Object deleted"})
}
