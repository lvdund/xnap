package ies

import (
	"github.com/lvdund/asn1go/per"
)

var uEContextRefAtSNHORequestConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "globalNG-RANNode-ID"},
		{Name: "sN-NG-RANnodeUEXnAPID"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type UEContextRefAtSNHORequest struct {
	GlobalNGRANNodeID   GlobalNGRANNodeID
	SNNGRANnodeUEXnAPID NGRANnodeUEXnAPID
	IEExtensions        []byte
}

func (ie *UEContextRefAtSNHORequest) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uEContextRefAtSNHORequestConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.GlobalNGRANNodeID.Encode(e); err != nil {
		return err
	}
	if err := ie.SNNGRANnodeUEXnAPID.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *UEContextRefAtSNHORequest) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uEContextRefAtSNHORequestConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.GlobalNGRANNodeID.Decode(d); err != nil {
		return err
	}
	if err := ie.SNNGRANnodeUEXnAPID.Decode(d); err != nil {
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
