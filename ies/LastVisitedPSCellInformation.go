package ies

import (
	"github.com/lvdund/asn1go/per"
)

var lastVisitedPSCellInformationConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        nil,
	Max:        nil,
}

type LastVisitedPSCellInformation struct {
	Value []byte
}

func (ie *LastVisitedPSCellInformation) Encode(e *per.Encoder) error {
	return e.EncodeOctetString(ie.Value, lastVisitedPSCellInformationConstraints)
}

func (ie *LastVisitedPSCellInformation) Decode(d *per.Decoder) error {
	val, err := d.DecodeOctetString(lastVisitedPSCellInformationConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
