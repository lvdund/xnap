package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var cellBasedUETrajectoryPredictionConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofCellsTrajectoryPredict)),
}

type CellBasedUETrajectoryPrediction struct {
	Value []*PredictedUETrajectoryItem
}

func (ie *CellBasedUETrajectoryPrediction) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(cellBasedUETrajectoryPredictionConstraints)
	if err := seqOf.EncodeLength(int64(len(ie.Value))); err != nil {
		return err
	}
	for _, item := range ie.Value {
		if err := item.Encode(e); err != nil {
			return err
		}
	}
	return nil
}

func (ie *CellBasedUETrajectoryPrediction) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(cellBasedUETrajectoryPredictionConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*PredictedUETrajectoryItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(PredictedUETrajectoryItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
