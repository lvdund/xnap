package ies

import (
	"github.com/lvdund/asn1go/per"
)

var cPACInformationModRequiredConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "candidate-pscells"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type CPACInformationModRequired struct {
	CandidatePscells CPACcandidatePSCellsList
	IEExtensions     []byte
}

func (ie *CPACInformationModRequired) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cPACInformationModRequiredConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.CandidatePscells.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *CPACInformationModRequired) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cPACInformationModRequiredConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.CandidatePscells.Decode(d); err != nil {
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
