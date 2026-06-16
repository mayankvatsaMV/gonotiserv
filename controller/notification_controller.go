package controller

import (
	dto "gonotiserv/dto"
	"gonotiserv/model"
	"gonotiserv/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NotificationController struct {
	service *service.NotificationService
}

func NewNotificationController(
	service *service.NotificationService,
) *NotificationController {

	return &NotificationController{
		service: service,
	}
}

func (NC *NotificationController) CreateNotification(c *gin.Context) {
	var req dto.CreateNotificationRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request payload",
			"details": err.Error(),
		})
		return
	}
	notification := &model.Notification{
		Recipient: req.Recipient,
		Subject:   req.Subject,
		Body:      req.Body,
	}
	id, err := NC.service.CreateNotification(notification)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error":   "Invalid request payload",
				"details": err.Error(),
			})
		return
	}
	c.JSON(
		http.StatusAccepted,
		dto.CreateNotificationResponse{
			ID:      id,
			Status:  model.Pending,
			Message: "notification queued successfully",
		},
	)
}

func (NC *NotificationController) GetNotification(c *gin.Context) {
	id := c.Query("id")

	if id == "" {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "id query parameter is required",
			},
		)
		return
	}
	notification, err :=
		NC.service.GetNotification(id)

	if err != nil {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}
	c.JSON(
		http.StatusOK,
		notification,
	)
}

func (NC *NotificationController) GetAllNotification(c *gin.Context) {

}
