package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var assistanceInformationQoEMeasConstraints = per.IntegerConstraints{
	Extensible: true,
	LowerBound: common.Ptr(int64(1)),
	UpperBound: common.Ptr(int64(16)),
}

type AssistanceInformationQoEMeas struct {
	Value int64
}

func (ie *AssistanceInformationQoEMeas) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, assistanceInformationQoEMeasConstraints)
}

func (ie *AssistanceInformationQoEMeas) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(assistanceInformationQoEMeasConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
