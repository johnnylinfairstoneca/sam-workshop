package main

import (
	"encoding/json"
	"errors"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	// DefaultHTTPGetAddress Default Address
	DefaultHTTPGetAddress = "https://checkip.amazonaws.com"

	// ErrNoIP No IP found in response
	ErrNoIP = errors.New("no IP in HTTP response")

	// ErrNon200Response non 200 status code in response
	ErrNon200Response = errors.New("non 200 Response found")
)

type OutputMessage struct {
	Message string `json:"message"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	output := OutputMessage{
		Message: "I'm using canary deployments",
	}

	outputStr, _ := json.Marshal(output)

	return events.APIGatewayProxyResponse{
		Body:       string(outputStr),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
