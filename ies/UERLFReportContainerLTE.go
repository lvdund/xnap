package ies

import (
	"github.com/lvdund/asn1go/per"
)

var uERLFReportContainerLTEConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        nil,
	Max:        nil,
}

type UERLFReportContainerLTE struct {
	Value []byte
}

func (ie *UERLFReportContainerLTE) Encode(e *per.Encoder) error {
	return e.EncodeOctetString(ie.Value, uERLFReportContainerLTEConstraints)
}

func (ie *UERLFReportContainerLTE) Decode(d *per.Decoder) error {
	val, err := d.DecodeOctetString(uERLFReportContainerLTEConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
