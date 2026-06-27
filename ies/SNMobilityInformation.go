package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var sNMobilityInformationConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(32)),
	Max:        common.Ptr(int64(32)),
}

type SNMobilityInformation struct {
	Value per.BitString
}

func (ie *SNMobilityInformation) Encode(e *per.Encoder) error {
	return e.EncodeBitString(ie.Value, sNMobilityInformationConstraints)
}

func (ie *SNMobilityInformation) Decode(d *per.Decoder) error {
	val, err := d.DecodeBitString(sNMobilityInformationConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
