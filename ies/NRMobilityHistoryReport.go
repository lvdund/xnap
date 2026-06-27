package ies

import (
	"github.com/lvdund/asn1go/per"
)

var nRMobilityHistoryReportConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        nil,
	Max:        nil,
}

type NRMobilityHistoryReport struct {
	Value []byte
}

func (ie *NRMobilityHistoryReport) Encode(e *per.Encoder) error {
	return e.EncodeOctetString(ie.Value, nRMobilityHistoryReportConstraints)
}

func (ie *NRMobilityHistoryReport) Decode(d *per.Decoder) error {
	val, err := d.DecodeOctetString(nRMobilityHistoryReportConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
