package ies

import (
	"github.com/lvdund/asn1go/per"
)

var fullAndShortIRNTIProfileListConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "full-I-RNTI-Profile-List"},
		{Name: "short-I-RNTI-Profile-List"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type FullAndShortIRNTIProfileList struct {
	FullIRNTIProfileList  FullIRNTIProfileList
	ShortIRNTIProfileList ShortIRNTIProfileList
	IEExtensions          []byte
}

func (ie *FullAndShortIRNTIProfileList) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(fullAndShortIRNTIProfileListConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.FullIRNTIProfileList.Encode(e); err != nil {
		return err
	}
	if err := ie.ShortIRNTIProfileList.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *FullAndShortIRNTIProfileList) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(fullAndShortIRNTIProfileListConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.FullIRNTIProfileList.Decode(d); err != nil {
		return err
	}
	if err := ie.ShortIRNTIProfileList.Decode(d); err != nil {
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
