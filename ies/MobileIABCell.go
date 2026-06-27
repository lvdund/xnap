package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	MobileIABCellTrue int64 = 0
)

var mobileIABCellConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type MobileIABCell struct {
	Value int64
}

func (ie *MobileIABCell) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, mobileIABCellConstraints)
}

func (ie *MobileIABCell) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(mobileIABCellConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
