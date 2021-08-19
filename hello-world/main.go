package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Message struct {
	Message string `json:"message"`
}

func handler(res events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var msg string

	switch res.Resource {
	case "": // sam local invoke
	case "/hello":
		msg = "hello world!"

	case "/hello/{name}":
		msg = fmt.Sprintf("hello %s!", res.PathParameters["name"])

	case "/goodbye":
		msg = "goodbye cruel world..."

	default:
		return events.APIGatewayProxyResponse{}, errors.New("oops!")
	}

	s, err := json.Marshal(Message{msg})
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(s),
	}, nil
}

func main() {
	lambda.Start(handler)
}
