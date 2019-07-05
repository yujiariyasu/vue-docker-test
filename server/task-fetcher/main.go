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

func index(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	userName := request.PathParameters["userName"]

	session, err := session.NewSession()
	conn := dynamodb.New(session)
	params, err := conn.Query(&dynamodb.QueryInput{
		TableName: aws.String("Task"),
		ExpressionAttributeNames: map[string]*string{
			"#ID":       aws.String("ID"),
			"#Name":     aws.String("Name"),
			"#UserID":   aws.String("UserID"),
			"#UserName": aws.String("UserName"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":userName": {
				S: aws.String(userName),
			},
		},
		KeyConditionExpression: aws.String("#UserName=:userName"),
		ProjectionExpression:   aws.String("#UserID, #ID, #UserName, #Name"),
	})
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println(params)
	tasks := make([]*TaskRes, 0)
	if err := dynamodbattribute.UnmarshalListOfMaps(params.Items, &tasks); err != nil {
		fmt.Println("error: ", err)
	}
	jsonBytes, _ := json.Marshal(tasks)

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
	lambda.Start(index)
}

type TaskRes struct {
	ID       string `json:id`
	Name     string `json:name`
	UserID   string `json:"userID"`
	UserName string `json:"userName"`
}
