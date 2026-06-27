package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	DAPSResponseInfoItemDapsResponseIndicatorDapsHOAccepted    int64 = 0
	DAPSResponseInfoItemDapsResponseIndicatorDapsHONotAccepted int64 = 1
)

var dAPSResponseInfoItemDapsResponseIndicatorConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type DAPSResponseInfoItemDapsResponseIndicator struct {
	Value int64
}

func (ie *DAPSResponseInfoItemDapsResponseIndicator) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, dAPSResponseInfoItemDapsResponseIndicatorConstraints)
}

func (ie *DAPSResponseInfoItemDapsResponseIndicator) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(dAPSResponseInfoItemDapsResponseIndicatorConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

var dAPSResponseInfoItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "drbID"},
		{Name: "dapsResponseIndicator"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type DAPSResponseInfoItem struct {
	DrbID                 DRBID
	DapsResponseIndicator DAPSResponseInfoItemDapsResponseIndicator
	IEExtensions          []byte
}

func (ie *DAPSResponseInfoItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dAPSResponseInfoItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.DrbID.Encode(e); err != nil {
		return err
	}
	if err := ie.DapsResponseIndicator.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *DAPSResponseInfoItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dAPSResponseInfoItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.DrbID.Decode(d); err != nil {
		return err
	}
	if err := ie.DapsResponseIndicator.Decode(d); err != nil {
		return err
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
