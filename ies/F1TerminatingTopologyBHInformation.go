package ies

import (
	"github.com/lvdund/asn1go/per"
)

var f1TerminatingTopologyBHInformationConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "f1TerminatingBHInformation-List"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type F1TerminatingTopologyBHInformation struct {
	F1TerminatingBHInformationList F1TerminatingBHInformationList
	IEExtensions                   []byte
}

func (ie *F1TerminatingTopologyBHInformation) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(f1TerminatingTopologyBHInformationConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.F1TerminatingBHInformationList.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *F1TerminatingTopologyBHInformation) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(f1TerminatingTopologyBHInformationConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.F1TerminatingBHInformationList.Decode(d); err != nil {
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
