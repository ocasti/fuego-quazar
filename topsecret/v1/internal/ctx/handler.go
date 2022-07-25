package ctx

import (
	"encoding/json"
	cc "github.com/ocasti/fuego-quazar/common/contracts"
	"github.com/ocasti/fuego-quazar/common/uc"
	"github.com/ocasti/fuego-quazar/topsecret/v1/internal/contracts"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
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

	var response = cc.Response{
		Position: cc.Position{
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
