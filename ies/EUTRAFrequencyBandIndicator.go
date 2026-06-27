package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var eUTRAFrequencyBandIndicatorConstraints = per.IntegerConstraints{
	Extensible: true,
	LowerBound: common.Ptr(int64(1)),
	UpperBound: common.Ptr(int64(256)),
}

type EUTRAFrequencyBandIndicator struct {
	Value int64
}

func (ie *EUTRAFrequencyBandIndicator) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, eUTRAFrequencyBandIndicatorConstraints)
}

func (ie *EUTRAFrequencyBandIndicator) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(eUTRAFrequencyBandIndicatorConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
