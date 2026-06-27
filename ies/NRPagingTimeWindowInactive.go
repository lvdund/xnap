package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	NRPagingTimeWindowInactiveS1  int64 = 0
	NRPagingTimeWindowInactiveS2  int64 = 1
	NRPagingTimeWindowInactiveS3  int64 = 2
	NRPagingTimeWindowInactiveS4  int64 = 3
	NRPagingTimeWindowInactiveS5  int64 = 4
	NRPagingTimeWindowInactiveS6  int64 = 5
	NRPagingTimeWindowInactiveS7  int64 = 6
	NRPagingTimeWindowInactiveS8  int64 = 7
	NRPagingTimeWindowInactiveS9  int64 = 8
	NRPagingTimeWindowInactiveS10 int64 = 9
	NRPagingTimeWindowInactiveS11 int64 = 10
	NRPagingTimeWindowInactiveS12 int64 = 11
	NRPagingTimeWindowInactiveS13 int64 = 12
	NRPagingTimeWindowInactiveS14 int64 = 13
	NRPagingTimeWindowInactiveS15 int64 = 14
	NRPagingTimeWindowInactiveS16 int64 = 15
	NRPagingTimeWindowInactiveS17 int64 = 16
	NRPagingTimeWindowInactiveS18 int64 = 17
	NRPagingTimeWindowInactiveS19 int64 = 18
	NRPagingTimeWindowInactiveS20 int64 = 19
	NRPagingTimeWindowInactiveS21 int64 = 20
	NRPagingTimeWindowInactiveS22 int64 = 21
	NRPagingTimeWindowInactiveS23 int64 = 22
	NRPagingTimeWindowInactiveS24 int64 = 23
	NRPagingTimeWindowInactiveS25 int64 = 24
	NRPagingTimeWindowInactiveS26 int64 = 25
	NRPagingTimeWindowInactiveS27 int64 = 26
	NRPagingTimeWindowInactiveS28 int64 = 27
	NRPagingTimeWindowInactiveS29 int64 = 28
	NRPagingTimeWindowInactiveS30 int64 = 29
	NRPagingTimeWindowInactiveS31 int64 = 30
	NRPagingTimeWindowInactiveS32 int64 = 31
)

var nRPagingTimeWindowInactiveConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31},
	ExtValues:  nil,
}

type NRPagingTimeWindowInactive struct {
	Value int64
}

func (ie *NRPagingTimeWindowInactive) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, nRPagingTimeWindowInactiveConstraints)
}

func (ie *NRPagingTimeWindowInactive) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(nRPagingTimeWindowInactiveConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
