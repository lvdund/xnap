package ies

import (
	"github.com/lvdund/asn1go/per"
)

var compositeAvailableCapacityConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "cellCapacityClassValue", Optional: true},
		{Name: "capacityValueInfo"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type CompositeAvailableCapacity struct {
	CellCapacityClassValue *CellCapacityClassValue
	CapacityValueInfo      CapacityValueInfo
	IEExtensions           []byte
}

func (ie *CompositeAvailableCapacity) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(compositeAvailableCapacityConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.CellCapacityClassValue != nil, false}); err != nil {
		return err
	}
	if ie.CellCapacityClassValue != nil {
		if err := ie.CellCapacityClassValue.Encode(e); err != nil {
			return err
		}
	}
	if err := ie.CapacityValueInfo.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *CompositeAvailableCapacity) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(compositeAvailableCapacityConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if seq.IsComponentPresent(0) {
		ie.CellCapacityClassValue = new(CellCapacityClassValue)
		if err := ie.CellCapacityClassValue.Decode(d); err != nil {
			return err
		}
	}
	if err := ie.CapacityValueInfo.Decode(d); err != nil {
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
