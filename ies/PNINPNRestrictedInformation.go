package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	PNINPNRestrictedInformationRestriced     int64 = 0
	PNINPNRestrictedInformationNotRestricted int64 = 1
)

var pNINPNRestrictedInformationConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type PNINPNRestrictedInformation struct {
	Value int64
}

func (ie *PNINPNRestrictedInformation) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, pNINPNRestrictedInformationConstraints)
}

func (ie *PNINPNRestrictedInformation) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(pNINPNRestrictedInformationConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
