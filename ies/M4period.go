package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	M4periodMs1024  int64 = 0
	M4periodMs2048  int64 = 1
	M4periodMs5120  int64 = 2
	M4periodMs10240 int64 = 3
	M4periodMin1    int64 = 4
)

var m4periodConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2, 3, 4},
	ExtValues:  nil,
}

type M4period struct {
	Value int64
}

func (ie *M4period) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, m4periodConstraints)
}

func (ie *M4period) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(m4periodConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
