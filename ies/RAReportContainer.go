package ies

import (
	"github.com/lvdund/asn1go/per"
)

var rAReportContainerConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        nil,
	Max:        nil,
}

type RAReportContainer struct {
	Value []byte
}

func (ie *RAReportContainer) Encode(e *per.Encoder) error {
	return e.EncodeOctetString(ie.Value, rAReportContainerConstraints)
}

func (ie *RAReportContainer) Decode(d *per.Decoder) error {
	val, err := d.DecodeOctetString(rAReportContainerConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
