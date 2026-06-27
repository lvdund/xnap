package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var sULFrequencyBandConstraints = per.IntegerConstraints{
	Extensible: false,
	LowerBound: common.Ptr(int64(1)),
	UpperBound: common.Ptr(int64(1024)),
}

type SULFrequencyBand struct {
	Value int64
}

func (ie *SULFrequencyBand) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, sULFrequencyBandConstraints)
}

func (ie *SULFrequencyBand) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(sULFrequencyBandConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
