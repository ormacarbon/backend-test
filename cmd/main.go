package main

import (
	"fmt"
	"os"

	"github.com/Andreffelipe/carbon_offsets_api/config"
	"github.com/Andreffelipe/carbon_offsets_api/internal"
	"github.com/gin-gonic/gin"
)

func main() {
	configLogger := internal.Config{
		OutputPaths: []string{"stdout"},
		Level:       "debug",
		PrettyPrint: true,
		WithCaller:  true,
	}

	log, err := internal.NewLogger(configLogger)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao inicializar logger: %v\n", err)
		os.Exit(1)
	}

	router := gin.Default()
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	smtp := internal.NewSMTP(config.SMTPHost, config.SMTPPort, config.SMTPUser, config.SMTPPass)

	event := internal.NewEventBus()

	connection := internal.ConnectDB(config)
	db := internal.NewPostgres(connection, log)

	createAuthor := internal.NewCreateAuthor(db, event)
	createpost := internal.NewPostCreate(db)
	postByAuthor := internal.NewFindPostByAuthor(db)
	findallPosts := internal.NewFindPost(db)
	findPostByID := internal.NewFindPostByID(db, event)
	endCompetition := internal.NewEndCompetition(db, smtp)

	increasePoint := internal.NewIncreasePoint(db, smtp)
	internal.RegisterIncreasePointHandler(event, increasePoint)

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	v1 := router.Group("/v1")
	v1.POST("/create/author", internal.CreateAuthorHttp(createAuthor, log))
	v1.POST("/create/post/:author_id", internal.CreatePostHttp(createpost, log))
	v1.GET("/find/post/:author_id", internal.FindPostByAuthorHttp(postByAuthor))
	v1.GET("/find/post/:author_id/:post_id", internal.FindPostByIDHttp(findPostByID))
	v1.GET("/find/posts", internal.FindPostHttp(findallPosts))
	v1.GET("/finish", internal.EndCompetitionHttp(endCompetition))

	router.Run(config.ServerAddress)
}
