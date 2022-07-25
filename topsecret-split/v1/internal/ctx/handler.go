package ctx

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	commonContracts "github.com/ocasti/fuego-quazar/common/contracts"
	"github.com/ocasti/fuego-quazar/topsecret-split/v1/internal/contracts"
	"net/http"
)

type Handler struct {
	GetTrilaterationUC      contracts.GetTrilaterationUC
	PostLocationSatelliteUC contracts.PostLocationSatelliteUC
}

func (h *Handler) Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch req.HTTPMethod {

	case http.MethodGet:
		x, y, msg, err := h.GetTrilaterationUC.Handler()

		if err != nil {
			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusNotFound,
				Body:       fmt.Sprintf(`{"message":%s}`, err.Error()),
			}, nil
		}

		response := commonContracts.Response{
			Position: commonContracts.Position{
				X: x,
				Y: y,
			},
			Message: msg,
		}
		body, err := json.Marshal(response)
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
	case http.MethodPost:
		request := contracts.RequestBody{}
		err := json.Unmarshal([]byte(req.Body), &request)
		if err != nil {
			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Body:       fmt.Sprintf(`{"message":%s}`, err.Error()),
			}, nil
		}
		satelliteName := req.PathParameters["satellite_name"]
		if satelliteName == "" {
			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Body:       fmt.Sprintf(`{"message":%s}`, "satellite_name is required"),
			}, nil
		}

		request.SatelliteName = satelliteName
		err = h.PostLocationSatelliteUC.Handler(request)
		if err != nil {
			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Body:       fmt.Sprintf(`{"message":%s}`, err.Error()),
			}, nil
		}
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
		}, nil
	}
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusMethodNotAllowed,
	}, nil
}

func NewHandler(
	getTrilaterationUC contracts.GetTrilaterationUC,
	postLocationSatelliteUC contracts.PostLocationSatelliteUC,
) *Handler {
	return &Handler{
		GetTrilaterationUC:      getTrilaterationUC,
		PostLocationSatelliteUC: postLocationSatelliteUC,
	}
}
