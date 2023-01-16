package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"testing"

	"github.com/joho/godotenv"
)

func goDotEnvVariable(key string) string {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("Error loading ../.env file")
	}

	return os.Getenv(key)
}
func getAwsCredentials() AwsCredentials {
	awsCred := AwsCredentials{
		accessKeyID:     goDotEnvVariable("AWS_ACCESS_KEY_ID"),
		secretAccessKey: goDotEnvVariable("AWS_SECRET_ACCESS_KEY"),
	}
	return awsCred
}
func TestGetAuthenticatedUserEmail(t *testing.T) {
	awsCred := getAwsCredentials()
	dynamoDBClient, _ := GetDynamoDBClient(awsCred)

	actualEmail, _ := GetAuthenticatedUserEmail("token1", dynamoDBClient)
	expectedEmail := "useremail1@gmail.com"
	if actualEmail != expectedEmail {
		t.Errorf("Actual:%q\tExpected %q", actualEmail, expectedEmail)
	}
}

func TestAuthenticationHeader400(t *testing.T) {
	// API gateway endpoint
	apiGatewayEndpoint := "https://djte904663.execute-api.ap-southeast-1.amazonaws.com/api"
	// create request
	resp, err := http.Get(apiGatewayEndpoint)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if resp.StatusCode != 400 {
		t.Errorf("Actual:%q\tExpected %q", resp.StatusCode, 400)
	}
	// check body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	expectedBody := "authentication header is missing"

	if string(body) != expectedBody {
		t.Errorf("Actual:%q\tExpected %q", string(body), expectedBody)
	}

}

func TestInvalidToken403(t *testing.T) {
	// API gateway endpoint
	apiGatewayEndpoint := "https://djte904663.execute-api.ap-southeast-1.amazonaws.com/api"
	// create request
	client := &http.Client{}
	req, err := http.NewRequest("GET", apiGatewayEndpoint, nil)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	headers := http.Header{
		"authentication": []string{"invalidtoken"},
	}
	req.Header = headers

	// fmt.Printf("Request Values: %v\n", req.Header)
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	defer resp.Body.Close()

	// fmt.Printf("StatusCode: %v", resp.StatusCode)
	if resp.StatusCode != 403 {
		t.Errorf("Actual:%q\tExpected %q", resp.StatusCode, 400)
	}
	// check body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	// fmt.Printf("Body: %v", string(body))
	expectedBody := "invalid token"
	if string(body) != expectedBody {
		t.Errorf("Actual:%q\tExpected %q", string(body), expectedBody)
	}

}

func TestValidToken(t *testing.T) {
	// API gateway endpoint
	apiGatewayEndpoint := "https://djte904663.execute-api.ap-southeast-1.amazonaws.com/api"
	// create request
	client := &http.Client{}
	req, err := http.NewRequest("GET", apiGatewayEndpoint, nil)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	headers := http.Header{
		"authentication": []string{"token1"},
	}
	req.Header = headers

	// fmt.Printf("Request Values: %v\n", req.Header)
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Actual:%q\tExpected %q", strconv.Itoa(resp.StatusCode), "200")
	}
	// check body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// fmt.Printf("Body: %v", string(body))
	expectedToken1Results, err := os.ReadFile("expectedToken1Results.json")
	if err != nil {
		log.Fatalln(err)
	}
	// convert to string
	buffer := new(bytes.Buffer)
	if err := json.Compact(buffer, expectedToken1Results); err != nil {
		fmt.Println(err)
	}
	// convert buffer to string

	expectedBody := buffer.String()
	if string(body) != expectedBody {
		t.Errorf("Actual:%q\tExpected %q", string(body), expectedBody)
	}

}
