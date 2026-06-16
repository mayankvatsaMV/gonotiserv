package service

import (
	"gonotiserv/model"
	"gonotiserv/repository"
	"time"
)

type NotificationService struct {
	repo     *repository.NotificationRepository
	jobQueue chan<- model.Job
}

func NewNotificationService(
	repo *repository.NotificationRepository,
	jobQueue chan<- model.Job,

) *NotificationService {

	return &NotificationService{
		repo:     repo,
		jobQueue: jobQueue,
	}
}

func (s *NotificationService) CreateNotification(n *model.Notification) (string, error) {
	n.Status = model.Pending
	n.Attempts = 0
	n.CreatedAt = time.Now()
	n.UpdatedAt = time.Now()
	id := s.repo.Save(n)
	s.jobQueue <- model.Job{
		NotificationID: id,
		Email:          n.Recipient,
		Subject:        n.Subject,
		Body:           n.Body,
		RetryCount:     0,
	}
	return id, nil
}

func (s *NotificationService) GetNotification(id string) (*model.Notification, error) {
	n, err := s.repo.GetByID(id)
	return n, err

}
func (s *NotificationService) GetAllNotifications(email string) []*model.Notification {
	return s.repo.GetAll(email)
}

func (s *NotificationService) DeleteNotification(id string) error {
	return s.repo.Delete(id)
}
