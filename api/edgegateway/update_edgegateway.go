package edgegateway

import (
	"github.com/sky-uk/gonsx/api"
	"net/http"
)

// UpdateEdgeGatewayAPI struct
type UpdateEdgeGatewayAPI struct {
	*api.BaseAPI
}

// NewUpdate returns a new delete method object of UpdateEdgeGatewayAPI
func NewUpdate(edgegatewayID string, edgegateway EdgeGateway) *UpdateEdgeGatewayAPI {
	this := new(UpdateEdgeGatewayAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodPut, "/api/4.0/edges/"+edgegatewayID, edgegateway, nil)
	return this
}
