package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var nRPCIConstraints = per.IntegerConstraints{
	Extensible: true,
	LowerBound: common.Ptr(int64(0)),
	UpperBound: common.Ptr(int64(1007)),
}

type NRPCI struct {
	Value int64
}

func (ie *NRPCI) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, nRPCIConstraints)
}

func (ie *NRPCI) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(nRPCIConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
