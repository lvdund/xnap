package ies

import (
	"github.com/lvdund/asn1go/per"
)

var cHOinformationAckConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "requestedTargetCellGlobalID"},
		{Name: "maxCHOoperations", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type CHOinformationAck struct {
	RequestedTargetCellGlobalID TargetCGI
	MaxCHOoperations            *MaxCHOpreparations
	IEExtensions                []byte
}

func (ie *CHOinformationAck) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cHOinformationAckConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.MaxCHOoperations != nil, false}); err != nil {
		return err
	}
	if err := ie.RequestedTargetCellGlobalID.Encode(e); err != nil {
		return err
	}
	if ie.MaxCHOoperations != nil {
		if err := ie.MaxCHOoperations.Encode(e); err != nil {
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

func (ie *CHOinformationAck) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cHOinformationAckConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.RequestedTargetCellGlobalID.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.MaxCHOoperations = new(MaxCHOpreparations)
		if err := ie.MaxCHOoperations.Decode(d); err != nil {
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
