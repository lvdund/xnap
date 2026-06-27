package ies

import (
	"github.com/lvdund/asn1go/per"
)

var successfulPSCellChangeReportContainerConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        nil,
	Max:        nil,
}

type SuccessfulPSCellChangeReportContainer struct {
	Value []byte
}

func (ie *SuccessfulPSCellChangeReportContainer) Encode(e *per.Encoder) error {
	return e.EncodeOctetString(ie.Value, successfulPSCellChangeReportContainerConstraints)
}

func (ie *SuccessfulPSCellChangeReportContainer) Decode(d *per.Decoder) error {
	val, err := d.DecodeOctetString(successfulPSCellChangeReportContainerConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
