package repository

import (
	"fmt"
	"gonotiserv/model"
	"gonotiserv/store"
)

type NotificationRepository struct {
	store *store.DATA
}

func NewNotificationRepository(
	store *store.DATA,
) *NotificationRepository {

	return &NotificationRepository{
		store: store,
	}
}

func (r *NotificationRepository) Save(
	n *model.Notification,
) string {

	return r.store.AddNotification(n)
}

func (r *NotificationRepository) GetByID(
	id string,
) (*model.Notification, error) {

	n := r.store.GetNotification(id)

	if n == nil {
		return nil, fmt.Errorf(
			"notification not found: %s",
			id,
		)
	}

	return n, nil
}

func (r *NotificationRepository) UpdateStatus(
	id string,
	status string,
	attempts int,
) error {

	return r.store.UpdateStatus(
		id,
		status,
		attempts,
	)
}

func (r *NotificationRepository) GetAll(email string) []*model.Notification {
	return r.store.GetAllNotifications(email)
}

func (r *NotificationRepository) Delete(
	id string,
) error {

	return r.store.DeleteNotification(id)
}
