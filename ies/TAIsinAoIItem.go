package ies

import (
	"github.com/lvdund/asn1go/per"
)

var tAIsinAoIItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "pLMN-Identity"},
		{Name: "tAC"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type TAIsinAoIItem struct {
	PLMNIdentity PLMNIdentity
	TAC          TAC
	IEExtensions []byte
}

func (ie *TAIsinAoIItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(tAIsinAoIItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.PLMNIdentity.Encode(e); err != nil {
		return err
	}
	if err := ie.TAC.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *TAIsinAoIItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(tAIsinAoIItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.PLMNIdentity.Decode(d); err != nil {
		return err
	}
	if err := ie.TAC.Decode(d); err != nil {
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
