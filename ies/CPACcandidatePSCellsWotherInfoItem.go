package ies

import (
	"github.com/lvdund/asn1go/per"
)

var cPACcandidatePSCellsWotherInfoItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "pscell-id"},
		{Name: "s-CPAC-CompleteCandidateConfig-Indicator", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type CPACcandidatePSCellsWotherInfoItem struct {
	PscellId                              NRCGI
	SCPACCompleteCandidateConfigIndicator *CompleteCandidateConfigIndicator
	IEExtensions                          []byte
}

func (ie *CPACcandidatePSCellsWotherInfoItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cPACcandidatePSCellsWotherInfoItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.SCPACCompleteCandidateConfigIndicator != nil, false}); err != nil {
		return err
	}
	if err := ie.PscellId.Encode(e); err != nil {
		return err
	}
	if ie.SCPACCompleteCandidateConfigIndicator != nil {
		if err := ie.SCPACCompleteCandidateConfigIndicator.Encode(e); err != nil {
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

func (ie *CPACcandidatePSCellsWotherInfoItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cPACcandidatePSCellsWotherInfoItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.PscellId.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.SCPACCompleteCandidateConfigIndicator = new(CompleteCandidateConfigIndicator)
		if err := ie.SCPACCompleteCandidateConfigIndicator.Decode(d); err != nil {
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
