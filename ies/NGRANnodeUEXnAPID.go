package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var nGRANnodeUEXnAPIDConstraints = per.IntegerConstraints{
	Extensible: false,
	LowerBound: common.Ptr(int64(0)),
	UpperBound: common.Ptr(int64(4294967295)),
}

type NGRANnodeUEXnAPID struct {
	Value int64
}

func (ie *NGRANnodeUEXnAPID) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, nGRANnodeUEXnAPIDConstraints)
}

func (ie *NGRANnodeUEXnAPID) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(nGRANnodeUEXnAPIDConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
