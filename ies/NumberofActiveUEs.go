package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var numberofActiveUEsConstraints = per.IntegerConstraints{
	Extensible: true,
	LowerBound: common.Ptr(int64(0)),
	UpperBound: common.Ptr(int64(16777215)),
}

type NumberofActiveUEs struct {
	Value int64
}

func (ie *NumberofActiveUEs) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, numberofActiveUEsConstraints)
}

func (ie *NumberofActiveUEs) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(numberofActiveUEsConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
