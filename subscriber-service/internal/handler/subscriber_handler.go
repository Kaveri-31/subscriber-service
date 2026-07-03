package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"subscriber-service/internal/model"
	"subscriber-service/internal/service"
	"subscriber-service/internal/metrics"
)

// GetSubscribers godoc
//
// @Summary Get all subscribers
// @Description Returns all subscribers
// @Tags Subscribers
// @Produce json
// @Success 200 {array} model.Subscriber
// @Failure 500 {object} map[string]string
// @Router /subscribers [get]
// Get All Subscribers
func GetSubscribers(c *gin.Context) {
	
	metrics.SubscriberRequestsTotal.Inc()

	subscribers, err := service.GetAllSubscribers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, subscribers)
}

// GetSubscriberByID godoc
//
// @Summary Get subscriber by ID
// @Description Returns a subscriber
// @Tags Subscribers
// @Produce json
// @Param id path string true "Subscriber ID"
// @Success 200 {object} model.Subscriber
// @Failure 404 {object} map[string]string
// @Router /subscribers/{id} [get]
// Get Subscriber By ID
func GetSubscriberByID(c *gin.Context) {
	
	metrics.SubscriberRequestsTotal.Inc()

	id := c.Param("id")

	subscriber, err := service.GetSubscriberByID(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, subscriber)
}

// CreateSubscriber godoc
//
// @Summary Create subscriber
// @Description Creates a new subscriber
// @Tags Subscribers
// @Accept json
// @Produce json
// @Param subscriber body model.Subscriber true "Subscriber"
// @Success 201 {object} model.Subscriber
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /subscribers [post]
// Create Subscriber
func CreateSubscriber(c *gin.Context) {
	
	metrics.SubscriberRequestsTotal.Inc()
	metrics.SubscriberCreateTotal.Inc()

	var subscriber model.Subscriber

	if err := c.ShouldBindJSON(&subscriber); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	createdSubscriber, err := service.CreateSubscriber(subscriber)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, createdSubscriber)
}

// UpdateSubscriber godoc
//
// @Summary Update subscriber
// @Description Updates subscriber details
// @Tags Subscribers
// @Accept json
// @Produce json
// @Param id path string true "Subscriber ID"
// @Param subscriber body model.Subscriber true "Subscriber"
// @Success 200 {object} model.Subscriber
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /subscribers/{id} [put]
// Update Subscriber
func UpdateSubscriber(c *gin.Context) {
	
	metrics.SubscriberRequestsTotal.Inc()
	metrics.SubscriberUpdateTotal.Inc()

	id := c.Param("id")

	var subscriber model.Subscriber

	if err := c.ShouldBindJSON(&subscriber); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	updatedSubscriber, err := service.UpdateSubscriber(id, subscriber)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, updatedSubscriber)
}

// DeleteSubscriber godoc
//
// @Summary Delete subscriber
// @Description Deletes a subscriber
// @Tags Subscribers
// @Produce json
// @Param id path string true "Subscriber ID"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /subscribers/{id} [delete]
// Delete Subscriber
func DeleteSubscriber(c *gin.Context) {
	
	metrics.SubscriberRequestsTotal.Inc()
	metrics.SubscriberDeleteTotal.Inc()

	id := c.Param("id")

	err := service.DeleteSubscriber(id)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Subscriber deleted successfully",
	})
}
