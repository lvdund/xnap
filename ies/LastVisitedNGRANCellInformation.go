package ies

import (
	"github.com/lvdund/asn1go/per"
)

var lastVisitedNGRANCellInformationConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        nil,
	Max:        nil,
}

type LastVisitedNGRANCellInformation struct {
	Value []byte
}

func (ie *LastVisitedNGRANCellInformation) Encode(e *per.Encoder) error {
	return e.EncodeOctetString(ie.Value, lastVisitedNGRANCellInformationConstraints)
}

func (ie *LastVisitedNGRANCellInformation) Decode(d *per.Decoder) error {
	val, err := d.DecodeOctetString(lastVisitedNGRANCellInformationConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
