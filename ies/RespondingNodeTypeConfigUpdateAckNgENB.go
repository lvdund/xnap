package ies

import (
	"github.com/lvdund/asn1go/per"
)

var respondingNodeTypeConfigUpdateAckNgENBConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "iE-Extension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type RespondingNodeTypeConfigUpdateAckNgENB struct {
	IEExtensions []byte
}

func (ie *RespondingNodeTypeConfigUpdateAckNgENB) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(respondingNodeTypeConfigUpdateAckNgENBConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *RespondingNodeTypeConfigUpdateAckNgENB) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(respondingNodeTypeConfigUpdateAckNgENBConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
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
