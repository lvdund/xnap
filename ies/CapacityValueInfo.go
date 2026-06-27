package ies

import (
	"github.com/lvdund/asn1go/per"
)

var capacityValueInfoConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "capacityValue"},
		{Name: "ssbAreaCapacityValueList", Optional: true},
		{Name: "iE-Extension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type CapacityValueInfo struct {
	CapacityValue            CapacityValue
	SsbAreaCapacityValueList *SSBAreaCapacityValueList
	IEExtensions             []byte
}

func (ie *CapacityValueInfo) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(capacityValueInfoConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.SsbAreaCapacityValueList != nil, false}); err != nil {
		return err
	}
	if err := ie.CapacityValue.Encode(e); err != nil {
		return err
	}
	if ie.SsbAreaCapacityValueList != nil {
		if err := ie.SsbAreaCapacityValueList.Encode(e); err != nil {
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

func (ie *CapacityValueInfo) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(capacityValueInfoConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.CapacityValue.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.SsbAreaCapacityValueList = new(SSBAreaCapacityValueList)
		if err := ie.SsbAreaCapacityValueList.Decode(d); err != nil {
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
