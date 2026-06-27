package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	SNGRANnodeAdditionTriggerIndSnChange  int64 = 0
	SNGRANnodeAdditionTriggerIndInterMNHO int64 = 1
	SNGRANnodeAdditionTriggerIndIntraMNHO int64 = 2
)

var sNGRANnodeAdditionTriggerIndConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2},
	ExtValues:  nil,
}

type SNGRANnodeAdditionTriggerInd struct {
	Value int64
}

func (ie *SNGRANnodeAdditionTriggerInd) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, sNGRANnodeAdditionTriggerIndConstraints)
}

func (ie *SNGRANnodeAdditionTriggerInd) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(sNGRANnodeAdditionTriggerIndConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
