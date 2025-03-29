package routes

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joaooliveira247/backend-test/src/controllers"
	"github.com/joaooliveira247/backend-test/src/db"
	"github.com/joaooliveira247/backend-test/src/repositories"
)

func CompetitionsRoute(eng *gin.Engine) {
	gormDB, err := db.GetDBConnection()

	if err != nil {
		log.Fatalf("Users Route: %v", err)
	}

	usersReposiotry := repositories.NewUsersRepository(gormDB)
	compsRepository := repositories.NewCompetiotionsRepository(gormDB)
	pointsRepository := repositories.NewPointsRepository(gormDB)

	controller := controllers.NewCompetitionsController(
		usersReposiotry,
		compsRepository,
		pointsRepository,
	)

	competitionsGroup := eng.Group("/competitions")
	{
		competitionsGroup.POST("/", controller.Create)
		competitionsGroup.GET("/", controller.GetCompetition)
		competitionsGroup.PUT("/", controller.CloseCompetition)
		competitionsGroup.GET("/reports/", controller.GetCompetitionReport)
	}
}