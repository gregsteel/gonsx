package edgegateway

import (
	"github.com/sky-uk/gonsx/api"
	"net/http"
)

// DeleteEdgeGatewayAPI struct
type DeleteEdgeGatewayAPI struct {
	*api.BaseAPI
}

// NewDelete returns a new delete method object of DeleteEdgeGatewayAPI
func NewDelete(edgegatewayID string) *DeleteEdgeGatewayAPI {
	this := new(DeleteEdgeGatewayAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodDelete, "/api/4.0/edges/"+edgegatewayID, nil, nil)
	return this
}
