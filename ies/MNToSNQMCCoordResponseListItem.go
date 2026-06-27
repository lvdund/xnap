package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	MNToSNQMCCoordResponseListItemQoEConfigSendingPathMn int64 = 0
	MNToSNQMCCoordResponseListItemQoEConfigSendingPathSn int64 = 1
)

var mNToSNQMCCoordResponseListItemQoEConfigSendingPathConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type MNToSNQMCCoordResponseListItemQoEConfigSendingPath struct {
	Value int64
}

func (ie *MNToSNQMCCoordResponseListItemQoEConfigSendingPath) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, mNToSNQMCCoordResponseListItemQoEConfigSendingPathConstraints)
}

func (ie *MNToSNQMCCoordResponseListItemQoEConfigSendingPath) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(mNToSNQMCCoordResponseListItemQoEConfigSendingPathConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

const (
	MNToSNQMCCoordResponseListItemQoEReportingPathResponseAccepted int64 = 0
	MNToSNQMCCoordResponseListItemQoEReportingPathResponseRejected int64 = 1
)

var mNToSNQMCCoordResponseListItemQoEReportingPathResponseConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type MNToSNQMCCoordResponseListItemQoEReportingPathResponse struct {
	Value int64
}

func (ie *MNToSNQMCCoordResponseListItemQoEReportingPathResponse) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, mNToSNQMCCoordResponseListItemQoEReportingPathResponseConstraints)
}

func (ie *MNToSNQMCCoordResponseListItemQoEReportingPathResponse) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(mNToSNQMCCoordResponseListItemQoEReportingPathResponseConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

const (
	MNToSNQMCCoordResponseListItemRVQoEReportingPathResponseAccepted int64 = 0
	MNToSNQMCCoordResponseListItemRVQoEReportingPathResponseRejected int64 = 1
)

var mNToSNQMCCoordResponseListItemRVQoEReportingPathResponseConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type MNToSNQMCCoordResponseListItemRVQoEReportingPathResponse struct {
	Value int64
}

func (ie *MNToSNQMCCoordResponseListItemRVQoEReportingPathResponse) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, mNToSNQMCCoordResponseListItemRVQoEReportingPathResponseConstraints)
}

func (ie *MNToSNQMCCoordResponseListItemRVQoEReportingPathResponse) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(mNToSNQMCCoordResponseListItemRVQoEReportingPathResponseConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

const (
	MNToSNQMCCoordResponseListItemFurtherRVQoEInterestResponseInterested    int64 = 0
	MNToSNQMCCoordResponseListItemFurtherRVQoEInterestResponseNotInterested int64 = 1
)

var mNToSNQMCCoordResponseListItemFurtherRVQoEInterestResponseConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type MNToSNQMCCoordResponseListItemFurtherRVQoEInterestResponse struct {
	Value int64
}

func (ie *MNToSNQMCCoordResponseListItemFurtherRVQoEInterestResponse) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, mNToSNQMCCoordResponseListItemFurtherRVQoEInterestResponseConstraints)
}

func (ie *MNToSNQMCCoordResponseListItemFurtherRVQoEInterestResponse) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(mNToSNQMCCoordResponseListItemFurtherRVQoEInterestResponseConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

const (
	MNToSNQMCCoordResponseListItemFurtherRVQoEReportingPathResponseSrb4 int64 = 0
	MNToSNQMCCoordResponseListItemFurtherRVQoEReportingPathResponseSrb5 int64 = 1
)

var mNToSNQMCCoordResponseListItemFurtherRVQoEReportingPathResponseConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type MNToSNQMCCoordResponseListItemFurtherRVQoEReportingPathResponse struct {
	Value int64
}

func (ie *MNToSNQMCCoordResponseListItemFurtherRVQoEReportingPathResponse) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, mNToSNQMCCoordResponseListItemFurtherRVQoEReportingPathResponseConstraints)
}

