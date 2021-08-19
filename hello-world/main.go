package main

import (
	"errors"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

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

	return events.APIGatewayProxyResponse{
		Body:       msg,
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
