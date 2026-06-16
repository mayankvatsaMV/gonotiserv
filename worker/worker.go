package worker

import (
	"gonotiserv/model"
	"gonotiserv/repository"
	"gonotiserv/websocket"
	"log"
)

func StartWorker(
	id int,
	jobs <-chan model.Job,
	repo *repository.NotificationRepository,
	hub *websocket.Hub,
) {
	log.Printf(
		"worker %d started",
		id,
	)

	for job := range jobs {

		log.Printf(
			"worker %d processing notification %s",
			id,
			job.NotificationID,
		)

		err := repo.UpdateStatus(
			job.NotificationID,
			model.Processing,
			job.RetryCount,
		)
		hub.Broadcast(
			model.NotificationStatusEvent{
				NotificationID: job.NotificationID,
				Status:         model.Processing,
				Attempts:       job.RetryCount,
			},
		)

		if err != nil {
			log.Printf(
				"worker %d failed updating status: %v",
				id,
				err,
			)
			continue
		}

		// err = mail.SendEmail(
		// 	job.Email,
		// 	job.Subject,
		// 	job.Body,
		// )

		if err != nil {

			log.Printf(
				"worker %d failed sending email: %v",
				id,
				err,
			)

			_ = repo.UpdateStatus(
				job.NotificationID,
				model.Failed,
				job.RetryCount+1,
			)
			hub.Broadcast(
				model.NotificationStatusEvent{
					NotificationID: job.NotificationID,
					Status:         model.Failed,
					Attempts:       job.RetryCount + 1,
				},
			)

			continue
		}

		_ = repo.UpdateStatus(
			job.NotificationID,
			model.Sent,
			job.RetryCount+1,
		)
		hub.Broadcast(
			model.NotificationStatusEvent{
				NotificationID: job.NotificationID,
				Status:         model.Sent,
				Attempts:       job.RetryCount + 1,
			},
		)
		log.Printf(
			"worker %d completed notification %s",
			id,
			job.NotificationID,
		)
	}
}
