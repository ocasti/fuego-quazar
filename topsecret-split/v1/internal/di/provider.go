package di

import (
	"errors"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/ocasti/fuego-quazar/topsecret-split/v1/internal/ctx"
	"github.com/ocasti/fuego-quazar/topsecret-split/v1/internal/repository"
	"github.com/ocasti/fuego-quazar/topsecret-split/v1/internal/uc"
)

func providerHandler(
	getTrilaterationUC *uc.GetTrilaterationUC,
	postLocationUC *uc.PostLocationSatelliteUC,
) *ctx.Handler {
	return ctx.NewHandler(getTrilaterationUC, postLocationUC)
}

func providerGetTrilaterationUC(satelliteRepository *repository.SatelliteRepository) (*uc.GetTrilaterationUC, error) {
	return uc.NewGetTrilaterationUC(satelliteRepository), nil
}

func providerPostLocationSatelliteUC(satelliteRepository *repository.SatelliteRepository) (*uc.PostLocationSatelliteUC, error) {
	return uc.NewPostLocationSatelliteUC(satelliteRepository), nil
}

func awsConfigProvider() []*aws.Config {
	return nil
}

func providerSatelliteRepository(
	client *dynamodb.DynamoDB,
) (*repository.SatelliteRepository, error) {
	table := os.Getenv("DYNAMODB_TABLE")
	if table == "" {
		return nil, errors.New("DYNAMODB_TABLE is empty")
	}

	return repository.NewSatelliteRepository(
		client,
		table,
	), nil
}
