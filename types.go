package dep

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"log"
	"net/http"
)

type AppMessage struct {
	Message string
}

type UserRequest struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

type AppResponse struct {
	App  string `json:"app"`
	User string `json:"user"`
}

func (app AppMessage) HandleRequest(ctx context.Context, request UserRequest) (events.APIGatewayProxyResponse, error) {
	log.Println("normal", "request", "context:", ctx, ", request:", request)
	code := request.StatusCode
	switch code {
	case http.StatusOK:
	case http.StatusCreated:
	case http.StatusNoContent:
	case http.StatusBadRequest:
	case http.StatusForbidden:
	case http.StatusNotFound:
	case http.StatusConflict:
		return (&AppResponse{App: app.Message, User: request.Message}).Json(code)
	}
	return events.APIGatewayProxyResponse{StatusCode: http.StatusNotFound, Body: "¯\\_(ツ)_/¯"}, nil
}

func (res *AppResponse) Json(statusCode int) (events.APIGatewayProxyResponse, error) {
	bytes, e := json.Marshal(*res)
	if e != nil {
		log.Println("error", "Json", "marshal", e)
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "internal error"}, e
	}
	jsonString := string(bytes)
	return events.APIGatewayProxyResponse{StatusCode: statusCode, Body: jsonString}, nil
}
