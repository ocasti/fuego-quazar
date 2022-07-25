package repository

import (
	"errors"
	"log"
	"net"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/ory/dockertest"
)

func postfix() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}

func portActive(network, address string, max int) error {
	for i := 0; i < max; i++ {
		s, err := net.Dial(network, address)
		if err == nil {
			s.Close()
			return nil
		}
		time.Sleep(10 * time.Millisecond)
	}
	return errors.New("port is not open")
}

func dynamodbServerStart(t *testing.T) (func(), *dynamodb.DynamoDB) {

	os.Setenv("AWS_REGION", "us-east-1")
	os.Setenv("AWS_ACCESS_KEY_ID", "x")
	os.Setenv("AWS_SECRET_ACCESS_KEY", "x")

	dynamodbURL := os.Getenv("DYNAMODB_URL")

	closer := func() {}

	if dynamodbURL == "" {
		pool, err := dockertest.NewPool("")
		if err != nil {
			log.Fatalf("Could not connect to docker: %s", err)
		}
		// pulls an image, creates a container based on it and runs it
		resource, err := pool.RunWithOptions(&dockertest.RunOptions{
			Repository:   "amazon/dynamodb-local",
			Tag:          "latest",
			ExposedPorts: []string{"8000"},
		})
		if err != nil {
			t.Fatalf("Could not start resource: %s", err)
		}
		err = portActive("tcp", resource.GetHostPort("8000/tcp"), 1000)
		if err != nil {
			t.Fatalf("Could not connect resource: %s", resource.GetHostPort("8000/tcp"))
		}
		dynamodbURL = "http://" + resource.GetHostPort("8000/tcp")
		closer = func() {
			if err := pool.Purge(resource); err != nil {
				t.Fatal(err)
			}
		}
	}

	session, err := session.NewSession()
	if err != nil {
		t.Errorf("Error while creating dynamodb local server: %v\n", err)
	}
	client := dynamodb.New(session, &aws.Config{Endpoint: aws.String(dynamodbURL)})
	
	return closer, client
}
