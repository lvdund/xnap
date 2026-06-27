package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	SNToMNQMCCoordRequestListItemQoEReportingPathRequestSrb4 int64 = 0
	SNToMNQMCCoordRequestListItemQoEReportingPathRequestSrb5 int64 = 1
)

var sNToMNQMCCoordRequestListItemQoEReportingPathRequestConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type SNToMNQMCCoordRequestListItemQoEReportingPathRequest struct {
	Value int64
}

func (ie *SNToMNQMCCoordRequestListItemQoEReportingPathRequest) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, sNToMNQMCCoordRequestListItemQoEReportingPathRequestConstraints)
}

func (ie *SNToMNQMCCoordRequestListItemQoEReportingPathRequest) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(sNToMNQMCCoordRequestListItemQoEReportingPathRequestConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

const (
	SNToMNQMCCoordRequestListItemRVQoEReportingPathRequestSrb4 int64 = 0
	SNToMNQMCCoordRequestListItemRVQoEReportingPathRequestSrb5 int64 = 1
)

var sNToMNQMCCoordRequestListItemRVQoEReportingPathRequestConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type SNToMNQMCCoordRequestListItemRVQoEReportingPathRequest struct {
	Value int64
}

func (ie *SNToMNQMCCoordRequestListItemRVQoEReportingPathRequest) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, sNToMNQMCCoordRequestListItemRVQoEReportingPathRequestConstraints)
}

func (ie *SNToMNQMCCoordRequestListItemRVQoEReportingPathRequest) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(sNToMNQMCCoordRequestListItemRVQoEReportingPathRequestConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

const (
	SNToMNQMCCoordRequestListItemFurtherRVQoEInterestInquiryTrue int64 = 0
)

var sNToMNQMCCoordRequestListItemFurtherRVQoEInterestInquiryConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type SNToMNQMCCoordRequestListItemFurtherRVQoEInterestInquiry struct {
	Value int64
}

func (ie *SNToMNQMCCoordRequestListItemFurtherRVQoEInterestInquiry) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, sNToMNQMCCoordRequestListItemFurtherRVQoEInterestInquiryConstraints)
}

func (ie *SNToMNQMCCoordRequestListItemFurtherRVQoEInterestInquiry) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(sNToMNQMCCoordRequestListItemFurtherRVQoEInterestInquiryConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

const (
	SNToMNQMCCoordRequestListItemFurtherRVQoEReportingPathInquiryTrue int64 = 0
)

var sNToMNQMCCoordRequestListItemFurtherRVQoEReportingPathInquiryConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type SNToMNQMCCoordRequestListItemFurtherRVQoEReportingPathInquiry struct {
	Value int64
}

func (ie *SNToMNQMCCoordRequestListItemFurtherRVQoEReportingPathInquiry) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, sNToMNQMCCoordRequestListItemFurtherRVQoEReportingPathInquiryConstraints)
}

func (ie *SNToMNQMCCoordRequestListItemFurtherRVQoEReportingPathInquiry) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(sNToMNQMCCoordRequestListItemFurtherRVQoEReportingPathInquiryConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

const (
	SNToMNQMCCoordRequestListItemConfigReleaseIndicationRvqoe       int64 = 0
	SNToMNQMCCoordRequestListItemConfigReleaseIndicationQoeAndRvqoe int64 = 1
)

var sNToMNQMCCoordRequestListItemConfigReleaseIndicationConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type SNToMNQMCCoordRequestListItemConfigReleaseIndication struct {
	Value int64
}

func (ie *SNToMNQMCCoordRequestListItemConfigReleaseIndication) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, sNToMNQMCCoordRequestListItemConfigReleaseIndicationConstraints)
}

