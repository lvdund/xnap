package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var eRedcapBcastInformationConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(8)),
	Max:        common.Ptr(int64(8)),
}

type ERedcapBcastInformation struct {
	Value per.BitString
}

func (ie *ERedcapBcastInformation) Encode(e *per.Encoder) error {
	return e.EncodeBitString(ie.Value, eRedcapBcastInformationConstraints)
}

func (ie *ERedcapBcastInformation) Decode(d *per.Decoder) error {
	val, err := d.DecodeBitString(eRedcapBcastInformationConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
