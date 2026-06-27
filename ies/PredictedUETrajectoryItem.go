package ies

import (
	"github.com/lvdund/asn1go/per"
)

var predictedUETrajectoryItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "predictedtrajectoryCellInfo"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type PredictedUETrajectoryItem struct {
	PredictedtrajectoryCellInfo PredictedTrajectoryCellInfo
	IEExtensions                []byte
}

func (ie *PredictedUETrajectoryItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(predictedUETrajectoryItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.PredictedtrajectoryCellInfo.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *PredictedUETrajectoryItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(predictedUETrajectoryItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.PredictedtrajectoryCellInfo.Decode(d); err != nil {
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
