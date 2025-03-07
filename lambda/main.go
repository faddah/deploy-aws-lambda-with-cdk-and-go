package main

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type App struct {
	id string
}

func newApp(id string) *App {
	return &App{id: id}
}

func (app *App) Handler() error {
	return nil
}

func main() {

	id := "faddah-tutorial-09876543210"

	app := newApp(id)

	lambda.Start(app.Handler)
}
