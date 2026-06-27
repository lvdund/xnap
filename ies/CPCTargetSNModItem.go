package ies

import (
	"github.com/lvdund/asn1go/per"
)

var cPCTargetSNModItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "target-S-NG-RANnodeID"},
		{Name: "candidate-pscells"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type CPCTargetSNModItem struct {
	TargetSNGRANnodeID GlobalNGRANNodeID
	CandidatePscells   CPCInformationUpdatePSCellsList
	IEExtensions       []byte
}

func (ie *CPCTargetSNModItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cPCTargetSNModItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.TargetSNGRANnodeID.Encode(e); err != nil {
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

func (ie *CPCTargetSNModItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cPCTargetSNModItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.TargetSNGRANnodeID.Decode(d); err != nil {
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
