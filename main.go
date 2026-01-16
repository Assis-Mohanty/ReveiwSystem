package main

import (
	"os"
	"reviewservice/app"
	"reviewservice/config"
	db "reviewservice/db/repository"
)

func main() {
	config.Load()
	cnf:=app.NewConfig(os.Getenv("PORT"))
	app:=app.NewApplication(cnf,db.Storage{})
	app.Run()
}