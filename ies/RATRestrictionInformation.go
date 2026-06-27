package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var rATRestrictionInformationConstraints = per.SizeConstraints{
	Extensible: true,
	Min:        common.Ptr(int64(8)),
	Max:        common.Ptr(int64(8)),
}

type RATRestrictionInformation struct {
	Value per.BitString
}

func (ie *RATRestrictionInformation) Encode(e *per.Encoder) error {
	return e.EncodeBitString(ie.Value, rATRestrictionInformationConstraints)
}

func (ie *RATRestrictionInformation) Decode(d *per.Decoder) error {
	val, err := d.DecodeBitString(rATRestrictionInformationConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
