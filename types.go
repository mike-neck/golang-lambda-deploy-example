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

func (app AppMessage) HandleRequest(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	request := UserRequest{}
	e := json.Unmarshal([]byte(event.Body), &request)
	if e != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest, Body: "¯\\_(ツ)_/¯"}, e
	}
	log.Println("normal", "request", "context:", ctx, ", request:", request)
	code := request.StatusCode
	if (200 <= code && code < 300) || (400 <= code && code < 500) {
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
