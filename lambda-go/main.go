package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// create struct for TokenLookupItem
type TokenLookupItem struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

// create struct for user note
type UserNote struct {
	User string `json:"user"`
	Text string `json:"text"`
	Date string `json:"create_date"`
	ID   string `json:"id"`
}

//	func GetDynamoDBClient() *dynamodb.DynamoDB {
//		sess := session.Must(session.NewSessionWithOptions(session.Options{
//			SharedConfigState: session.SharedConfigEnable,
//		}))
//		svc := dynamodb.New(sess)
//		log.Printf("GetDynamoDBClient func ran")
//		return svc
//	}
type AwsCredentials struct {
	accessKeyID     string
	secretAccessKey string
}

func GetDynamoDBClient(params ...interface{}) (*dynamodb.DynamoDB, error) {
	log.Printf("GetDynamoDBClient func ran")
	var sess *session.Session
	var err error
	if params != nil {
		// convert to AwsCredentials
		awsCreds := params[0].(AwsCredentials)
		sess, err = session.NewSession(&aws.Config{
			Region: aws.String("ap-southeast-1"),
			Credentials: credentials.NewStaticCredentials(
				awsCreds.accessKeyID,
				awsCreds.secretAccessKey,
				""),
		})
	} else {
		sess, err = session.NewSession(&aws.Config{
			Region: aws.String("ap-southeast-1"),
		})
	}
	if err != nil {
		return nil, err
	}
	return dynamodb.New(sess), nil
}
func GetAuthenticatedUserEmail(token string, dynamoDBClient *dynamodb.DynamoDB) (email string, ok bool) {
	log.Printf("GetAuthenticatedUserEmail func ran")

	// dynamoDBClient, _ := GetDynamoDBClient()

	tableName := "token-email-lookup"

	result, err := dynamoDBClient.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		AttributesToGet: []*string{
			aws.String("token"),
			aws.String("email"),
		},
		/* ... */
		Key: map[string]*dynamodb.AttributeValue{
			"token": {
				S: aws.String(token),
			},
		},
	},
	)
	if err != nil {
		log.Println("DynamoDB Error!", err)
		return "", false
	}

	// log.Printf("dynamoDB result: %v", result)

	item := TokenLookupItem{
		Email: email,
		Token: token,
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if item.Email == "" {
		log.Println("DynamoDB Error!", err)
		return "", false
	}
	// Validate the given token with one from the database
	// and return user email if the tokens match ...

	// return "", false
	return item.Email, true
}

func QueryUserNotes(email string, dynamoDBClient *dynamodb.DynamoDB) []UserNote {
	// dynamoDBClient, _ := GetDynamoDBClient()

	// User the following date format for "now"
	// dateNow := time.Now().Format(time.RFC3339)

	result, err := dynamoDBClient.Query(&dynamodb.QueryInput{
		TableName: aws.String("user-notes"),
		KeyConditions: map[string]*dynamodb.Condition{
			"user": {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String(email),
					},
				},
			},
		},
		ScanIndexForward: aws.Bool(false), // sort by date descending
		Limit:            aws.Int64(10),
		/* ... */
	})
	if err != nil {
		log.Println("DynamoDB - user-notes Error!", err)
		return nil
	}
	if result.Items == nil {
		log.Println("DynamoDB user-notes empty!", err)
		return nil
	}

	// unmarshal the result
	userNotes := []UserNote{}
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &userNotes)
	if err != nil {
		log.Println("error getting notes", err.Error())
		return nil
	}

	return userNotes
}

func AuthenticateUser(headers map[string]string) (string, error) {
	// You can get Authentication header in the following way:
	// get or values
	authenticationHeaders := []string{
		headers["Authentication"],
		headers["authentication"],
	}
	// get non empty value
	var authenticationHeader string
	for _, v := range authenticationHeaders {
		if v != "" {
			authenticationHeader = v
			break
		}
	}
	// authenticationHeader := headers["Authentication"]
	// get small caps authentication
	// authenticationHeader2 := headers["authentication"]
	// get token from header Authorization

	log.Printf("authenticationHeader: %v", authenticationHeader)
	// if authenticationHeader is missing or malformed
	if authenticationHeader == "" {
		return "", errors.New("authentication header is missing")
	}

	// token := strings.Split(authenticationHeader, "Bearer ")[1]
	// log.Printf("token: %v", token)
	// Validate the Authentication header and retrieve token
	// remove whitespace
	// token := strings.TrimSpace(authenticationHeader)
	token := authenticationHeader
	fmt.Printf("token: %q", token)
	// initiate db
	dynamoDBClient, _ := GetDynamoDBClient()
	email, ok := GetAuthenticatedUserEmail(token, dynamoDBClient)
	log.Printf("email: %v", email)
	if !ok {
		return "", errors.New("invalid token")
	}

	// return email, nil
	return email, nil
}

// Handler is your Lambda function handler
// It uses Amazon API Gateway request/responses provided by the aws-lambda-go/events package,
// However you could use other event sources (S3, Kinesis etc), or JSON-decoded primitive types such as 'string'.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	headers := request.Headers
	for key, value := range headers {
		log.Printf("Header key=%s, value=%s", key, value)
	}

	email, err := AuthenticateUser(headers)

	if err != nil {
		// Return appropriate responses on failed authentication
		if err.Error() == "authentication header is missing" {
			return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 400}, nil
		}
		if err.Error() == "invalid token" {
			return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 403}, nil
		}
	}

	// user notes section
	dynamoDBClient, _ := GetDynamoDBClient()
	userNotes := QueryUserNotes(email, dynamoDBClient)

	var buf bytes.Buffer

	body, err := json.Marshal(userNotes)

	if err != nil {
		return events.APIGatewayProxyResponse{Body: request.Body, StatusCode: 400}, nil
	}

	json.HTMLEscape(&buf, body)

	resp := events.APIGatewayProxyResponse{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
	return resp, nil

	// stdout and stderr are sent to AWS CloudWatch Logs
	// log.Printf("Processing Lambda request %v\n", request.RequestContext)
	// outputString := fmt.Sprintf("Hello world, token: %v", email)
	// return events.APIGatewayProxyResponse{
	// 	Body:       outputString,
	// 	StatusCode: 200,
	// }, nil
}

func main() {
	log.Printf("Start lambda new")
	lambda.Start(Handler)
}
