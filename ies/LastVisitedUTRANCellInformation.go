package ies

import (
	"github.com/lvdund/asn1go/per"
)

var lastVisitedUTRANCellInformationConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        nil,
	Max:        nil,
}

type LastVisitedUTRANCellInformation struct {
	Value []byte
}

func (ie *LastVisitedUTRANCellInformation) Encode(e *per.Encoder) error {
	return e.EncodeOctetString(ie.Value, lastVisitedUTRANCellInformationConstraints)
}

func (ie *LastVisitedUTRANCellInformation) Decode(d *per.Decoder) error {
	val, err := d.DecodeOctetString(lastVisitedUTRANCellInformationConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
