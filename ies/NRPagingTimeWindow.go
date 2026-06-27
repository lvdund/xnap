package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	NRPagingTimeWindowS1  int64 = 0
	NRPagingTimeWindowS2  int64 = 1
	NRPagingTimeWindowS3  int64 = 2
	NRPagingTimeWindowS4  int64 = 3
	NRPagingTimeWindowS5  int64 = 4
	NRPagingTimeWindowS6  int64 = 5
	NRPagingTimeWindowS7  int64 = 6
	NRPagingTimeWindowS8  int64 = 7
	NRPagingTimeWindowS9  int64 = 8
	NRPagingTimeWindowS10 int64 = 9
	NRPagingTimeWindowS11 int64 = 10
	NRPagingTimeWindowS12 int64 = 11
	NRPagingTimeWindowS13 int64 = 12
	NRPagingTimeWindowS14 int64 = 13
	NRPagingTimeWindowS15 int64 = 14
	NRPagingTimeWindowS16 int64 = 15
	NRPagingTimeWindowS17 int64 = 16
	NRPagingTimeWindowS18 int64 = 17
	NRPagingTimeWindowS19 int64 = 18
	NRPagingTimeWindowS20 int64 = 19
	NRPagingTimeWindowS21 int64 = 20
	NRPagingTimeWindowS22 int64 = 21
	NRPagingTimeWindowS23 int64 = 22
	NRPagingTimeWindowS24 int64 = 23
	NRPagingTimeWindowS25 int64 = 24
	NRPagingTimeWindowS26 int64 = 25
	NRPagingTimeWindowS27 int64 = 26
	NRPagingTimeWindowS28 int64 = 27
	NRPagingTimeWindowS29 int64 = 28
	NRPagingTimeWindowS30 int64 = 29
	NRPagingTimeWindowS31 int64 = 30
	NRPagingTimeWindowS32 int64 = 31
)

var nRPagingTimeWindowConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
	ExtValues:  []int64{16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31},
}

type NRPagingTimeWindow struct {
	Value int64
}

func (ie *NRPagingTimeWindow) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, nRPagingTimeWindowConstraints)
}

func (ie *NRPagingTimeWindow) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(nRPagingTimeWindowConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
