package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var maxNrofRSIndexesToReportConstraints = per.IntegerConstraints{
	Extensible: true,
	LowerBound: common.Ptr(int64(1)),
	UpperBound: common.Ptr(int64(64)),
}

type MaxNrofRSIndexesToReport struct {
	Value int64
}

func (ie *MaxNrofRSIndexesToReport) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, maxNrofRSIndexesToReportConstraints)
}

func (ie *MaxNrofRSIndexesToReport) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(maxNrofRSIndexesToReportConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
