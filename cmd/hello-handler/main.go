package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mike-neck/golang-lambda-deploy-example"
	"github.com/pelletier/go-toml"
	"log"
)

func main() {
	config, e := toml.LoadFile("config.toml")
	if e != nil {
		log.Fatalln("error", "main", "fail to load config.toml.", e)
	}
	app := config.Get("app")
	if app == nil {
		log.Fatalln("error", "main", "config has no app configuration.", e)
	}
	msg := app.(*toml.Tree).Get("message")
	if msg == nil {
		log.Fatalln("error", "main", "app config has no message entry.", e)
	}
	message := dep.AppMessage{Message: msg.(string)}
	lambda.Start(message.HandleRequest)
}
