package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Andreffelipe/carbon_offsets_api/config"
	"github.com/Andreffelipe/carbon_offsets_api/internal/application/usecase"
	"github.com/Andreffelipe/carbon_offsets_api/internal/infra/database"
	"github.com/Andreffelipe/carbon_offsets_api/internal/infra/eventbus"
	"github.com/Andreffelipe/carbon_offsets_api/internal/infra/http"
	"github.com/Andreffelipe/carbon_offsets_api/internal/infra/logger"
	"github.com/Andreffelipe/carbon_offsets_api/internal/infra/repository"
	"github.com/Andreffelipe/carbon_offsets_api/internal/infra/smtp"
	"github.com/gin-gonic/gin"
)

func main() {
	configLogger := logger.Config{
		OutputPaths: []string{"stdout"},
		Level:       "debug",
		PrettyPrint: true,
		WithCaller:  true,
	}

	log, err := logger.NewLogger(configLogger)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao inicializar logger: %v\n", err)
		os.Exit(1)
	}

	router := gin.Default()
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	smtp := smtp.NewSMTP(config.SMTPHost, config.SMTPPort, config.SMTPUser, config.SMTPPass)

	event := eventbus.NewEventBus()

	connection := database.ConnectDB(config)
	db := repository.NewPostgres(connection, log)

	createAuthor := usecase.NewCreateAuthor(db, event)
	createpost := usecase.NewPostCreate(db)
	postByAuthor := usecase.NewFindPostByAuthor(db)
	findallPosts := usecase.NewFindPost(db)
	findPostByID := usecase.NewFindPostByID(db, event)
	endCompetition := usecase.NewEndCompetition(db, smtp)

	increasePoint := usecase.NewIncreasePoint(db, smtp)
	RegisterIncreasePointHandler(event, increasePoint)

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	v1 := router.Group("/v1")
	v1.POST("/create/author", http.CreateAuthorHttp(createAuthor, log))
	v1.POST("/create/post/:author_id", http.CreatePostHttp(createpost, log))
	v1.GET("/find/post/:author_id", http.FindPostByAuthorHttp(postByAuthor))
	v1.GET("/find/post/:author_id/:post_id", http.FindPostByIDHttp(findPostByID))
	v1.GET("/find/posts", http.FindPostHttp(findallPosts))
	v1.GET("/finish", http.EndCompetitionHttp(endCompetition))

	if err := router.Run(config.ServerAddress); err != nil {
		log.Fatal("cannot start server:", err)
	}
}

func RegisterIncreasePointHandler(eventBus *eventbus.EventBus, usecase *usecase.IncreasePoint) {
	ch := make(chan eventbus.Event)

	_ = eventBus.Subscribe(eventbus.EventTypeIncreasePoint, ch)

	go func() {
		for evt := range ch {
			data, ok := evt.Data.(eventbus.IncreasePointEventData)
			if !ok {
				continue
			}

			_ = usecase.Execute(context.Background(), usecase.InputIncreasePoint(data))
		}
	}()
}
