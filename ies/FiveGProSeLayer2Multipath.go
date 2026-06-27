package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	FiveGProSeLayer2MultipathAuthorized    int64 = 0
	FiveGProSeLayer2MultipathNotAuthorized int64 = 1
)

var fiveGProSeLayer2MultipathConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type FiveGProSeLayer2Multipath struct {
	Value int64
}

func (ie *FiveGProSeLayer2Multipath) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, fiveGProSeLayer2MultipathConstraints)
}

func (ie *FiveGProSeLayer2Multipath) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(fiveGProSeLayer2MultipathConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
