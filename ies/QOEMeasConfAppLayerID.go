package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var qOEMeasConfAppLayerIDConstraints = per.IntegerConstraints{
	Extensible: true,
	LowerBound: common.Ptr(int64(0)),
	UpperBound: common.Ptr(int64(15)),
}

type QOEMeasConfAppLayerID struct {
	Value int64
}

func (ie *QOEMeasConfAppLayerID) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, qOEMeasConfAppLayerIDConstraints)
}

func (ie *QOEMeasConfAppLayerID) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(qOEMeasConfAppLayerIDConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
