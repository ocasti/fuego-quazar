package ctx

import (
	"encoding/json"
	"github.com/ocasti/fuego-quazar/topsecret/v1/internal/uc"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/ocasti/fuego-quazar/topsecret/v1/internal/contracts"
)

type Handler struct {
}

func (h *Handler) Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	r := contracts.RequestBody{}
	err := json.Unmarshal([]byte(req.Body), &r)
	if err != nil {
		return events.APIGatewayProxyResponse{}, nil
	}

	var mss [][]string
	var pss []float32

	for _, satellite := range r.Satellites {
		mss = append(mss, satellite.Message)
		pss = append(pss, float32(satellite.Distance))
	}

	message := uc.GetMessage(mss...)

	x, y, err := uc.GetLocation(pss...)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusNotFound,
		}, nil
	}

	var response = contracts.Response{
		Position: contracts.Position{
			X: float64(x),
			Y: float64(y),
		},
		Message: message,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusNotFound,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(jsonResponse),
	}, nil
}

func NewHandler() *Handler {
	return &Handler{}
}
