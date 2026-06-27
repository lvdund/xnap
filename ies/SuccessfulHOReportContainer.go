package ies

import (
	"github.com/lvdund/asn1go/per"
)

var successfulHOReportContainerConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        nil,
	Max:        nil,
}

type SuccessfulHOReportContainer struct {
	Value []byte
}

func (ie *SuccessfulHOReportContainer) Encode(e *per.Encoder) error {
	return e.EncodeOctetString(ie.Value, successfulHOReportContainerConstraints)
}

func (ie *SuccessfulHOReportContainer) Decode(d *per.Decoder) error {
	val, err := d.DecodeOctetString(successfulHOReportContainerConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
