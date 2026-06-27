package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	M6reportIntervalMs120   int64 = 0
	M6reportIntervalMs240   int64 = 1
	M6reportIntervalMs480   int64 = 2
	M6reportIntervalMs640   int64 = 3
	M6reportIntervalMs1024  int64 = 4
	M6reportIntervalMs2048  int64 = 5
	M6reportIntervalMs5120  int64 = 6
	M6reportIntervalMs10240 int64 = 7
	M6reportIntervalMs20480 int64 = 8
	M6reportIntervalMs40960 int64 = 9
	M6reportIntervalMin1    int64 = 10
	M6reportIntervalMin6    int64 = 11
	M6reportIntervalMin12   int64 = 12
	M6reportIntervalMin30   int64 = 13
)

var m6reportIntervalConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13},
	ExtValues:  nil,
}

type M6reportInterval struct {
	Value int64
}

func (ie *M6reportInterval) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, m6reportIntervalConstraints)
}

func (ie *M6reportInterval) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(m6reportIntervalConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
