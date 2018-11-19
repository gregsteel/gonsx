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
	XMLName        xml.Name   `xml:"edge"`
	Id             string     `xml:"id,omitempty"`
	DatacenterMoid string     `xml:"datacenterMoid"`
	Name           string     `xml:"name"`
	Type           string     `xml:"type,omitempty"`
	Tenant         string     `xml:"tenant,omitempty"`
	Fqdn           string     `xml:"fqdn,omitempty"`
	VseLogLevel    string     `xml:"vseLogLevel,omitempty"`
	EnableAesni    bool       `xml:"enableAesni,omitempty"`
	EnableFips     bool       `xml:"enableFips,omitempty"`
	Appliances     Appliances `xml:"appliances"`
	Vnics          EdgeVnics  `xml:"vnics"`
	Features       Features   `xml:"features"`
}

// Appliances within EdgeGateway.
type Appliances struct {
	ApplianceSize  string      `xml:"applianceSize"`
	EnableCoreDump bool        `xml:"enableCoreDump,omitempty"`
	ApplianceList  []Appliance `xml:"appliance"`
}

// Appliance object within Appliance list.
type Appliance struct {
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

// Features within EdgeGateway.
type Features struct {
	//lots of other features to be added
	Firewall Firewall `xml:"firewall"`
	Routing  Routing  `xml:"routing"`
	Dhcp     Dhcp     `xml:"dhcp"`
}

// Firewall object within Edge features list.
type Firewall struct {
	DefaultPolicy FirewallDefaultPolicy `xml:"defaultPolicy"`
	GlobalConfig  FirewallGlobalConfig  `xml:"globalConfig"`
}

// FirewallDefaultPolicy ...
type FirewallDefaultPolicy struct {
	Action         string `xml:"action"`
	LoggingEnabled bool   `xml:"loggingEnabled"`
}

// FirewallGlobalConfig ...
type FirewallGlobalConfig struct {
	//lots of other config
	TcpTimeoutEstablished int `xml:"tcpTimeoutEstablished"`
}

// Routing object within Edge features list.
type Routing struct {
	GlobalConfig  RoutingGlobalConfig `xml:"routingGlobalConfig"`
	StaticRouting StaticRouting       `xml:"staticRouting"`
	OSPFRouting   OSPFRouting         `xml:"ospf"`
}

//RoutingGlobalConfig ...
type RoutingGlobalConfig struct {
	RouterId string `xml:"routerId"`
	ECMP     bool   `xml:"ecmp"`
	//Logging 	RoutingGlobalConfigLogging 	`xml:"logging"`
}

//StaticRouting ...
type StaticRouting struct {
	DefaultRoute StaticRoutingDefault `xml:"defaultRoute"`
	//StaticRoutes 	StaticRoutes 			`xml:"staticRoutes"`
}

//StaticRoutingDefault ...
type StaticRoutingDefault struct {
	Description    string `xml:"description"`
	Vnic           int    `xml:"vnic"`
	GatewayAddress string `xml:"gatewayAddress"`
	Mtu            int    `xml:"mtu"`
}

//OSPFRouting ...
type OSPFRouting struct {
	Enabled          bool               `xml:"enabled"`
	OSPFAreas        OSPFAreas          `xml:"ospfAreas"`
	OSPFInterfaces   OSPFInterfaces     `xml:"ospfInterfaces"`
	Redistribution   OSPFRedistribution `xml:"redistribution"`
	GracefulRestart  bool               `xml:"gracefulRestart"`
	DefaultOriginate bool               `xml:"defaultOriginate"`
}

//OSPFAreas ...
type OSPFAreas struct {
	OSPFAreaList []OSPFArea `xml:"ospfArea"`
}

//OSPFArea ...
type OSPFArea struct {
	AreaId int    `xml:"areaId"`
	Type   string `xml:"type"`
}

//OSPFInterfaces ...
type OSPFInterfaces struct {
	OSPFInterfaceList []OSPFInterface `xml:"ospfInterface"`
}

//OSPFInterface ...
type OSPFInterface struct {
	Vnic   int `xml:"vnic"`
	AreaId int `xml:"areaId"`
	//more attributes ...
}

//OSPFRedistribution ...
type OSPFRedistribution struct {
	Enabled bool       `xml:"enabled"`
	Rules   []OSPFRule `xml:"rules"`
	//more attributes ...
}

//OSPFRule ...
type OSPFRule struct {
	//??
}

//Dhcp
type Dhcp struct {
	StaticBindings DhcpStaticBindings `xml:"staticBindings"`
	IpPools        DhcpIpPools        `xml:"ipPools"`
}

//DhcpStaticBindings
type DhcpStaticBindings struct {
	StaticBindingList []DhcpStaticBinding `xml:"staticBinding"`
}

//DhcpStaticBinding
type DhcpStaticBinding struct {
	MacAddress          string `xml:"macAddress"`
	VmId                string `xml:"vmId"`
	VnicId              int    `xml:"vnicId"`
	Hostname            string `xml:"hostname"`
	IpAddress           string `xml:"ipAddress"`
	SubnetMask          string `xml:"subnetMask"`
	DefaultGateway      string `xml:"defaultGateway"`
	DomainName          string `xml:"domainName"`
	PrimaryNameServer   string `xml:"primaryNameServer"`
	SecondaryNameServer string `xml:"secondaryNameServer"`
	LeaseTime           string `xml:"leaseTime"`
	AutoConfigureDNS    bool   `xml:"autoConfigureDNS"`
}

//IpPools
type DhcpIpPools struct {
	IpPoolList []DhcpIpPool `xml:"ipPool"`
}

//DhcpIpPool
type DhcpIpPool struct {
	IpRange             string `xml:"ipRange"`
	DefaultGateway      string `xml:"defaultGateway"`
	SubnetMask          string `xml:"subnetMask"`
	DomainName          string `xml:"domainName"`
	PrimaryNameServer   string `xml:"primaryNameServer"`
	SecondaryNameServer string `xml:"secondaryNameServer"`
	LeaseTime           string `xml:"leaseTime"`
	AutoConfigureDNS    bool   `xml:"autoConfigureDNS"`
	NextServer          string `xml:"nextServer,omitempty"`
	//dhcpOptions ...
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
