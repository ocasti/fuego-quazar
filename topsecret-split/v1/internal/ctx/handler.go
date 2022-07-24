package ctx

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/ocasti/fuego-quazar/topsecret-split/v1/internal/contracts"
	"net/http"
)

type Handler struct {
}

func (h *Handler) Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	r := contracts.RequestBody{}
	err := json.Unmarshal([]byte(req.Body), &r)
	if err != nil {
		return events.APIGatewayProxyResponse{}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       "",
	}, nil
}

func NewHandler() *Handler {
	return &Handler{}
}
