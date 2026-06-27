package ies

import (
	"github.com/lvdund/asn1go/per"
)

var lastVisitedEUTRANCellInformationConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        nil,
	Max:        nil,
}

type LastVisitedEUTRANCellInformation struct {
	Value []byte
}

func (ie *LastVisitedEUTRANCellInformation) Encode(e *per.Encoder) error {
	return e.EncodeOctetString(ie.Value, lastVisitedEUTRANCellInformationConstraints)
}

func (ie *LastVisitedEUTRANCellInformation) Decode(d *per.Decoder) error {
	val, err := d.DecodeOctetString(lastVisitedEUTRANCellInformationConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
