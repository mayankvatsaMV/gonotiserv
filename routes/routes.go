package routes

import (
	"gonotiserv/controller"
	"gonotiserv/websocket"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	router *gin.Engine,
	notificationController *controller.NotificationController,
	hub *websocket.Hub,
) {
	router.GET(
		"/ws",
		controller.WebSocketHandler(hub),
	)
	notifications := router.Group("/notifications")

	{
		notifications.POST(
			"",
			notificationController.CreateNotification,
		)

		notifications.GET(
			"",
			notificationController.GetNotification,
		)

		// notifications.GET(
		// 	"",
		// 	notificationController.GetAllNotifications,
		// )

		// notifications.DELETE(
		// 	"/:id",
		// 	notificationController.DeleteNotification,
		// )
	}
}
