package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var priorityLevelQoSConstraints = per.IntegerConstraints{
	Extensible: true,
	LowerBound: common.Ptr(int64(1)),
	UpperBound: common.Ptr(int64(127)),
}

type PriorityLevelQoS struct {
	Value int64
}

func (ie *PriorityLevelQoS) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, priorityLevelQoSConstraints)
}

func (ie *PriorityLevelQoS) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(priorityLevelQoSConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
