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

func (app *App) Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	responseBody := map[string]string{
		"message": "OH HAI! You have this route & Lambda API functions with the ID: " + app.id,
	}

	responseJSON, err := json.Marshal(responseBody)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body: `{"error": "internal server error: "}` + err.Error(),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			StatusCode: http.StatusInternalServerError,
		}, nil
	}

	response := events.APIGatewayProxyResponse{
		Body: string(responseJSON),
		Headers: map[string]string{
			"Content-Type":                     "application/json",
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Methods":     "GET, POST, PUT, DELETE, OPTIONS",
			"Access-Control-Allow-Headers":     "Content-Type, Authorization, X-Requested-With",
			"Access-Control-Allow-Credentials": "true",
		},
		StatusCode: http.StatusOK,
	}

	return response, nil
}

func main() {

	id := "faddah-tutorial-09876543210"

	app := newApp(id)

	lambda.Start(app.Handler)
}
