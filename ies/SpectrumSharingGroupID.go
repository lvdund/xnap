package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var spectrumSharingGroupIDConstraints = per.IntegerConstraints{
	Extensible: false,
	LowerBound: common.Ptr(int64(1)),
	UpperBound: common.Ptr(int64(common.MaxnoofCellsinNG_RANnode)),
}

type SpectrumSharingGroupID struct {
	Value int64
}

func (ie *SpectrumSharingGroupID) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, spectrumSharingGroupIDConstraints)
}

func (ie *SpectrumSharingGroupID) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(spectrumSharingGroupIDConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
