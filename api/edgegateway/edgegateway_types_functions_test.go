//Work in progress, needs to be completed

package edgegateway

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func constructEdgeGateway(edgeId string, name string) (edgeGateway *EdgeGateway) {

	edgeGateway = &EdgeGateway{
		Id:             edgeId,
		Name:           name,
		DatacenterMoid: "dc-1",
		Tenant:         "default",
		Appliances: EdgeGatewayAppliances{
			ApplianceSize:  "compact",
			EnableCoreDump: true,
			ApplianceList: []EdgeGatewayAppliance{
				EdgeGatewayAppliance{
					ResourcePoolId: "pool-1",
					DatastoreId:    "datastore-1",
				},
			},
		},
	}
	return edgeGateway
}

func TestMarshalToXML(t *testing.T) {
	edgeGateway := constructEdgeGateway("edge-1", "Edge Gateway Test")
	convertedXML := edgeGateway.MarshalToXML()
	expectedXML := "<edge><id>edge-1</id><datacenterMoid>dc-1</datacenterMoid><name>Edge Gateway Test</name><tenant>default</tenant><appliances><applianceSize>compact</applianceSize><enableCoreDump>true</enableCoreDump><appliance><resourcePoolId>pool-1</resourcePoolId><datastoreId>datastore-1</datastoreId><vmFolderId></vmFolderId></appliance></appliances><vnics></vnics></edge>"
	assert.Equal(t, expectedXML, convertedXML)
}

func TestStringImplementation(t *testing.T) {
	edgeGateway := constructEdgeGateway("edge-1", "Edge Gateway Test")

	edgeGatewayString := fmt.Sprintln(edgeGateway)
	assert.Equal(t, "name: Edge Gateway Test\n", edgeGatewayString)

}
