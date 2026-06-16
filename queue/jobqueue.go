package queue

import "gonotiserv/model"

func NewJobQueue() chan model.Job {
	return make(chan model.Job, 10000)
}
