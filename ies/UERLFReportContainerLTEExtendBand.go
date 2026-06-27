package ies

import (
	"github.com/lvdund/asn1go/per"
)

var uERLFReportContainerLTEExtendBandConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        nil,
	Max:        nil,
}

type UERLFReportContainerLTEExtendBand struct {
	Value []byte
}

func (ie *UERLFReportContainerLTEExtendBand) Encode(e *per.Encoder) error {
	return e.EncodeOctetString(ie.Value, uERLFReportContainerLTEExtendBandConstraints)
}

func (ie *UERLFReportContainerLTEExtendBand) Decode(d *per.Decoder) error {
	val, err := d.DecodeOctetString(uERLFReportContainerLTEExtendBandConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
