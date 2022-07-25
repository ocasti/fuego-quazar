package repository

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/ocasti/fuego-quazar/topsecret-split/v1/internal/model"
	"testing"
)

func createTable(client *dynamodb.DynamoDB, table string, t *testing.T) {
	_, err := client.CreateTable(&dynamodb.CreateTableInput{
		TableName: aws.String(table),
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("pk"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("sk"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("pk"),
				KeyType:       aws.String("HASH"),
			}, {
				AttributeName: aws.String("sk"),
				KeyType:       aws.String("RANGE"),
			},
		},

		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(5),
			WriteCapacityUnits: aws.Int64(5),
		},
	})
	if err != nil {
		t.Fatalf("Could not create table: %s", err)
	}
}

func TestSatelliteRepository_Save(t *testing.T) {
	closer, client := dynamodbServerStart(t)
	defer closer()

	table := "satellite"
	createTable(client, table, t)

	satellite := []model.Satellite{
		{
			SatelliteName: "skywalker",
			Distance:      100,
			Message:       `{"este", "", "", "secreto"}`,
		},
		{
			SatelliteName: "sato",
			Distance:      100,
			Message:       `{"", "es", "", ""}`,
		},
	}

	repo := NewSatelliteRepository(client, table)

	err := repo.Save(satellite[0])
	if err != nil {
		t.Errorf("error saving data %v", err.Error())
	}
	err = repo.Save(satellite[1])
	if err != nil {
		t.Errorf("error saving data %v", err.Error())
	}

	satellites, err := repo.Get()
	if err != nil {
		t.Errorf("error getting data %v", err.Error())
	}

	if len(satellites) != 2 {
		t.Errorf("error test must return %v elements but return %v", 2, len(satellites))
	}

}
