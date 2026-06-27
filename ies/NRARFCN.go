package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var nRARFCNConstraints = per.IntegerConstraints{
	Extensible: false,
	LowerBound: common.Ptr(int64(0)),
	UpperBound: common.Ptr(int64(common.MaxNRARFCN)),
}

type NRARFCN struct {
	Value int64
}

func (ie *NRARFCN) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, nRARFCNConstraints)
}

func (ie *NRARFCN) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(nRARFCNConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
