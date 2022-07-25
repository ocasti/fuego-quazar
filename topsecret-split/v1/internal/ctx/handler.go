package ctx

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	cc "github.com/ocasti/fuego-quazar/common/contracts"
	"github.com/ocasti/fuego-quazar/topsecret-split/v1/internal/contracts"
	"net/http"
)

type Handler struct {
	GetTrilaterationUC      contracts.GetTrilaterationUC
	PostLocationSatelliteUC contracts.PostLocationSatelliteUC
}

func (h *Handler) Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch req.HTTPMethod {
	case "GET":
		data, err := h.GetTrilaterationUC.Handler()
		if err != nil {
			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusNotFound,
				Body:       fmt.Sprintf(`{"message":%s}`, err.Error()),
			}, nil
		}
		body, err := json.Marshal(data)
		if err != nil {
			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusNotFound,
				Body:       fmt.Sprintf(`{"message":%s}`, err.Error()),
			}, nil
		}
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Body:       string(body),
		}, nil
	case "POST":
		request := cc.SatelliteDetail{}
		err := json.Unmarshal([]byte(req.Body), &request)

		if err != nil {
			return generateInternalServerError(err.Error()), nil
		}
		
		satelliteName := req.PathParameters["satellite_name"]
		if satelliteName == "" {
			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Body:       fmt.Sprintf(`{"message":%s}`, "satellite_name is required"),
			}, nil
		}

		request.Name = satelliteName
		err = h.PostLocationSatelliteUC.Handler(request)
		if err != nil {
			return generateInternalServerError(err.Error()), nil
		}
	}
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusMethodNotAllowed,
	}, nil
}

func generateInternalServerError(err string) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       fmt.Sprintf(`{"message":%s}`, err),
	}
}

func NewHandler() *Handler {
	return &Handler{}
}
