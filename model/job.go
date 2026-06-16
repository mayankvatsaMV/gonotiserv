package model

type Job struct {
	NotificationID string
	Email          string
	Subject        string
	Body           string
	RetryCount     int
}
