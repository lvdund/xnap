package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var redcapBcastInformationConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(8)),
	Max:        common.Ptr(int64(8)),
}

type RedcapBcastInformation struct {
	Value per.BitString
}

func (ie *RedcapBcastInformation) Encode(e *per.Encoder) error {
	return e.EncodeBitString(ie.Value, redcapBcastInformationConstraints)
}

func (ie *RedcapBcastInformation) Decode(d *per.Decoder) error {
	val, err := d.DecodeBitString(redcapBcastInformationConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
