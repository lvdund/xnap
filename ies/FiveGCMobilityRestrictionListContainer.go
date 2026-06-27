package ies

import (
	"github.com/lvdund/asn1go/per"
)

var fiveGCMobilityRestrictionListContainerConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        nil,
	Max:        nil,
}

type FiveGCMobilityRestrictionListContainer struct {
	Value []byte
}

func (ie *FiveGCMobilityRestrictionListContainer) Encode(e *per.Encoder) error {
	return e.EncodeOctetString(ie.Value, fiveGCMobilityRestrictionListContainerConstraints)
}

func (ie *FiveGCMobilityRestrictionListContainer) Decode(d *per.Decoder) error {
	val, err := d.DecodeOctetString(fiveGCMobilityRestrictionListContainerConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
