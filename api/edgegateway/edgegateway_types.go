package edgegateway

import (
	"encoding/xml"
)

// EdgeGateways top level xml element
type EdgeGateways struct {
	EdgeGateways []EdgeGateway `xml:"edges"`
}

// EdgeGateway object within EdgeGateways list.
type EdgeGateway struct {
	XMLName        xml.Name              `xml:"edge"`
	Id             string                `xml:"id,omitempty"`
	DatacenterMoid string                `xml:"datacenterMoid"`
	Name           string                `xml:"name"`
	Type           string                `xml:"type,omitempty"`
	Tenant         string                `xml:"tenant,omitempty"`
	Fqdn           string                `xml:"fqdn,omitempty"`
	VseLogLevel    string                `xml:"vseLogLevel,omitempty"`
	EnableAesni    bool                  `xml:"enableAesni,omitempty"`
	EnableFips     bool                  `xml:"enableFips,omitempty"`
	Appliances     EdgeGatewayAppliances `xml:"appliances"`
	Vnics          EdgeVnics             `xml:"vnics"`
}

// Appliances within EdgeGateway.
type EdgeGatewayAppliances struct {
	ApplianceSize  string                 `xml:"applianceSize"`
	EnableCoreDump bool                   `xml:"enableCoreDump,omitempty"`
	ApplianceList  []EdgeGatewayAppliance `xml:"appliance"`
}

// Appliance object within Appliance list.
type EdgeGatewayAppliance struct {
	XMLName        xml.Name `xml:"appliance"`
	ResourcePoolId string   `xml:"resourcePoolId"`
	DatastoreId    string   `xml:"datastoreId"`
	VmFolderId     string   `xml:"vmFolderId"`
}

// EdgeVnics top level list object
type EdgeVnics struct {
	VnicList []EdgeVnic `xml:"vnic"`
}

// EdgeVnic object within EdgeVnics list.
type EdgeVnic struct {
	XMLName       xml.Name      `xml:"vnic"`
	Index         int           `xml:"index"`
	Name          string        `xml:"name,omitempty"`
	Type          string        `xml:"type,omitempty"`
	Mtu           int           `xml:"mtu,omitempty"`
	PortgroupId   string        `xml:"portgroupId"`
	IsConnected   bool          `xml:"isConnected"`
	AddressGroups AddressGroups `xml:"addressGroups"`
}

// AddressGroups within EdgeVnic.
type AddressGroups struct {
	AddressGroups []AddressGroup `xml:"addressGroup"`
}

// AddressGroup object within AddressGroup list.
type AddressGroup struct {
	XMLName        xml.Name `xml:"addressGroup"`
	PrimaryAddress string   `xml:"primaryAddress"`
	SubnetMask     string   `xml:"subnetMask"`
}
