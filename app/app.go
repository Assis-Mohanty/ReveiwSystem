package app

import (
	"fmt"
	"log"
	"os"
	"reviewservice/config"
	"reviewservice/controllers"
	repo "reviewservice/db/repository"
	"reviewservice/routes"
	"reviewservice/services"

	"github.com/gofiber/fiber/v2"
)

type Config struct {
	Address string
}

type Application struct {
	Config  Config
	Storage repo.Storage
}

func NewConfig(address string)Config{
	return Config{
		Address: address,
	}
}

func NewApplication (config Config,storage repo.Storage)Application{
	return Application{
		Config:  config,
		Storage: storage,

	}
}

func (app *Application) Run() error{
	db,err:=config.SetupDb()
	if err != nil {
		log.Panic(err)
	}
	application:=fiber.New()
	rr:=repo.NewReviewRepository(db)
	rs:=services.NewReviewService(rr)
	rc:=controllers.NewReviewController(rs)
	rrt:=routes.NewReviewRouter(rc)

	routes.SetupRouter(application,rrt)
	fmt.Println("Server starting at :",app.Config.Address)
	application.Listen(os.Getenv("PORT"))
	return nil
}