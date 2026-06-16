package dto

type CreateNotificationRequest struct {
	Recipient string `json:"recipient" binding:"required,email"`
	Subject   string `json:"subject" binding:"required"`
	Body      string `json:"body" binding:"required"`
}
