package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var eUTRAARFCNConstraints = per.IntegerConstraints{
	Extensible: false,
	LowerBound: common.Ptr(int64(0)),
	UpperBound: common.Ptr(int64(common.MaxEARFCN)),
}

type EUTRAARFCN struct {
	Value int64
}

func (ie *EUTRAARFCN) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, eUTRAARFCNConstraints)
}

func (ie *EUTRAARFCN) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(eUTRAARFCNConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
