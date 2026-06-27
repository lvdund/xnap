package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	EUTRAPagingTimeWindowS1  int64 = 0
	EUTRAPagingTimeWindowS2  int64 = 1
	EUTRAPagingTimeWindowS3  int64 = 2
	EUTRAPagingTimeWindowS4  int64 = 3
	EUTRAPagingTimeWindowS5  int64 = 4
	EUTRAPagingTimeWindowS6  int64 = 5
	EUTRAPagingTimeWindowS7  int64 = 6
	EUTRAPagingTimeWindowS8  int64 = 7
	EUTRAPagingTimeWindowS9  int64 = 8
	EUTRAPagingTimeWindowS10 int64 = 9
	EUTRAPagingTimeWindowS11 int64 = 10
	EUTRAPagingTimeWindowS12 int64 = 11
	EUTRAPagingTimeWindowS13 int64 = 12
	EUTRAPagingTimeWindowS14 int64 = 13
	EUTRAPagingTimeWindowS15 int64 = 14
	EUTRAPagingTimeWindowS16 int64 = 15
)

var eUTRAPagingTimeWindowConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
	ExtValues:  nil,
}

type EUTRAPagingTimeWindow struct {
	Value int64
}

func (ie *EUTRAPagingTimeWindow) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, eUTRAPagingTimeWindowConstraints)
}

func (ie *EUTRAPagingTimeWindow) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(eUTRAPagingTimeWindowConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
