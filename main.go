package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"gonotiserv/config"
	"gonotiserv/controller"
	"gonotiserv/mail"
	"gonotiserv/queue"
	"gonotiserv/repository"
	"gonotiserv/routes"
	"gonotiserv/service"
	"gonotiserv/store"
	"gonotiserv/websocket"
	"gonotiserv/worker"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal(".env file not found")
	}

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	err = mail.InitSMTP(cfg)
	if err != nil {
		log.Fatal(err)
	}

	defer mail.CloseSMTP()

	log.Printf(
		"SMTP Config Loaded: %s:%s",
		cfg.SMTPHost,
		cfg.SMTPPort,
	)

	st := store.NewStore()

	repo := repository.NewNotificationRepository(st)
	hub := websocket.NewHub()

	jobs := queue.NewJobQueue()
	workerCount := 5

	for i := 1; i <= workerCount; i++ {
		go worker.StartWorker(
			i,
			jobs,
			repo,
			hub,
		)
	}

	notificationService :=
		service.NewNotificationService(repo,
			jobs)

	notificationController :=
		controller.NewNotificationController(
			notificationService,
		)

	router := gin.Default()

	routes.RegisterRoutes(
		router,
		notificationController,
		hub,
	)

	log.Println("Server started on :8080")

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
