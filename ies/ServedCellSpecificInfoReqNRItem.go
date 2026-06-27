package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	ServedCellSpecificInfoReqNRItemAdditionalMTCListRequestIndicatorAdditionalMTCListRequested int64 = 0
)

var servedCellSpecificInfoReqNRItemAdditionalMTCListRequestIndicatorConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type ServedCellSpecificInfoReqNRItemAdditionalMTCListRequestIndicator struct {
	Value int64
}

func (ie *ServedCellSpecificInfoReqNRItemAdditionalMTCListRequestIndicator) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, servedCellSpecificInfoReqNRItemAdditionalMTCListRequestIndicatorConstraints)
}

func (ie *ServedCellSpecificInfoReqNRItemAdditionalMTCListRequestIndicator) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(servedCellSpecificInfoReqNRItemAdditionalMTCListRequestIndicatorConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

var servedCellSpecificInfoReqNRItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "nRCGI"},
		{Name: "additionalMTCListRequestIndicator", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type ServedCellSpecificInfoReqNRItem struct {
	NRCGI                             NRCGI
	AdditionalMTCListRequestIndicator *ServedCellSpecificInfoReqNRItemAdditionalMTCListRequestIndicator
	IEExtensions                      []byte
}

func (ie *ServedCellSpecificInfoReqNRItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(servedCellSpecificInfoReqNRItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.AdditionalMTCListRequestIndicator != nil, false}); err != nil {
		return err
	}
	if err := ie.NRCGI.Encode(e); err != nil {
		return err
	}
	if ie.AdditionalMTCListRequestIndicator != nil {
		if err := ie.AdditionalMTCListRequestIndicator.Encode(e); err != nil {
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

func (ie *ServedCellSpecificInfoReqNRItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(servedCellSpecificInfoReqNRItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.NRCGI.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.AdditionalMTCListRequestIndicator = new(ServedCellSpecificInfoReqNRItemAdditionalMTCListRequestIndicator)
		if err := ie.AdditionalMTCListRequestIndicator.Decode(d); err != nil {
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
