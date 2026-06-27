package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var dLSchedulingPDCCHCCEUsageConstraints = per.IntegerConstraints{
	Extensible: false,
	LowerBound: common.Ptr(int64(0)),
	UpperBound: common.Ptr(int64(100)),
}

type DLSchedulingPDCCHCCEUsage struct {
	Value int64
}

func (ie *DLSchedulingPDCCHCCEUsage) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, dLSchedulingPDCCHCCEUsageConstraints)
}

func (ie *DLSchedulingPDCCHCCEUsage) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(dLSchedulingPDCCHCCEUsageConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
