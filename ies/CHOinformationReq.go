package ies

import (
	"github.com/lvdund/asn1go/per"
)

var cHOinformationReqConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "cho-trigger"},
		{Name: "targetNG-RANnodeUEXnAPID", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type CHOinformationReq struct {
	ChoTrigger              CHOtrigger
	TargetNGRANnodeUEXnAPID *NGRANnodeUEXnAPID
	IEExtensions            []byte
}

func (ie *CHOinformationReq) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cHOinformationReqConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.TargetNGRANnodeUEXnAPID != nil, false}); err != nil {
		return err
	}
	if err := ie.ChoTrigger.Encode(e); err != nil {
		return err
	}
	if ie.TargetNGRANnodeUEXnAPID != nil {
		if err := ie.TargetNGRANnodeUEXnAPID.Encode(e); err != nil {
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

func (ie *CHOinformationReq) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cHOinformationReqConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.ChoTrigger.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.TargetNGRANnodeUEXnAPID = new(NGRANnodeUEXnAPID)
		if err := ie.TargetNGRANnodeUEXnAPID.Decode(d); err != nil {
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
