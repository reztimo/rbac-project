package controllers

import (
	"maoelana/RBAC-project/initializers"
	"maoelana/RBAC-project/models"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateResource(c *gin.Context) {
	var eventData struct {
		Title       string
		Description string
		StartTime   time.Time
		EndTime     time.Time
	}

	c.Bind(&eventData)

	event := models.Event{
		Title:       eventData.Title,
		Description: eventData.Description,
		StartTime:   eventData.StartTime,
		EndTime:     eventData.EndTime,
	}

	result := initializers.DB.Create(&event)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"event": event,
	})
}

func GetAllResource(c *gin.Context) {
	var events []models.Event
	initializers.DB.Find(&events)

	c.JSON(200, gin.H{
		"event": events,
	})
}

func GetResource(c *gin.Context) {
	id := c.Param("id")

	var event models.Event

	initializers.DB.First(&event, id)

	c.JSON(200, gin.H{
		"event": event,
	})
}

func UpdateResource(c *gin.Context) {
	id := c.Param("id")

	var eventData struct {
		Title       string
		Description string
		StartTime   time.Time
		EndTime     time.Time
	}

	c.Bind(&eventData)

	var event models.Event
	initializers.DB.First(&event, id)

	initializers.DB.Model(&event).Updates(models.Event{
		Title:       eventData.Title,
		Description: eventData.Description,
		StartTime:   eventData.StartTime,
		EndTime:     eventData.EndTime,
	})

	c.JSON(200, gin.H{
		"evet": event,
	})
}

func DeleteResource(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Event{}, id)

	c.Status(200)
}
