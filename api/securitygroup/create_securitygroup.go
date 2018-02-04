package securitygroup

import (
	"github.com/sky-uk/gonsx/api"
	"net/http"
)

// CreateSecurityGroupAPI api object
type CreateSecurityGroupAPI struct {
	*api.BaseAPI
}

// NewCreate returns a new object of CreateSecurityGroupAPI.
func NewCreate(scopeID, securityGroupName string, dynamicMemberDefinition *DynamicMemberDefinition, memberDefinition *MemberDefinition) *CreateSecurityGroupAPI {
	this := new(CreateSecurityGroupAPI)
	requestPayload := new(SecurityGroup)
	requestPayload.Name = securityGroupName

	if dynamicMemberDefinition.DynamicSet != nil {
		requestPayload.DynamicMemberDefinition = dynamicMemberDefinition
	}
	if memberDefinition.ObjectID != "" {
		requestPayload.MemberDefinition = memberDefinition
	}
	this.BaseAPI = api.NewBaseAPI(http.MethodPost, "/api/2.0/services/securitygroup/bulk/"+scopeID, requestPayload, new(string))
	return this
}

// GetResponse returns a ResponseObject of CreateSecurityGroupAPI.
func (ca CreateSecurityGroupAPI) GetResponse() string {
	return ca.ResponseObject().(string)
}
