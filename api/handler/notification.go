package handler

import (
	pb "api_gateway/genproto/notification"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateNotification godoc
// @Summary Create a new notification
// @Description Create a new notification for a user
// @Tags notifications
// @Accept  json
// @Produce  json
// @Param notification body notification.CreateNotificationRequest true "Create Notification"
// @Success 200 {object} notification.CreateNotificationResponse
// @Failure 400 {object} gin.H "Invalid request"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /notifications [post]
func (h *Handler) CreateNotification(c *gin.Context) {
	var req pb.CreateNotificationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	resp, err := h.NotificationService.Create(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetNotification godoc
// @Summary Get a notification by ID
// @Description Get notification details by ID
// @Tags notifications
// @Accept  json
// @Produce  json
// @Param id path string true "Notification ID"
// @Success 200 {object} notification.GetNotificationResponse
// @Failure 500 {object} gin.H "Internal server error"
// @Router /notifications/get/{id} [get]
func (h *Handler) GetNotification(c *gin.Context) {
	id := c.Param("id")

	req := &pb.GetNotificationRequest{Id: id}
	resp, err := h.NotificationService.Get(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if resp.Notification == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Notification not found"})
		return
	}

	c.JSON(http.StatusOK, resp)
}
