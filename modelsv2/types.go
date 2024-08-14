package models2

import "encoding/xml"

// CPIXRequest represents the CPIX XML request structure
type CPIXRequest struct {
	XMLName        xml.Name       `xml:"cpix:CPIX"`
	ID             string         `xml:"id,attr"`
	XMLNSCPIX      string         `xml:"xmlns:cpix,attr"`
	XMLNSPSKC      string         `xml:"xmlns:pskc,attr"`
	XMLNSSPEKE     string         `xml:"xmlns:speke,attr"`
	XMLNSDS        string         `xml:"xmlns:ds,attr"`
	XMLNSENC       string         `xml:"xmlns:enc,attr"`
	ContentKeyList ContentKeyList `xml:"cpix:ContentKeyList"`
	DRMSystemList  DRMSystemList  `xml:"cpix:DRMSystemList"`
}

type ContentKeyList struct {
	ContentKeys []ContentKey `xml:"cpix:ContentKey"`
}

type ContentKey struct {
	KID string `xml:"kid,attr"`
}

type DRMSystemList struct {
	DRMSystems []DRMSystem `xml:"cpix:DRMSystem"`
}

type DRMSystem struct {
	KID                   string `xml:"kid,attr"`
	SystemID              string `xml:"systemId,attr"`
	ContentProtectionData string `xml:"cpix:ContentProtectionData"`
	ProtectionHeader      string `xml:"speke:ProtectionHeader"`
	PSSH                  string `xml:"cpix:PSSH"`
	URIExtXKey            string `xml:"cpix:URIExtXKey"`
	KeyFormat             string `xml:"speke:KeyFormat"`
	KeyFormatVersions     string `xml:"speke:KeyFormatVersions"`
}

////

type CPIXResponse struct {
	XMLName        xml.Name               `xml:"cpix:CPIX"`
	ID             string                 `xml:"id,attr"`
	ContentKeyList ContentKeyListResponse `xml:"cpix:ContentKeyList"`
	DRMSystemList  DRMSystemListResponse  `xml:"cpix:DRMSystemList"`
}

type ContentKeyListResponse struct {
	ContentKeys []ContentKeyResponse `xml:"cpix:ContentKey"`
}

type ContentKeyResponse struct {
	KID        string       `xml:"kid,attr"`
	ExplicitIV string       `xml:"explicitIV,attr,omitempty"`
	Data       DataResponse `xml:"cpix:Data"`
}

type DataResponse struct {
	Secret SecretResponse `xml:"pskc:Secret"`
}

type SecretResponse struct {
	PlainValue string `xml:"pskc:PlainValue"`
}

type DRMSystemListResponse struct {
	DRMSystems []DRMSystemResponse `xml:"cpix:DRMSystem"`
}

type DRMSystemResponse struct {
	KID                   string `xml:"kid,attr"`
	SystemID              string `xml:"systemId,attr"`
	PSSH                  string `xml:"cpix:PSSH,omitempty"`
	ContentProtectionData string `xml:"cpix:ContentProtectionData,omitempty"`
	URIExtXKey            string `xml:"cpix:URIExtXKey,omitempty"`
	ProtectionHeader      string `xml:"speke:ProtectionHeader,omitempty"`
	KeyFormat             string `xml:"speke:KeyFormat,omitempty"`
	KeyFormatVersions     string `xml:"speke:KeyFormatVersions,omitempty"`
}
