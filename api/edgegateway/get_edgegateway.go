package edgegateway

import (
	"net/http"

	"github.com/sky-uk/gonsx/api"
)

// GetEdgeGatewayAPI base object.
type GetEdgeGatewayAPI struct {
	*api.BaseAPI
}

// NewGet returns the api object of GetEdgeGatewayAPI
func NewGet(edgegatewayID string) *GetEdgeGatewayAPI {
	this := new(GetEdgeGatewayAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodGet, "/api/4.0/edges/"+edgegatewayID, nil, new(EdgeGateway))
	return this
}

// GetResponse returns ResponseObject of GetEdgeGatewayAPI.
func (getAPI GetEdgeGatewayAPI) GetResponse() *EdgeGateway {
	return getAPI.ResponseObject().(*EdgeGateway)
}
