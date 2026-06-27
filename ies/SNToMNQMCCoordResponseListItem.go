package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	SNToMNQMCCoordResponseListItemQoEReportingPathResponseAccepted int64 = 0
	SNToMNQMCCoordResponseListItemQoEReportingPathResponseRejected int64 = 1
)

var sNToMNQMCCoordResponseListItemQoEReportingPathResponseConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type SNToMNQMCCoordResponseListItemQoEReportingPathResponse struct {
	Value int64
}

func (ie *SNToMNQMCCoordResponseListItemQoEReportingPathResponse) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, sNToMNQMCCoordResponseListItemQoEReportingPathResponseConstraints)
}

func (ie *SNToMNQMCCoordResponseListItemQoEReportingPathResponse) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(sNToMNQMCCoordResponseListItemQoEReportingPathResponseConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

const (
	SNToMNQMCCoordResponseListItemRVQoEReportingPathResponseAccepted int64 = 0
	SNToMNQMCCoordResponseListItemRVQoEReportingPathResponseRejected int64 = 1
)

var sNToMNQMCCoordResponseListItemRVQoEReportingPathResponseConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type SNToMNQMCCoordResponseListItemRVQoEReportingPathResponse struct {
	Value int64
}

func (ie *SNToMNQMCCoordResponseListItemRVQoEReportingPathResponse) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, sNToMNQMCCoordResponseListItemRVQoEReportingPathResponseConstraints)
}

func (ie *SNToMNQMCCoordResponseListItemRVQoEReportingPathResponse) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(sNToMNQMCCoordResponseListItemRVQoEReportingPathResponseConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

const (
	SNToMNQMCCoordResponseListItemFurtherRVQoEInterestResponseInterested    int64 = 0
	SNToMNQMCCoordResponseListItemFurtherRVQoEInterestResponseNotInterested int64 = 1
)

var sNToMNQMCCoordResponseListItemFurtherRVQoEInterestResponseConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type SNToMNQMCCoordResponseListItemFurtherRVQoEInterestResponse struct {
	Value int64
}

func (ie *SNToMNQMCCoordResponseListItemFurtherRVQoEInterestResponse) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, sNToMNQMCCoordResponseListItemFurtherRVQoEInterestResponseConstraints)
}

func (ie *SNToMNQMCCoordResponseListItemFurtherRVQoEInterestResponse) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(sNToMNQMCCoordResponseListItemFurtherRVQoEInterestResponseConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

const (
	SNToMNQMCCoordResponseListItemFurtherRVQoEReportingPathResponseSrb4 int64 = 0
	SNToMNQMCCoordResponseListItemFurtherRVQoEReportingPathResponseSrb5 int64 = 1
)

var sNToMNQMCCoordResponseListItemFurtherRVQoEReportingPathResponseConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type SNToMNQMCCoordResponseListItemFurtherRVQoEReportingPathResponse struct {
	Value int64
}

func (ie *SNToMNQMCCoordResponseListItemFurtherRVQoEReportingPathResponse) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, sNToMNQMCCoordResponseListItemFurtherRVQoEReportingPathResponseConstraints)
}

func (ie *SNToMNQMCCoordResponseListItemFurtherRVQoEReportingPathResponse) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(sNToMNQMCCoordResponseListItemFurtherRVQoEReportingPathResponseConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

var sNToMNQMCCoordResponseListItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "qOEReference"},
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

type SNToMNQMCCoordResponseListItem struct {
	QOEReference                      QOEReference
	QoEReportingPathResponse          *SNToMNQMCCoordResponseListItemQoEReportingPathResponse
	RVQoEReportingPathResponse        *SNToMNQMCCoordResponseListItemRVQoEReportingPathResponse
	FurtherRVQoEInterestResponse      *SNToMNQMCCoordResponseListItemFurtherRVQoEInterestResponse
	FurtherRVQoEReportingPathResponse *SNToMNQMCCoordResponseListItemFurtherRVQoEReportingPathResponse
	PreferredRVQoEConfig              *RVQoEConfig
	IEExtensions                      []byte
}

func (ie *SNToMNQMCCoordResponseListItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sNToMNQMCCoordResponseListItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.QoEReportingPathResponse != nil, ie.RVQoEReportingPathResponse != nil, ie.FurtherRVQoEInterestResponse != nil, ie.FurtherRVQoEReportingPathResponse != nil, ie.PreferredRVQoEConfig != nil, false}); err != nil {
		return err
	}
	if err := ie.QOEReference.Encode(e); err != nil {
		return err
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

func (ie *SNToMNQMCCoordResponseListItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sNToMNQMCCoordResponseListItemConstraints)
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
		ie.QoEReportingPathResponse = new(SNToMNQMCCoordResponseListItemQoEReportingPathResponse)
		if err := ie.QoEReportingPathResponse.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(2) {
		ie.RVQoEReportingPathResponse = new(SNToMNQMCCoordResponseListItemRVQoEReportingPathResponse)
		if err := ie.RVQoEReportingPathResponse.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(3) {
		ie.FurtherRVQoEInterestResponse = new(SNToMNQMCCoordResponseListItemFurtherRVQoEInterestResponse)
		if err := ie.FurtherRVQoEInterestResponse.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(4) {
		ie.FurtherRVQoEReportingPathResponse = new(SNToMNQMCCoordResponseListItemFurtherRVQoEReportingPathResponse)
		if err := ie.FurtherRVQoEReportingPathResponse.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(5) {
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
