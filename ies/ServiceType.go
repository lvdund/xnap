package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	ServiceTypeQMCForStreamingService int64 = 0
	ServiceTypeQMCForMTSIService      int64 = 1
	ServiceTypeQMCForVRService        int64 = 2
)

var serviceTypeConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2},
	ExtValues:  nil,
}

type ServiceType struct {
	Value int64
}

func (ie *ServiceType) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, serviceTypeConstraints)
}

func (ie *ServiceType) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(serviceTypeConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
