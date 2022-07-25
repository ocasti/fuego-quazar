package contracts

import "github.com/ocasti/fuego-quazar/common/contracts"

type RequestBody struct {
	Satellites []contracts.SatelliteDetail `json:"satellites"`
}
