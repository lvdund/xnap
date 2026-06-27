package ies

import (
	"github.com/lvdund/asn1go/per"
)

var cPACCandidateCellItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "cpacCandidateCellID"},
		{Name: "cpacExecutionCondition-List"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type CPACCandidateCellItem struct {
	CpacCandidateCellID        GlobalNGRANCellID
	CpacExecutionConditionList CPACExecutionConditionList
	IEExtensions               []byte
}

func (ie *CPACCandidateCellItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cPACCandidateCellItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.CpacCandidateCellID.Encode(e); err != nil {
		return err
	}
	if err := ie.CpacExecutionConditionList.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *CPACCandidateCellItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cPACCandidateCellItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.CpacCandidateCellID.Decode(d); err != nil {
		return err
	}
	if err := ie.CpacExecutionConditionList.Decode(d); err != nil {
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
