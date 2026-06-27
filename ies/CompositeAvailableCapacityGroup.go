package ies

import (
	"github.com/lvdund/asn1go/per"
)

var compositeAvailableCapacityGroupConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "compositeAvailableCapacityDownlink"},
		{Name: "compositeAvailableCapacityUplink"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type CompositeAvailableCapacityGroup struct {
	CompositeAvailableCapacityDownlink CompositeAvailableCapacity
	CompositeAvailableCapacityUplink   CompositeAvailableCapacity
	IEExtensions                       []byte
}

func (ie *CompositeAvailableCapacityGroup) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(compositeAvailableCapacityGroupConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.CompositeAvailableCapacityDownlink.Encode(e); err != nil {
		return err
	}
	if err := ie.CompositeAvailableCapacityUplink.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *CompositeAvailableCapacityGroup) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(compositeAvailableCapacityGroupConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.CompositeAvailableCapacityDownlink.Decode(d); err != nil {
		return err
	}
	if err := ie.CompositeAvailableCapacityUplink.Decode(d); err != nil {
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
