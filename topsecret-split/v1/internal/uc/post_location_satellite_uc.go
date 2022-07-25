package uc

import (
	"encoding/json"
	"github.com/ocasti/fuego-quazar/topsecret-split/v1/internal/contracts"
	"github.com/ocasti/fuego-quazar/topsecret-split/v1/internal/model"
)

type PostLocationSatelliteUC struct {
	SatelliteRepository contracts.SatelliteRepository
}

func (p *PostLocationSatelliteUC) Handler(request contracts.RequestBody) error {
	message, err := json.Marshal(request.Message)
	if err != nil {
		return err
	}
	satellite := model.Satellite{
		SatelliteName: request.SatelliteName,
		Distance:      request.Distance,
		Message:       string(message),
	}

	return p.SatelliteRepository.Save(satellite)

}

func NewPostLocationSatelliteUC(satelliteRepository contracts.SatelliteRepository) *PostLocationSatelliteUC {
	return &PostLocationSatelliteUC{
		SatelliteRepository: satelliteRepository,
	}
}
