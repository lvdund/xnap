package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	PredictedTrajectoryCellInfoChNGRANCellPredicted = 0
	PredictedTrajectoryCellInfoChChoiceExtension    = 1
)

var predictedTrajectoryCellInfoConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "nG-RAN-Cell-Predicted"},
		{Name: "choice-extension"},
	},
	ExtAlternatives: nil,
}

type PredictedTrajectoryCellInfo struct {
	Choice             int
	NGRANCellPredicted *PredictedTrajectoryNGRANCellInfo
	ChoiceExtension    []byte
}

func (ie *PredictedTrajectoryCellInfo) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(predictedTrajectoryCellInfoConstraints)
	switch ie.Choice {
	case 0: // nG-RAN-Cell-Predicted
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err := ie.NGRANCellPredicted.Encode(e); err != nil {
			return err
		}
	case 1: // choice-extension
		if err := choice.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		// TODO encode field ChoiceExtension (kind=ext)
	}
	return nil
}

func (ie *PredictedTrajectoryCellInfo) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(predictedTrajectoryCellInfoConstraints)
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // nG-RAN-Cell-Predicted
		ie.NGRANCellPredicted = new(PredictedTrajectoryNGRANCellInfo)
		if err := ie.NGRANCellPredicted.Decode(d); err != nil {
			return err
		}
	case 1: // choice-extension
		// TODO decode field ChoiceExtension (kind=ext)
	}
	return nil
}
