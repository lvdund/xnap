package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	UserPlaneTrafficActivityReportInactive    int64 = 0
	UserPlaneTrafficActivityReportReActivated int64 = 1
)

var userPlaneTrafficActivityReportConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type UserPlaneTrafficActivityReport struct {
	Value int64
}

func (ie *UserPlaneTrafficActivityReport) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, userPlaneTrafficActivityReportConstraints)
}

func (ie *UserPlaneTrafficActivityReport) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(userPlaneTrafficActivityReportConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
