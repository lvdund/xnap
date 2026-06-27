package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	FiveGProSeLayer2UEtoUERemoteAuthorized    int64 = 0
	FiveGProSeLayer2UEtoUERemoteNotAuthorized int64 = 1
)

var fiveGProSeLayer2UEtoUERemoteConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type FiveGProSeLayer2UEtoUERemote struct {
	Value int64
}

func (ie *FiveGProSeLayer2UEtoUERemote) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, fiveGProSeLayer2UEtoUERemoteConstraints)
}

func (ie *FiveGProSeLayer2UEtoUERemote) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(fiveGProSeLayer2UEtoUERemoteConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
