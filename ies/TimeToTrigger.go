package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	TimeToTriggerMs0    int64 = 0
	TimeToTriggerMs40   int64 = 1
	TimeToTriggerMs64   int64 = 2
	TimeToTriggerMs80   int64 = 3
	TimeToTriggerMs100  int64 = 4
	TimeToTriggerMs128  int64 = 5
	TimeToTriggerMs160  int64 = 6
	TimeToTriggerMs256  int64 = 7
	TimeToTriggerMs320  int64 = 8
	TimeToTriggerMs480  int64 = 9
	TimeToTriggerMs512  int64 = 10
	TimeToTriggerMs640  int64 = 11
	TimeToTriggerMs1024 int64 = 12
	TimeToTriggerMs1280 int64 = 13
	TimeToTriggerMs2560 int64 = 14
	TimeToTriggerMs5120 int64 = 15
)

var timeToTriggerConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
	ExtValues:  nil,
}

type TimeToTrigger struct {
	Value int64
}

func (ie *TimeToTrigger) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, timeToTriggerConstraints)
}

func (ie *TimeToTrigger) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(timeToTriggerConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
