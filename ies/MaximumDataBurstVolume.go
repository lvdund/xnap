package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var maximumDataBurstVolumeConstraints = per.IntegerConstraints{
	Extensible: true,
	LowerBound: common.Ptr(int64(0)),
	UpperBound: common.Ptr(int64(4095)),
}

type MaximumDataBurstVolume struct {
	Value int64
}

func (ie *MaximumDataBurstVolume) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, maximumDataBurstVolumeConstraints)
}

func (ie *MaximumDataBurstVolume) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(maximumDataBurstVolumeConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
