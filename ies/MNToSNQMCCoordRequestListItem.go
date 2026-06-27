package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	MNToSNQMCCoordRequestListItemQoEReportingPathRequestSrb4 int64 = 0
	MNToSNQMCCoordRequestListItemQoEReportingPathRequestSrb5 int64 = 1
)

var mNToSNQMCCoordRequestListItemQoEReportingPathRequestConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type MNToSNQMCCoordRequestListItemQoEReportingPathRequest struct {
	Value int64
}

func (ie *MNToSNQMCCoordRequestListItemQoEReportingPathRequest) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, mNToSNQMCCoordRequestListItemQoEReportingPathRequestConstraints)
}

func (ie *MNToSNQMCCoordRequestListItemQoEReportingPathRequest) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(mNToSNQMCCoordRequestListItemQoEReportingPathRequestConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

const (
	MNToSNQMCCoordRequestListItemRVQoEReportingPathRequestSrb4 int64 = 0
	MNToSNQMCCoordRequestListItemRVQoEReportingPathRequestSrb5 int64 = 1
)

var mNToSNQMCCoordRequestListItemRVQoEReportingPathRequestConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type MNToSNQMCCoordRequestListItemRVQoEReportingPathRequest struct {
	Value int64
}

func (ie *MNToSNQMCCoordRequestListItemRVQoEReportingPathRequest) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, mNToSNQMCCoordRequestListItemRVQoEReportingPathRequestConstraints)
}

func (ie *MNToSNQMCCoordRequestListItemRVQoEReportingPathRequest) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(mNToSNQMCCoordRequestListItemRVQoEReportingPathRequestConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

const (
	MNToSNQMCCoordRequestListItemFurtherRVQoEInterestInquiryTrue int64 = 0
)

var mNToSNQMCCoordRequestListItemFurtherRVQoEInterestInquiryConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type MNToSNQMCCoordRequestListItemFurtherRVQoEInterestInquiry struct {
	Value int64
}

func (ie *MNToSNQMCCoordRequestListItemFurtherRVQoEInterestInquiry) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, mNToSNQMCCoordRequestListItemFurtherRVQoEInterestInquiryConstraints)
}

func (ie *MNToSNQMCCoordRequestListItemFurtherRVQoEInterestInquiry) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(mNToSNQMCCoordRequestListItemFurtherRVQoEInterestInquiryConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

const (
	MNToSNQMCCoordRequestListItemFurtherRVQoEReportingPathInquiryTrue int64 = 0
)

var mNToSNQMCCoordRequestListItemFurtherRVQoEReportingPathInquiryConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type MNToSNQMCCoordRequestListItemFurtherRVQoEReportingPathInquiry struct {
	Value int64
}

func (ie *MNToSNQMCCoordRequestListItemFurtherRVQoEReportingPathInquiry) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, mNToSNQMCCoordRequestListItemFurtherRVQoEReportingPathInquiryConstraints)
}

func (ie *MNToSNQMCCoordRequestListItemFurtherRVQoEReportingPathInquiry) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(mNToSNQMCCoordRequestListItemFurtherRVQoEReportingPathInquiryConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

const (
	MNToSNQMCCoordRequestListItemConfigReleaseIndicationRvqoe       int64 = 0
	MNToSNQMCCoordRequestListItemConfigReleaseIndicationQoeAndRvqoe int64 = 1
)

var mNToSNQMCCoordRequestListItemConfigReleaseIndicationConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type MNToSNQMCCoordRequestListItemConfigReleaseIndication struct {
	Value int64
}

func (ie *MNToSNQMCCoordRequestListItemConfigReleaseIndication) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, mNToSNQMCCoordRequestListItemConfigReleaseIndicationConstraints)
}

func (ie *MNToSNQMCCoordRequestListItemConfigReleaseIndication) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(mNToSNQMCCoordRequestListItemConfigReleaseIndicationConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

var mNToSNQMCCoordRequestListItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "qOEReference"},
		{Name: "qOEMeasConfigAppLayerID", Optional: true},
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

type MNToSNQMCCoordRequestListItem struct {
	QOEReference                     QOEReference
	QOEMeasConfigAppLayerID          *QOEMeasConfAppLayerID
	MeasCollectionEntityIPAddress    *MeasCollectionEntityIPAddress
	QoEReportingPathRequest          *MNToSNQMCCoordRequestListItemQoEReportingPathRequest
	RVQoEReportingPathRequest        *MNToSNQMCCoordRequestListItemRVQoEReportingPathRequest
	FurtherRVQoEInterestInquiry      *MNToSNQMCCoordRequestListItemFurtherRVQoEInterestInquiry
	FurtherRVQoEReportingPathInquiry *MNToSNQMCCoordRequestListItemFurtherRVQoEReportingPathInquiry
	CurrentRVQoEConfig               *RVQoEConfig
	AvailableRVQoEMetrics            *AvailableRVQoEMetrics
	ConfigReleaseIndication          *MNToSNQMCCoordRequestListItemConfigReleaseIndication
	IEExtensions                     []byte
}

func (ie *MNToSNQMCCoordRequestListItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mNToSNQMCCoordRequestListItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.QOEMeasConfigAppLayerID != nil, ie.MeasCollectionEntityIPAddress != nil, ie.QoEReportingPathRequest != nil, ie.RVQoEReportingPathRequest != nil, ie.FurtherRVQoEInterestInquiry != nil, ie.FurtherRVQoEReportingPathInquiry != nil, ie.CurrentRVQoEConfig != nil, ie.AvailableRVQoEMetrics != nil, ie.ConfigReleaseIndication != nil, false}); err != nil {
		return err
	}
	if err := ie.QOEReference.Encode(e); err != nil {
		return err
	}
	if ie.QOEMeasConfigAppLayerID != nil {
		if err := ie.QOEMeasConfigAppLayerID.Encode(e); err != nil {
			return err
		}
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

func (ie *MNToSNQMCCoordRequestListItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mNToSNQMCCoordRequestListItemConstraints)
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
		ie.QOEMeasConfigAppLayerID = new(QOEMeasConfAppLayerID)
		if err := ie.QOEMeasConfigAppLayerID.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(2) {
		ie.MeasCollectionEntityIPAddress = new(MeasCollectionEntityIPAddress)
		if err := ie.MeasCollectionEntityIPAddress.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(3) {
		ie.QoEReportingPathRequest = new(MNToSNQMCCoordRequestListItemQoEReportingPathRequest)
		if err := ie.QoEReportingPathRequest.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(4) {
		ie.RVQoEReportingPathRequest = new(MNToSNQMCCoordRequestListItemRVQoEReportingPathRequest)
		if err := ie.RVQoEReportingPathRequest.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(5) {
		ie.FurtherRVQoEInterestInquiry = new(MNToSNQMCCoordRequestListItemFurtherRVQoEInterestInquiry)
		if err := ie.FurtherRVQoEInterestInquiry.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(6) {
		ie.FurtherRVQoEReportingPathInquiry = new(MNToSNQMCCoordRequestListItemFurtherRVQoEReportingPathInquiry)
		if err := ie.FurtherRVQoEReportingPathInquiry.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(7) {
		ie.CurrentRVQoEConfig = new(RVQoEConfig)
		if err := ie.CurrentRVQoEConfig.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(8) {
		ie.AvailableRVQoEMetrics = new(AvailableRVQoEMetrics)
		if err := ie.AvailableRVQoEMetrics.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(9) {
		ie.ConfigReleaseIndication = new(MNToSNQMCCoordRequestListItemConfigReleaseIndication)
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
