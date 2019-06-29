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
	userID := taskReq.UserID
	userName := taskReq.UserName
	taskID := taskReq.ID
	taskName := taskReq.Name

	// DynamoDBへ永続化
	task := Task{
		UserID:   userID,
		UserName: userName,
		ID:       taskID,
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
	fmt.Println("success!!")

	return events.APIGatewayProxyResponse{
		Body:       string(byteBody),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(create)
}

type TaskRequest struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	UserID   string `json:"userID"`
	UserName string `json:"userName"`
}

type Task struct {
	ID       string
	Name     string
	UserID   string
	UserName string
}
