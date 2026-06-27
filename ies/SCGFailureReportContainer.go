package ies

import (
	"github.com/lvdund/asn1go/per"
)

var sCGFailureReportContainerConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        nil,
	Max:        nil,
}

type SCGFailureReportContainer struct {
	Value []byte
}

func (ie *SCGFailureReportContainer) Encode(e *per.Encoder) error {
	return e.EncodeOctetString(ie.Value, sCGFailureReportContainerConstraints)
}

func (ie *SCGFailureReportContainer) Decode(d *per.Decoder) error {
	val, err := d.DecodeOctetString(sCGFailureReportContainerConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
