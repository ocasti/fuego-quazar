// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/ocasti/fuego-quazar/topsecret-split/v1/internal/ctx"
)

// Injectors from wire.go:

func Initialize() (*ctx.Handler, error) {
	v := awsConfigProvider()
	sessionSession, err := session.NewSession(v...)
	if err != nil {
		return nil, err
	}
	dynamoDB := dynamodb.New(sessionSession, v...)
	satelliteRepository, err := providerSatelliteRepository(dynamoDB)
	if err != nil {
		return nil, err
	}
	getTrilaterationUC, err := providerGetTrilaterationUC(satelliteRepository)
	if err != nil {
		return nil, err
	}
	postLocationSatelliteUC, err := providerPostLocationSatelliteUC(satelliteRepository)
	if err != nil {
		return nil, err
	}
	handler := providerHandler(getTrilaterationUC, postLocationSatelliteUC)
	return handler, nil
}
