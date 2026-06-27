package ies

import (
	"github.com/lvdund/asn1go/per"
)

var dLDiscardingConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "dRBsSubjectToDLDiscarding"},
		{Name: "iE-Extension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type DLDiscarding struct {
	DRBsSubjectToDLDiscarding DRBsSubjectToDLDiscardingList
	IEExtensions              []byte
}

func (ie *DLDiscarding) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dLDiscardingConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.DRBsSubjectToDLDiscarding.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *DLDiscarding) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dLDiscardingConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.DRBsSubjectToDLDiscarding.Decode(d); err != nil {
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
