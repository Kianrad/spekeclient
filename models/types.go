package models

import "encoding/xml"

// CPIXRequest represents the CPIX XML request structure
type CPIXRequest struct {
	XMLName                 xml.Name                `xml:"cpix:CPIX"`
	XMLNsCpix               string                  `xml:"xmlns:cpix,attr"`
	XMLNsPskc               string                  `xml:"xmlns:pskc,attr"`
	XMLNsDs                 string                  `xml:"xmlns:ds,attr"`
	XMLNsEnc                string                  `xml:"xmlns:enc,attr"`
	ContentID               string                  `xml:"contentId,attr"`
	Version                 string                  `xml:"version,attr"`
	ContentKeyList          ContentKeyList          `xml:"cpix:ContentKeyList"`
	DRMSystemList           DRMSystemList           `xml:"cpix:DRMSystemList"`
	ContentKeyPeriodList    ContentKeyPeriodList    `xml:"cpix:ContentKeyPeriodList"`
	ContentKeyUsageRuleList ContentKeyUsageRuleList `xml:"cpix:ContentKeyUsageRuleList"`
}

type ContentKeyList struct {
	ContentKeys []ContentKey `xml:"cpix:ContentKey"`
}

type ContentKey struct {
	KID                    string `xml:"kid,attr"`
	CommonEncryptionScheme string `xml:"commonEncryptionScheme,attr"`
}

type DRMSystemList struct {
	DRMSystems []DRMSystem `xml:"cpix:DRMSystem"`
}

type DRMSystem struct {
	KID                   string             `xml:"kid,attr"`
	SystemID              string             `xml:"systemId,attr"`
	PSSH                  string             `xml:"cpix:PSSH,omitempty"`
	ContentProtectionData string             `xml:"cpix:ContentProtectionData,omitempty"`
	HLSSignalingData      []HLSSignalingData `xml:"cpix:HLSSignalingData"`
}

type HLSSignalingData struct {
	Playlist string `xml:"playlist,attr"`
}

type ContentKeyPeriodList struct {
	ContentKeyPeriods []ContentKeyPeriod `xml:"cpix:ContentKeyPeriod"`
}

type ContentKeyPeriod struct {
	ID    string `xml:"id,attr"`
	Index int    `xml:"index,attr"`
}

type ContentKeyUsageRuleList struct {
	ContentKeyUsageRules []ContentKeyUsageRule `xml:"cpix:ContentKeyUsageRule"`
}

type ContentKeyUsageRule struct {
	KID               string          `xml:"kid,attr"`
	IntendedTrackType string          `xml:"intendedTrackType,attr"`
	KeyPeriodFilter   KeyPeriodFilter `xml:"cpix:KeyPeriodFilter"`
	VideoFilter       *VideoFilter    `xml:"cpix:VideoFilter,omitempty"`
	AudioFilter       *AudioFilter    `xml:"cpix:AudioFilter,omitempty"`
}

type KeyPeriodFilter struct {
	PeriodID string `xml:"periodId,attr"`
}

type VideoFilter struct{}
type AudioFilter struct{}

// CPIXResponse represents the CPIX XML response structure
type CPIXResponse struct {
	XMLName        xml.Name               `xml:"cpix:CPIX"`
	ContentKeyList ContentKeyListResponse `xml:"cpix:ContentKeyList"`
	DRMSystemList  DRMSystemListResponse  `xml:"cpix:DRMSystemList"`
}

type ContentKeyListResponse struct {
	ContentKeys []ContentKeyResponse `xml:"cpix:ContentKey"`
}

type ContentKeyResponse struct {
	KID                    string `xml:"kid,attr"`
	ExplicitIV             string `xml:"explicitIV,attr"`
	CommonEncryptionScheme string `xml:"commonEncryptionScheme,attr"`
	Data                   Data   `xml:"cpix:Data"`
}

type Data struct {
	Secret Secret `xml:"pskc:Secret"`
}

type Secret struct {
	PlainValue string `xml:"pskc:PlainValue"`
}

type DRMSystemListResponse struct {
	DRMSystems []DRMSystemResponse `xml:"cpix:DRMSystem"`
}

type DRMSystemResponse struct {
	SystemID              string             `xml:"systemId,attr"`
	KID                   string             `xml:"kid,attr"`
	PSSH                  string             `xml:"cpix:PSSH"`
	ContentProtectionData string             `xml:"cpix:ContentProtectionData"`
	HLSSignalingData      []HLSSignalingData `xml:"cpix:HLSSignalingData"`
}
