package uc

import (
	cc "github.com/ocasti/fuego-quazar/common/contracts"
	"github.com/ocasti/fuego-quazar/topsecret-split/v1/internal/contracts"
	"github.com/ocasti/fuego-quazar/topsecret-split/v1/internal/model"
	"github.com/ocasti/fuego-quazar/topsecret-split/v1/internal/repository"
	"strings"
)

type PostLocationSatelliteUC struct {
	SatelliteRepository contracts.SatelliteRepository
}

func (p *PostLocationSatelliteUC) Handler(request cc.SatelliteDetail) error {

	satellite := model.Satellite{
		SatelliteName: request.Name,
		Distance:      request.Distance,
		Message:       strings.Join(request.Message, ","),
	}

	return p.SatelliteRepository.Save(satellite)
}

func NewPostLocationSatelliteUC(satelliteRepository *repository.SatelliteRepository) *PostLocationSatelliteUC {
	return &PostLocationSatelliteUC{
		SatelliteRepository: satelliteRepository,
	}
}
