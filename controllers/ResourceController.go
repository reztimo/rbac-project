package controllers

import (
	"maoelana/RBAC-project/initializers"
	"maoelana/RBAC-project/models"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateResource
// @Summary Create a new event resource
// @Description Create a new event by providing title, description, start time, and end time.
// @Tags resources
// @Produce json
// @Param title formData string true "Event title"
// @Param description formData string true "Event description"
// @Param startTime formData string true "Event start time (RFC3339 format)"
// @Param endTime formData string true "Event end time (RFC3339 format)"
// @Success 200 "Event created successfully"
// @Router /event [post]
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

// GetAllResource
// @Summary Get all event resources
// @Description Retrieve a list of all events.
// @Tags resources
// @Produce json
// @Success 200 "List of events retrieved successfully"
// @Router /event [get]
func GetAllResource(c *gin.Context) {
	var events []models.Event
	initializers.DB.Find(&events)

	c.JSON(200, gin.H{
		"event": events,
	})
}

// GetResource
// @Summary Get a specific event resource by ID
// @Description Retrieve a specific event by providing its ID.
// @Tags resources
// @Produce json
// @Param id path string true "Event ID"
// @Success 200 "Event retrieved successfully"
// @Router /event/{id} [get]
func GetResource(c *gin.Context) {
	id := c.Param("id")

	var event models.Event

	initializers.DB.First(&event, id)

	c.JSON(200, gin.H{
		"event": event,
	})
}

// UpdateResource
// @Summary Update a specific event resource by ID
// @Description Update a specific event by providing its ID and new data.
// @Tags resources
// @Produce json
// @Param id path string true "Event ID"
// @Param title formData string true "New event title"
// @Param description formData string true "New event description"
// @Param startTime formData string true "New event start time (RFC3339 format)"
// @Param endTime formData string true "New event end time (RFC3339 format)"
// @Success 200 "Event updated successfully"
// @Router /event/{id} [put]
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
		"event": event,
	})
}

// DeleteResource
// @Summary Delete a specific event resource by ID
// @Description Delete a specific event by providing its ID.
// @Tags resources
// @Produce json
// @Param id path string true "Event ID"
// @Success 200 "Event deleted successfully"
// @Router /event/{id} [delete]
func DeleteResource(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Event{}, id)

	c.Status(200)
}
