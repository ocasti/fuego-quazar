package uc

import (
	"github.com/ocasti/fuego-quazar/topsecret/v1/internal/contracts"
	"github.com/savaki/trilateration"
)

func GetLocation(distances ...float32) (x, y float32, err error) {
	kenobiCoordinate := trilateration.Point{
		X: contracts.MapCoordinates["kenobi"].X,
		Y: contracts.MapCoordinates["kenobi"].Y,
		R: float64(distances[0]),
	}

	skywalkerCoordinate := trilateration.Point{
		X: contracts.MapCoordinates["skywalker"].X,
		Y: contracts.MapCoordinates["skywalker"].Y,
		R: float64(distances[1]),
	}

	satoCoordinate := trilateration.Point{
		X: contracts.MapCoordinates["sato"].X,
		Y: contracts.MapCoordinates["sato"].Y,
		R: float64(distances[2]),
	}

	position, err := trilateration.Solve(kenobiCoordinate, skywalkerCoordinate, satoCoordinate)

	if err != nil {
		return 0, 0, err
	}

	return float32(position[0].X), float32(position[0].Y), nil
}
