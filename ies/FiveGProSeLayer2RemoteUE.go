package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	FiveGProSeLayer2RemoteUEAuthorized    int64 = 0
	FiveGProSeLayer2RemoteUENotAuthorized int64 = 1
)

var fiveGProSeLayer2RemoteUEConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type FiveGProSeLayer2RemoteUE struct {
	Value int64
}

func (ie *FiveGProSeLayer2RemoteUE) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, fiveGProSeLayer2RemoteUEConstraints)
}

func (ie *FiveGProSeLayer2RemoteUE) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(fiveGProSeLayer2RemoteUEConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
