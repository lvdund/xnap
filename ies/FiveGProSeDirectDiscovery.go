package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	FiveGProSeDirectDiscoveryAuthorized    int64 = 0
	FiveGProSeDirectDiscoveryNotAuthorized int64 = 1
)

var fiveGProSeDirectDiscoveryConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type FiveGProSeDirectDiscovery struct {
	Value int64
}

func (ie *FiveGProSeDirectDiscovery) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, fiveGProSeDirectDiscoveryConstraints)
}

func (ie *FiveGProSeDirectDiscovery) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(fiveGProSeDirectDiscoveryConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
