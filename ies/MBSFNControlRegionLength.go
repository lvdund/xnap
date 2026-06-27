package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var mBSFNControlRegionLengthConstraints = per.IntegerConstraints{
	Extensible: false,
	LowerBound: common.Ptr(int64(0)),
	UpperBound: common.Ptr(int64(3)),
}

type MBSFNControlRegionLength struct {
	Value int64
}

func (ie *MBSFNControlRegionLength) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, mBSFNControlRegionLengthConstraints)
}

func (ie *MBSFNControlRegionLength) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(mBSFNControlRegionLengthConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
