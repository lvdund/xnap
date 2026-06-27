package ies

import (
	"github.com/lvdund/asn1go/per"
)

var cHOinformationAddReqAckConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "pCell-ID", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type CHOinformationAddReqAck struct {
	PCellID      *GlobalNGRANCellID
	IEExtensions []byte
}

func (ie *CHOinformationAddReqAck) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cHOinformationAddReqAckConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.PCellID != nil, false}); err != nil {
		return err
	}
	if ie.PCellID != nil {
		if err := ie.PCellID.Encode(e); err != nil {
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

func (ie *CHOinformationAddReqAck) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cHOinformationAddReqAckConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if seq.IsComponentPresent(0) {
		ie.PCellID = new(GlobalNGRANCellID)
		if err := ie.PCellID.Decode(d); err != nil {
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
