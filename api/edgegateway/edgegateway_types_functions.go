package edgegateway

import (
	"encoding/xml"
	"fmt"
)

func (s EdgeGateways) String() string {
	return fmt.Sprintf("%v", s.EdgeGateways)
}

func (s EdgeGateway) String() string {
	return fmt.Sprintf("name: %s", s.Name)
}

// MarshalToXML converts the object into XML
func (s EdgeGateway) MarshalToXML() string {
	xmlBytes, _ := xml.Marshal(s)
	return string(xmlBytes)
}
