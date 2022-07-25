package model

type Satellite struct {
	SatelliteName string  `json:"satellite_name"`
	Distance      float64 `json:"distance"`
	Message       string  `json:"Message"`
}

func (s *Satellite) IsEmpty() bool {
	return s == &Satellite{}
}