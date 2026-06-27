package ies

import (
	"github.com/lvdund/asn1go/per"
)

var reportConfigContainerConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        nil,
	Max:        nil,
}

type ReportConfigContainer struct {
	Value []byte
}

func (ie *ReportConfigContainer) Encode(e *per.Encoder) error {
	return e.EncodeOctetString(ie.Value, reportConfigContainerConstraints)
}

func (ie *ReportConfigContainer) Decode(d *per.Decoder) error {
	val, err := d.DecodeOctetString(reportConfigContainerConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
