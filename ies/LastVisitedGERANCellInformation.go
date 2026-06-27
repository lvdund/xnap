package ies

import (
	"github.com/lvdund/asn1go/per"
)

var lastVisitedGERANCellInformationConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        nil,
	Max:        nil,
}

type LastVisitedGERANCellInformation struct {
	Value []byte
}

func (ie *LastVisitedGERANCellInformation) Encode(e *per.Encoder) error {
	return e.EncodeOctetString(ie.Value, lastVisitedGERANCellInformationConstraints)
}

func (ie *LastVisitedGERANCellInformation) Decode(d *per.Decoder) error {
	val, err := d.DecodeOctetString(lastVisitedGERANCellInformationConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
