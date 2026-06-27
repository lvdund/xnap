package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var measuredUETrajectoryConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofCellsTrajectory)),
}

type MeasuredUETrajectory struct {
	Value []*MeasuredUETrajectoryItem
}

func (ie *MeasuredUETrajectory) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(measuredUETrajectoryConstraints)
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

func (ie *MeasuredUETrajectory) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(measuredUETrajectoryConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*MeasuredUETrajectoryItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(MeasuredUETrajectoryItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
