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
		Appliances: Appliances{
			ApplianceSize:  "compact",
			EnableCoreDump: true,
			ApplianceList: []Appliance{
				Appliance{
					ResourcePoolId: "pool-1",
					DatastoreId:    "datastore-1",
				},
			},
		},
		Features: Features{
			Firewall: Firewall{
				DefaultPolicy: FirewallDefaultPolicy{
					Action:         "accept",
					LoggingEnabled: false,
				},
				GlobalConfig: FirewallGlobalConfig{
					TcpTimeoutEstablished: 300000,
				},
			},
			Routing: Routing{
				GlobalConfig: RoutingGlobalConfig{
					RouterId: "1.1.1.1",
				},
				StaticRouting: StaticRouting{
					DefaultRoute: StaticRoutingDefault{
						Description:    "defaultRoute",
						Vnic:           0,
						GatewayAddress: "1.1.1.1",
						Mtu:            150,
					},
				},
				OSPFRouting: OSPFRouting{
					Enabled: true,
					OSPFAreas: OSPFAreas{
						[]OSPFArea{
							OSPFArea{
								AreaId: 0,
								Type:   "normal",
							},
						},
					},
					OSPFInterfaces: OSPFInterfaces{
						[]OSPFInterface{
							OSPFInterface{
								Vnic:   0,
								AreaId: 0,
							},
						},
					},
					Redistribution: OSPFRedistribution{
						Enabled: false,
						Rules:   []OSPFRule{},
					},
					GracefulRestart:  true,
					DefaultOriginate: false,
				},
			},
			Dhcp: Dhcp{
				StaticBindings: StaticBindings{
					StaticBindingList: []StaticBinding{
						StaticBinding{
							MacAddress: "abcd",
							VmId:       "vm123",
							IpAddress:  "2.2.2.2",
						},
					},
				},
				IpPools: IpPools{
					IpPoolList: []IpPool{
						IpPool{
							IpRange:          "10.1.1.1-10.2.2.2",
							DefaultGateway:   "1.1.1.1",
							SubnetMask:       "255.255.255.0",
							AutoConfigureDNS: true,
						},
					},
				},
			},
		},
	}
	return edgeGateway
}

func TestMarshalToXML(t *testing.T) {
	edgeGateway := constructEdgeGateway("edge-1", "Edge Gateway Test")
	convertedXML := edgeGateway.MarshalToXML()
	expectedXML := "<edge><id>edge-1</id><datacenterMoid>dc-1</datacenterMoid><name>Edge Gateway Test</name><tenant>default</tenant><appliances><applianceSize>compact</applianceSize><enableCoreDump>true</enableCoreDump><appliance><resourcePoolId>pool-1</resourcePoolId><datastoreId>datastore-1</datastoreId><vmFolderId></vmFolderId></appliance></appliances><vnics></vnics><features><firewall><defaultPolicy><action>accept</action><loggingEnabled>false</loggingEnabled></defaultPolicy><globalConfig><tcpTimeoutEstablished>300000</tcpTimeoutEstablished></globalConfig></firewall><routing><routingGlobalConfig><routerId>1.1.1.1</routerId><ecmp>false</ecmp></routingGlobalConfig><staticRouting><defaultRoute><description>defaultRoute</description><vnic>0</vnic><gatewayAddress>1.1.1.1</gatewayAddress><mtu>150</mtu></defaultRoute></staticRouting><ospf><enabled>true</enabled><ospfAreas><ospfArea><areaId>0</areaId><type>normal</type></ospfArea></ospfAreas><ospfInterfaces><ospfInterface><vnic>0</vnic><areaId>0</areaId></ospfInterface></ospfInterfaces><redistribution><enabled>false</enabled></redistribution><gracefulRestart>true</gracefulRestart><defaultOriginate>false</defaultOriginate></ospf></routing><dhcp><staticBindings><staticBinding><macAddress>abcd</macAddress><vmId>vm123</vmId><vnicId>0</vnicId><hostname></hostname><ipAddress>2.2.2.2</ipAddress><subnetMask></subnetMask><defaultGateway></defaultGateway><domainName></domainName><primaryNameServer></primaryNameServer><secondaryNameServer></secondaryNameServer><leaseTime></leaseTime><autoConfigureDNS>false</autoConfigureDNS></staticBinding></staticBindings><ipPools><ipPool><ipRange>10.1.1.1-10.2.2.2</ipRange><defaultGateway>1.1.1.1</defaultGateway><subnetMask>255.255.255.0</subnetMask><domainName></domainName><primaryNameServer></primaryNameServer><secondaryNameServer></secondaryNameServer><leaseTime>0</leaseTime><autoConfigureDNS>true</autoConfigureDNS></ipPool></ipPools></dhcp></features></edge>"
	assert.Equal(t, expectedXML, convertedXML)
}

func TestStringImplementation(t *testing.T) {
	edgeGateway := constructEdgeGateway("edge-1", "Edge Gateway Test")

	edgeGatewayString := fmt.Sprintln(edgeGateway)
	assert.Equal(t, "name: Edge Gateway Test\n", edgeGatewayString)

}
