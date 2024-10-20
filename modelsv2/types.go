package models2

import "encoding/xml"

// CPIXRequest represents the CPIX XML request structure
type CPIXRequest struct {
	ContentID               string                  `xml:"contentId,attr"`
	Version                 string                  `xml:"version,attr"`
	XMLName                 xml.Name                `xml:"cpix:CPIX"`
	XMLNSCPIX               string                  `xml:"xmlns:cpix,attr"`
	XMLNSPSKC               string                  `xml:"xmlns:pskc,attr"`
	XMLNSDS                 string                  `xml:"xmlns:ds,attr"`
	XMLNSENC                string                  `xml:"xmlns:enc,attr"`
	ContentKeyList          ContentKeyList          `xml:"cpix:ContentKeyList"`
	DRMSystemList           DRMSystemList           `xml:"cpix:DRMSystemList"`
	ContentKeyPeriodList    ContentKeyPeriodList    `xml:"cpix:ContentKeyPeriodList"`
	ContentKeyUsageRuleList ContentKeyUsageRuleList `xml:"cpix:ContentKeyUsageRuleList"`
}

type ContentKeyList struct {
	ContentKeys []ContentKey `xml:"cpix:ContentKey"`
}

type ContentKey struct {
	ExplicitIV             string `xml:"explicitIV,attr,omitempty"`
	KID                    string `xml:"kid,attr"`
	CommonEncryptionScheme string `xml:"commonEncryptionScheme,attr"`
}

type DRMSystemList struct {
	DRMSystems []DRMSystem `xml:"cpix:DRMSystem"`
}

type DRMSystem struct {
	KID                                 string                               `xml:"kid,attr"`
	SystemID                            string                               `xml:"systemId,attr"`
	PSSH                                string                               `xml:"cpix:PSSH"`
	ContentProtectionData               *ContentProtectionData               `xml:"cpix:ContentProtectionData,omitempty"`
	HLSSignalingData                    []HLSSignalingData                   `xml:"cpix:HLSSignalingData"`
	SmoothStreamingProtectionHeaderData *SmoothStreamingProtectionHeaderData `xml:"cpix:SmoothStreamingProtectionHeaderData,omitempty"`
}

type HLSSignalingData struct {
	Playlist string `xml:"playlist,attr"`
}

type PSSH struct{}

type ContentProtectionData struct{}

type SmoothStreamingProtectionHeaderData struct{}

type ContentKeyPeriodList struct {
	ContentKeyPeriods []ContentKeyPeriod `xml:"cpix:ContentKeyPeriod"`
}

type ContentKeyPeriod struct {
	ID    string `xml:"id,attr"`
	Index string `xml:"index,attr"`
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

type CPIXResponse struct {
	XMLName                 xml.Name                        `xml:"cpix:CPIX"`
	ContentID               string                          `xml:"contentId,attr"`
	Version                 string                          `xml:"version,attr"`
	ContentKeyList          ResponseContentKeyList          `xml:"cpix:ContentKeyList"`
	DRMSystemList           ResponseDRMSystemList           `xml:"cpix:DRMSystemList"`
	ContentKeyPeriodList    ResponseContentKeyPeriodList    `xml:"cpix:ContentKeyPeriodList"`
	ContentKeyUsageRuleList ResponseContentKeyUsageRuleList `xml:"cpix:ContentKeyUsageRuleList"`
}

type ResponseContentKeyList struct {
	ContentKeys []ResponseContentKey `xml:"cpix:ContentKey"`
}

type ResponseContentKey struct {
	KID                    string       `xml:"kid,attr"`
	ExplicitIV             string       `xml:"explicitIV,attr"`
	CommonEncryptionScheme string       `xml:"commonEncryptionScheme,attr"`
	Data                   ResponseData `xml:"cpix:Data"`
}

type ResponseData struct {
	Secret ResponseSecret `xml:"pskc:Secret"`
}

type ResponseSecret struct {
	PlainValue string `xml:"pskc:PlainValue"`
}

type ResponseDRMSystemList struct {
	DRMSystems []DRMSystem `xml:"cpix:DRMSystem"`
}

type ResponseDRMSystem struct {
	SystemID              string                     `xml:"systemId,attr"`
	KID                   string                     `xml:"kid,attr"`
	PSSH                  string                     `xml:"cpix:PSSH"`
	ContentProtectionData string                     `xml:"cpix:ContentProtectionData"`
	HLSSignalingData      []ResponseHLSSignalingData `xml:"cpix:HLSSignalingData"`
	SmoothStreamingData   string                     `xml:"cpix:SmoothStreamingProtectionHeaderData"`
}

type ResponseHLSSignalingData struct {
	Playlist string `xml:"playlist,attr"`
	Value    string `xml:",chardata"`
}

type ResponseContentKeyPeriodList struct {
	ContentKeyPeriods []ResponseContentKeyPeriod `xml:"cpix:ContentKeyPeriod"`
}

type ResponseContentKeyPeriod struct {
	ID    string `xml:"id,attr"`
	Index string `xml:"index,attr"`
}

type ResponseContentKeyUsageRuleList struct {
	ContentKeyUsageRules []ResponseContentKeyUsageRule `xml:"cpix:ContentKeyUsageRule"`
}

type ResponseContentKeyUsageRule struct {
	KID               string                  `xml:"kid,attr"`
	IntendedTrackType string                  `xml:"intendedTrackType,attr"`
	KeyPeriodFilter   ResponseKeyPeriodFilter `xml:"cpix:KeyPeriodFilter"`
	VideoFilter       ResponseVideoFilter     `xml:"cpix:VideoFilter"`
	AudioFilter       ResponseAudioFilter     `xml:"cpix:AudioFilter"`
}

type ResponseKeyPeriodFilter struct {
	PeriodID string `xml:"periodId,attr"`
}

type ResponseVideoFilter struct {
}

type ResponseAudioFilter struct {
}
