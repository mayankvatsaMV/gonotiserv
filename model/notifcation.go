package model

import "time"

type Notification struct {
	ID          string     `json:"id"`
	Channel     string     `json:"channel"`
	Recipient   string     `json:"recipient"`
	Subject     string     `json:"subject"`
	Body        string     `json:"body"`
	Status      string     `json:"status"`
	Attempts    int        `json:"attempts"`
	ScheduledAt *time.Time `json:"scheduled_at,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

const (
	Pending    = "PENDING"
	Processing = "PROCESSING"
	Sent       = "SENT"
	Failed     = "FAILED"
	Scheduled  = "SCHEDULED"
)
