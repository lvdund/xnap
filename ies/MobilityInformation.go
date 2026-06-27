package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var mobilityInformationConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(32)),
	Max:        common.Ptr(int64(32)),
}

type MobilityInformation struct {
	Value per.BitString
}

func (ie *MobilityInformation) Encode(e *per.Encoder) error {
	return e.EncodeBitString(ie.Value, mobilityInformationConstraints)
}

func (ie *MobilityInformation) Decode(d *per.Decoder) error {
	val, err := d.DecodeBitString(mobilityInformationConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
