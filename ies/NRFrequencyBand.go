package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var nRFrequencyBandConstraints = per.IntegerConstraints{
	Extensible: true,
	LowerBound: common.Ptr(int64(1)),
	UpperBound: common.Ptr(int64(1024)),
}

type NRFrequencyBand struct {
	Value int64
}

func (ie *NRFrequencyBand) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, nRFrequencyBandConstraints)
}

func (ie *NRFrequencyBand) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(nRFrequencyBandConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
