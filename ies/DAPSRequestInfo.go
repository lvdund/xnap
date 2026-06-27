package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	DAPSRequestInfoDapsIndicatorDapsHORequired int64 = 0
)

var dAPSRequestInfoDapsIndicatorConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type DAPSRequestInfoDapsIndicator struct {
	Value int64
}

func (ie *DAPSRequestInfoDapsIndicator) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, dAPSRequestInfoDapsIndicatorConstraints)
}

func (ie *DAPSRequestInfoDapsIndicator) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(dAPSRequestInfoDapsIndicatorConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

var dAPSRequestInfoConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "dapsIndicator"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type DAPSRequestInfo struct {
	DapsIndicator DAPSRequestInfoDapsIndicator
	IEExtensions  []byte
}

func (ie *DAPSRequestInfo) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dAPSRequestInfoConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.DapsIndicator.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *DAPSRequestInfo) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dAPSRequestInfoConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.DapsIndicator.Decode(d); err != nil {
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
