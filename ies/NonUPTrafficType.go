package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	NonUPTrafficTypeUeassociatedf1ap    int64 = 0
	NonUPTrafficTypeNonueassociatedf1ap int64 = 1
	NonUPTrafficTypeNonf1               int64 = 2
)

var nonUPTrafficTypeConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2},
	ExtValues:  nil,
}

type NonUPTrafficType struct {
	Value int64
}

func (ie *NonUPTrafficType) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, nonUPTrafficTypeConstraints)
}

func (ie *NonUPTrafficType) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(nonUPTrafficTypeConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
