package ies

import (
	"github.com/lvdund/asn1go/per"
)

var tNLCapacityIndicatorConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "dLTNLOfferedCapacity"},
		{Name: "dLTNLAvailableCapacity"},
		{Name: "uLTNLOfferedCapacity"},
		{Name: "uLTNLAvailableCapacity"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type TNLCapacityIndicator struct {
	DLTNLOfferedCapacity   OfferedCapacity
	DLTNLAvailableCapacity AvailableCapacity
	ULTNLOfferedCapacity   OfferedCapacity
	ULTNLAvailableCapacity AvailableCapacity
	IEExtensions           []byte
}

func (ie *TNLCapacityIndicator) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(tNLCapacityIndicatorConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.DLTNLOfferedCapacity.Encode(e); err != nil {
		return err
	}
	if err := ie.DLTNLAvailableCapacity.Encode(e); err != nil {
		return err
	}
	if err := ie.ULTNLOfferedCapacity.Encode(e); err != nil {
		return err
	}
	if err := ie.ULTNLAvailableCapacity.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *TNLCapacityIndicator) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(tNLCapacityIndicatorConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.DLTNLOfferedCapacity.Decode(d); err != nil {
		return err
	}
	if err := ie.DLTNLAvailableCapacity.Decode(d); err != nil {
		return err
	}
	if err := ie.ULTNLOfferedCapacity.Decode(d); err != nil {
		return err
	}
	if err := ie.ULTNLAvailableCapacity.Decode(d); err != nil {
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
