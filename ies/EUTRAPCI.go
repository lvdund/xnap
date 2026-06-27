package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var eUTRAPCIConstraints = per.IntegerConstraints{
	Extensible: true,
	LowerBound: common.Ptr(int64(0)),
	UpperBound: common.Ptr(int64(503)),
}

type EUTRAPCI struct {
	Value int64
}

func (ie *EUTRAPCI) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, eUTRAPCIConstraints)
}

func (ie *EUTRAPCI) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(eUTRAPCIConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
