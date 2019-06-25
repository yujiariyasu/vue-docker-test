package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	personName := request.PathParameters["personName"]
	age := request.QueryStringParameters["age"]

	person := PersonResponse{
		PersonName: personName,
		Age:        age,
	}
	jsonBytes, _ := json.Marshal(person)

	return events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Headers": "origin,Accept,Authorization,Content-Type",
			"Content-Type":                 "application/json",
		},
		Body:       string(jsonBytes),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}

type PersonResponse struct {
	PersonName string `json:"personName"`
	Age        string `json:"age"`
}
