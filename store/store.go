package store

import (
	"fmt"
	"gonotiserv/model"
	"sync"
	"time"
)

type DATA struct {
	mu            sync.RWMutex
	count         int
	notifications map[string]*model.Notification
}

func NewStore() *DATA {
	return &DATA{
		count:         0,
		notifications: make(map[string]*model.Notification),
	}
}

func (s *DATA) AddNotification(n *model.Notification) string {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.count++

	id := fmt.Sprintf("Notif-%d", s.count)

	n.ID = id
	n.CreatedAt = time.Now()
	n.UpdatedAt = time.Now()

	s.notifications[id] = n

	return id
}

func (s *DATA) GetNotification(id string) *model.Notification {
	s.mu.RLock()
	defer s.mu.RUnlock()

	item, ok := s.notifications[id]
	if !ok {
		return nil
	}

	return item
}

func (s *DATA) UpdateStatus(
	id string,
	status string,
	attempts int,
) error {

	s.mu.Lock()
	defer s.mu.Unlock()

	n, ok := s.notifications[id]
	if !ok {
		return fmt.Errorf(
			"notification not found: %s",
			id,
		)
	}

	n.Status = status
	n.Attempts = attempts
	n.UpdatedAt = time.Now()

	return nil
}

func (s *DATA) GetAllNotifications(email string) []*model.Notification {
	s.mu.RLock()
	defer s.mu.RUnlock()

	notifications := make(
		[]*model.Notification,
		0,
		len(s.notifications),
	)

	for _, n := range s.notifications {

		if n.Recipient == email {
			notifications = append(
				notifications,
				n,
			)
		}
	}

	return notifications
}

// DeleteNotification removes a notification by ID.
func (s *DATA) DeleteNotification(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.notifications[id]; !ok {
		return fmt.Errorf(
			"notification not found: %s",
			id,
		)
	}

	delete(s.notifications, id)

	return nil
}
