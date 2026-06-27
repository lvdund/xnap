package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var rFSPIndexConstraints = per.IntegerConstraints{
	Extensible: false,
	LowerBound: common.Ptr(int64(1)),
	UpperBound: common.Ptr(int64(256)),
}

type RFSPIndex struct {
	Value int64
}

func (ie *RFSPIndex) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, rFSPIndexConstraints)
}

func (ie *RFSPIndex) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(rFSPIndexConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
