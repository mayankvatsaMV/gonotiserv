package model

type NotificationStatusEvent struct {
	NotificationID string `json:"notification_id"`
	Status         string `json:"status"`
	Attempts       int    `json:"attempts"`
}
