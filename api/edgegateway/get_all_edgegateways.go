package edgegateway

import (
	"github.com/sky-uk/gonsx/api"
	"net/http"
)

// GetAllEdgeGatewaysAPI base object.
type GetAllEdgeGatewaysAPI struct {
	*api.BaseAPI
}

// NewGetAll returns the api object of GetAllEdgeGatewaysAPI
func NewGetAll(edgegatewayID string) *GetAllEdgeGatewaysAPI {
	this := new(GetAllEdgeGatewaysAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodGet, "/api/4.0/edges", nil, new(EdgeGateways))
	return this
}

// GetResponse returns ResponseObject of GetAllEdgeGatewaysAPI
func (getAllAPI GetAllEdgeGatewaysAPI) GetResponse() *EdgeGateways {
	return getAllAPI.ResponseObject().(*EdgeGateways)
}
