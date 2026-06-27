package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	LocationInformationSNReportingPSCell int64 = 0
)

var locationInformationSNReportingConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type LocationInformationSNReporting struct {
	Value int64
}

func (ie *LocationInformationSNReporting) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, locationInformationSNReportingConstraints)
}

func (ie *LocationInformationSNReporting) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(locationInformationSNReportingConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
