package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var expectedUEMovingTrajectoryConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofCellsUEMovingTrajectory)),
}

type ExpectedUEMovingTrajectory struct {
	Value []*ExpectedUEMovingTrajectoryItem
}

func (ie *ExpectedUEMovingTrajectory) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(expectedUEMovingTrajectoryConstraints)
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

func (ie *ExpectedUEMovingTrajectory) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(expectedUEMovingTrajectoryConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*ExpectedUEMovingTrajectoryItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(ExpectedUEMovingTrajectoryItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
