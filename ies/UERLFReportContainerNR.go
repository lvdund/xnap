package ies

import (
	"github.com/lvdund/asn1go/per"
)

var uERLFReportContainerNRConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        nil,
	Max:        nil,
}

type UERLFReportContainerNR struct {
	Value []byte
}

func (ie *UERLFReportContainerNR) Encode(e *per.Encoder) error {
	return e.EncodeOctetString(ie.Value, uERLFReportContainerNRConstraints)
}

func (ie *UERLFReportContainerNR) Decode(d *per.Decoder) error {
	val, err := d.DecodeOctetString(uERLFReportContainerNRConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