func (ie *MNToSNQMCCoordResponseListItemFurtherRVQoEReportingPathResponse) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(mNToSNQMCCoordResponseListItemFurtherRVQoEReportingPathResponseConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

var mNToSNQMCCoordResponseListItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "qOEReference"},
		{Name: "qOEMeasConfigAppLayerID", Optional: true},
		{Name: "qoEConfigSendingPath", Optional: true},
		{Name: "qoEReportingPathResponse", Optional: true},
		{Name: "rVQoEReportingPathResponse", Optional: true},
		{Name: "furtherRVQoEInterestResponse", Optional: true},
		{Name: "furtherRVQoEReportingPathResponse", Optional: true},
		{Name: "preferredRVQoEConfig", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type MNToSNQMCCoordResponseListItem struct {
	QOEReference                      QOEReference
	QOEMeasConfigAppLayerID           *QOEMeasConfAppLayerID
	QoEConfigSendingPath              *MNToSNQMCCoordResponseListItemQoEConfigSendingPath
	QoEReportingPathResponse          *MNToSNQMCCoordResponseListItemQoEReportingPathResponse
	RVQoEReportingPathResponse        *MNToSNQMCCoordResponseListItemRVQoEReportingPathResponse
	FurtherRVQoEInterestResponse      *MNToSNQMCCoordResponseListItemFurtherRVQoEInterestResponse
	FurtherRVQoEReportingPathResponse *MNToSNQMCCoordResponseListItemFurtherRVQoEReportingPathResponse
	PreferredRVQoEConfig              *RVQoEConfig
	IEExtensions                      []byte
}

func (ie *MNToSNQMCCoordResponseListItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mNToSNQMCCoordResponseListItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.QOEMeasConfigAppLayerID != nil, ie.QoEConfigSendingPath != nil, ie.QoEReportingPathResponse != nil, ie.RVQoEReportingPathResponse != nil, ie.FurtherRVQoEInterestResponse != nil, ie.FurtherRVQoEReportingPathResponse != nil, ie.PreferredRVQoEConfig != nil, false}); err != nil {
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
	if ie.QoEConfigSendingPath != nil {
		if err := ie.QoEConfigSendingPath.Encode(e); err != nil {
			return err
		}
	}
	if ie.QoEReportingPathResponse != nil {
		if err := ie.QoEReportingPathResponse.Encode(e); err != nil {
			return err
		}
	}
	if ie.RVQoEReportingPathResponse != nil {
		if err := ie.RVQoEReportingPathResponse.Encode(e); err != nil {
			return err
		}
	}
	if ie.FurtherRVQoEInterestResponse != nil {
		if err := ie.FurtherRVQoEInterestResponse.Encode(e); err != nil {
			return err
		}
	}
	if ie.FurtherRVQoEReportingPathResponse != nil {
		if err := ie.FurtherRVQoEReportingPathResponse.Encode(e); err != nil {
			return err
		}
	}
	if ie.PreferredRVQoEConfig != nil {
		if err := ie.PreferredRVQoEConfig.Encode(e); err != nil {
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

func (ie *MNToSNQMCCoordResponseListItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mNToSNQMCCoordResponseListItemConstraints)
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
		ie.QoEConfigSendingPath = new(MNToSNQMCCoordResponseListItemQoEConfigSendingPath)
		if err := ie.QoEConfigSendingPath.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(3) {
		ie.QoEReportingPathResponse = new(MNToSNQMCCoordResponseListItemQoEReportingPathResponse)
		if err := ie.QoEReportingPathResponse.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(4) {
		ie.RVQoEReportingPathResponse = new(MNToSNQMCCoordResponseListItemRVQoEReportingPathResponse)
		if err := ie.RVQoEReportingPathResponse.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(5) {
		ie.FurtherRVQoEInterestResponse = new(MNToSNQMCCoordResponseListItemFurtherRVQoEInterestResponse)
		if err := ie.FurtherRVQoEInterestResponse.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(6) {
		ie.FurtherRVQoEReportingPathResponse = new(MNToSNQMCCoordResponseListItemFurtherRVQoEReportingPathResponse)
		if err := ie.FurtherRVQoEReportingPathResponse.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(7) {
		ie.PreferredRVQoEConfig = new(RVQoEConfig)
		if err := ie.PreferredRVQoEConfig.Decode(d); err != nil {
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
