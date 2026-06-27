package ies

import (
	"github.com/lvdund/asn1go/per"
)

var sCPACMultiTargetSNItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "target-S-NG-RANnodeID"},
		{Name: "recommendedCandidatePSCells"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type SCPACMultiTargetSNItem struct {
	TargetSNGRANnodeID          GlobalNGRANNodeID
	RecommendedCandidatePSCells []byte
	IEExtensions                []byte
}

func (ie *SCPACMultiTargetSNItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sCPACMultiTargetSNItemConstraints)
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
	if err := e.EncodeOctetString(ie.RecommendedCandidatePSCells, per.SizeConstraints{
		Extensible: false,
		Min:        nil,
		Max:        nil,
	}); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *SCPACMultiTargetSNItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sCPACMultiTargetSNItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.TargetSNGRANnodeID.Decode(d); err != nil {
		return err
	}
	{
		val, err := d.DecodeOctetString(per.SizeConstraints{
			Extensible: false,
			Min:        nil,
			Max:        nil,
		})
		if err != nil {
			return err
		}
		ie.RecommendedCandidatePSCells = val
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
