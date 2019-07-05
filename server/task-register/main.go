package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func create(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	byteBody := ([]byte)(request.Body)
	taskReq := new(TaskRequest)
	if err := json.Unmarshal(byteBody, taskReq); err != nil {
		fmt.Println("error: ", err)
	}
	userName := taskReq.UserName
	taskName := taskReq.Name

	// DynamoDBへ永続化
	task := Task{
		UserName: userName,
		Name:     taskName,
	}
	av, err := dynamodbattribute.MarshalMap(task)
	if err != nil {
		fmt.Println("error: ", err)
	}

	session, err := session.NewSession()
	conn := dynamodb.New(session)
	params, err := conn.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String("Task"),
		Item:      av,
	})
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println(params)

	return events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Headers": "origin,Accept,Authorization,Content-Type",
			"Content-Type":                 "application/json",
		},
		Body:       string(byteBody),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(create)
}

type TaskRequest struct {
	Name     string `json:"name"`
	UserName string `json:"userName"`
}

type Task struct {
	Name     string
	UserName string
}
