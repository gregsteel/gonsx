package edgegateway

import (
	"net/http"
	"strings"

	"github.com/sky-uk/gonsx/api"
)

// CreateEdgeGatewayAPI struct
type CreateEdgeGatewayAPI struct {
	*api.BaseAPI
}

// NewCreate returns a new object of CreateEdgeGatewayAPI
func NewCreate(edgegateway *EdgeGateway) *CreateEdgeGatewayAPI {

	this := new(CreateEdgeGatewayAPI)

	this.BaseAPI = api.NewBaseAPI(http.MethodPost, "/api/4.0/edges", edgegateway, new(string))
	return this
}

// GetResponse returns a ResponseObject of CreateEdgeGatewayAPI.
// NSX does not respond with the element id for edge, instead derive it from Location
func (createAPI CreateEdgeGatewayAPI) GetResponse() string {
	resp := createAPI.ResponseObject().(string)
	if resp == "" {
		elements := strings.Split(createAPI.Location(), "/")
		resp = elements[len(elements)-1]
	}
	return resp
}
