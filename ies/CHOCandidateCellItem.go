package ies

import (
	"github.com/lvdund/asn1go/per"
)

var cHOCandidateCellItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "choCandidateCellID"},
		{Name: "choExecutionCondition-List"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type CHOCandidateCellItem struct {
	ChoCandidateCellID        GlobalNGRANCellID
	ChoExecutionConditionList CHOExecutionConditionList
	IEExtensions              []byte
}

func (ie *CHOCandidateCellItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cHOCandidateCellItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.ChoCandidateCellID.Encode(e); err != nil {
		return err
	}
	if err := ie.ChoExecutionConditionList.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *CHOCandidateCellItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cHOCandidateCellItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.ChoCandidateCellID.Decode(d); err != nil {
		return err
	}
	if err := ie.ChoExecutionConditionList.Decode(d); err != nil {
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
