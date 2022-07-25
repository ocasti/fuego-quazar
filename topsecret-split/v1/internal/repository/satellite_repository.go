package repository

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/ocasti/fuego-quazar/topsecret-split/v1/internal/model"
	"strconv"
)

type SatelliteRepository struct {
	client *dynamodb.DynamoDB
	table  string
}

func (s *SatelliteRepository) Hydrate(items []map[string]*dynamodb.AttributeValue) ([]model.Satellite, error) {
	satellites := make([]model.Satellite, len(items))

	for i, item := range items {
		satellites[i].SatelliteName = *item["sk"].S

		if v, ok := item["distance"]; ok {
			value, err := strconv.ParseFloat(*v.N, 64)
			if err != nil {
				return []model.Satellite{}, err
			}
			satellites[i].Distance = value
		}

		if v, ok := item["message"]; ok {
			satellites[i].Message = *v.S
		}

	}

	return satellites, nil
}

func (s SatelliteRepository) Get() ([]model.Satellite, error) {
	out, err := s.client.Query(&dynamodb.QueryInput{
		TableName: aws.String(s.table),
		KeyConditions: map[string]*dynamodb.Condition{
			"pk": {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String("#satellite"),
					},
				},
			},
		},
	})
	if err != nil {
		return []model.Satellite{}, err
	}

	if len(out.Items) == 0 {
		return []model.Satellite{}, nil
	}

	return s.Hydrate(out.Items)
}

func (s SatelliteRepository) Save(satellite model.Satellite) error {
	item := &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"pk": {
				S: aws.String("#satellite"),
			},
			"sk": {
				S: aws.String(satellite.SatelliteName),
			},
			"distance": {
				N: aws.String(fmt.Sprintf("%f", satellite.Distance)),
			},
			"message": {
				S: aws.String(satellite.Message),
			},
		},
		TableName: aws.String(s.table),
	}
	_, err := s.client.PutItem(item)
	return err
}

func NewSatelliteRepository(client *dynamodb.DynamoDB, table string) *SatelliteRepository {
	return &SatelliteRepository{
		client: client,
		table:  table,
	}
}
