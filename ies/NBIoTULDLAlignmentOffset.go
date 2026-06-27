package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	NBIoTULDLAlignmentOffsetKhz7dot5   int64 = 0
	NBIoTULDLAlignmentOffsetKhz0       int64 = 1
	NBIoTULDLAlignmentOffsetKhz7dot5_2 int64 = 2
)

var nBIoTULDLAlignmentOffsetConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2},
	ExtValues:  nil,
}

type NBIoTULDLAlignmentOffset struct {
	Value int64
}

func (ie *NBIoTULDLAlignmentOffset) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, nBIoTULDLAlignmentOffsetConstraints)
}

func (ie *NBIoTULDLAlignmentOffset) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(nBIoTULDLAlignmentOffsetConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
