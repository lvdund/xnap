package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	M5periodMs1024  int64 = 0
	M5periodMs2048  int64 = 1
	M5periodMs5120  int64 = 2
	M5periodMs10240 int64 = 3
	M5periodMin1    int64 = 4
)

var m5periodConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2, 3, 4},
	ExtValues:  nil,
}

type M5period struct {
	Value int64
}

func (ie *M5period) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, m5periodConstraints)
}

func (ie *M5period) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(m5periodConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
