package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	TimeToWaitV1s  int64 = 0
	TimeToWaitV2s  int64 = 1
	TimeToWaitV5s  int64 = 2
	TimeToWaitV10s int64 = 3
	TimeToWaitV20s int64 = 4
	TimeToWaitV60s int64 = 5
)

var timeToWaitConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2, 3, 4, 5},
	ExtValues:  nil,
}

type TimeToWait struct {
	Value int64
}

func (ie *TimeToWait) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, timeToWaitConstraints)
}

func (ie *TimeToWait) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(timeToWaitConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
