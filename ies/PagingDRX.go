package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	PagingDRXV32   int64 = 0
	PagingDRXV64   int64 = 1
	PagingDRXV128  int64 = 2
	PagingDRXV256  int64 = 3
	PagingDRXV512  int64 = 4
	PagingDRXV1024 int64 = 5
)

var pagingDRXConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2, 3},
	ExtValues:  []int64{4, 5},
}

type PagingDRX struct {
	Value int64
}

func (ie *PagingDRX) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, pagingDRXConstraints)
}

func (ie *PagingDRX) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(pagingDRXConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