func (ie *SNToMNQMCCoordRequestListItemConfigReleaseIndication) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(sNToMNQMCCoordRequestListItemConfigReleaseIndicationConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

var sNToMNQMCCoordRequestListItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "qOEReference"},
		{Name: "measCollectionEntityIPAddress", Optional: true},
		{Name: "qoEReportingPathRequest", Optional: true},
		{Name: "rVQoEReportingPathRequest", Optional: true},
		{Name: "furtherRVQoEInterestInquiry", Optional: true},
		{Name: "furtherRVQoEReportingPathInquiry", Optional: true},
		{Name: "currentRVQoEConfig", Optional: true},
		{Name: "availableRVQoEMetrics", Optional: true},
		{Name: "configReleaseIndication", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type SNToMNQMCCoordRequestListItem struct {
	QOEReference                     QOEReference
	MeasCollectionEntityIPAddress    *MeasCollectionEntityIPAddress
	QoEReportingPathRequest          *SNToMNQMCCoordRequestListItemQoEReportingPathRequest
	RVQoEReportingPathRequest        *SNToMNQMCCoordRequestListItemRVQoEReportingPathRequest
	FurtherRVQoEInterestInquiry      *SNToMNQMCCoordRequestListItemFurtherRVQoEInterestInquiry
	FurtherRVQoEReportingPathInquiry *SNToMNQMCCoordRequestListItemFurtherRVQoEReportingPathInquiry
	CurrentRVQoEConfig               *RVQoEConfig
	AvailableRVQoEMetrics            *AvailableRVQoEMetrics
	ConfigReleaseIndication          *SNToMNQMCCoordRequestListItemConfigReleaseIndication
	IEExtensions                     []byte
}

func (ie *SNToMNQMCCoordRequestListItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sNToMNQMCCoordRequestListItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.MeasCollectionEntityIPAddress != nil, ie.QoEReportingPathRequest != nil, ie.RVQoEReportingPathRequest != nil, ie.FurtherRVQoEInterestInquiry != nil, ie.FurtherRVQoEReportingPathInquiry != nil, ie.CurrentRVQoEConfig != nil, ie.AvailableRVQoEMetrics != nil, ie.ConfigReleaseIndication != nil, false}); err != nil {
		return err
	}
	if err := ie.QOEReference.Encode(e); err != nil {
		return err
	}
	if ie.MeasCollectionEntityIPAddress != nil {
		if err := ie.MeasCollectionEntityIPAddress.Encode(e); err != nil {
			return err
		}
	}
	if ie.QoEReportingPathRequest != nil {
		if err := ie.QoEReportingPathRequest.Encode(e); err != nil {
			return err
		}
	}
	if ie.RVQoEReportingPathRequest != nil {
		if err := ie.RVQoEReportingPathRequest.Encode(e); err != nil {
			return err
		}
	}
	if ie.FurtherRVQoEInterestInquiry != nil {
		if err := ie.FurtherRVQoEInterestInquiry.Encode(e); err != nil {
			return err
		}
	}
	if ie.FurtherRVQoEReportingPathInquiry != nil {
		if err := ie.FurtherRVQoEReportingPathInquiry.Encode(e); err != nil {
			return err
		}
	}
	if ie.CurrentRVQoEConfig != nil {
		if err := ie.CurrentRVQoEConfig.Encode(e); err != nil {
			return err
		}
	}
	if ie.AvailableRVQoEMetrics != nil {
		if err := ie.AvailableRVQoEMetrics.Encode(e); err != nil {
			return err
		}
	}
	if ie.ConfigReleaseIndication != nil {
		if err := ie.ConfigReleaseIndication.Encode(e); err != nil {
			return err
		}
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *SNToMNQMCCoordRequestListItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sNToMNQMCCoordRequestListItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.QOEReference.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.MeasCollectionEntityIPAddress = new(MeasCollectionEntityIPAddress)
		if err := ie.MeasCollectionEntityIPAddress.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(2) {
		ie.QoEReportingPathRequest = new(SNToMNQMCCoordRequestListItemQoEReportingPathRequest)
		if err := ie.QoEReportingPathRequest.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(3) {
		ie.RVQoEReportingPathRequest = new(SNToMNQMCCoordRequestListItemRVQoEReportingPathRequest)
		if err := ie.RVQoEReportingPathRequest.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(4) {
		ie.FurtherRVQoEInterestInquiry = new(SNToMNQMCCoordRequestListItemFurtherRVQoEInterestInquiry)
		if err := ie.FurtherRVQoEInterestInquiry.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(5) {
		ie.FurtherRVQoEReportingPathInquiry = new(SNToMNQMCCoordRequestListItemFurtherRVQoEReportingPathInquiry)
		if err := ie.FurtherRVQoEReportingPathInquiry.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(6) {
		ie.CurrentRVQoEConfig = new(RVQoEConfig)
		if err := ie.CurrentRVQoEConfig.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(7) {
		ie.AvailableRVQoEMetrics = new(AvailableRVQoEMetrics)
		if err := ie.AvailableRVQoEMetrics.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(8) {
		ie.ConfigReleaseIndication = new(SNToMNQMCCoordRequestListItemConfigReleaseIndication)
		if err := ie.ConfigReleaseIndication.Decode(d); err != nil {
			return err
		}
	}
	extBytes, err := seq.DecodeExtensionAdditions()
	if err != nil {
		return err
	}
	if len(extBytes) > 0 {
		ie.IEExtensions = extBytes[0]
	}
	return nil
}
