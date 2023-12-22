package main

import (
	"encoding/json"

	"go-serverless-api/api/common"
	"go-serverless-api/models"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var newAuthor models.Author
	err := json.Unmarshal([]byte(request.Body), &newAuthor)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400}, err
	}

	db, err := common.ConnectToDB()
	if err != nil {
		return common.InternalServerError(), err
	}
	author, err := db.CreateAuthor(&newAuthor)
	if err != nil {
		return common.InternalServerError(), err
	}
	authorBts, err := json.Marshal(author)
	if err != nil {
		return common.InternalServerError(), err
	}
	resp := events.APIGatewayProxyResponse{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            string(authorBts),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
