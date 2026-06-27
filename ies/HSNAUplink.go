package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	HSNAUplinkHard         int64 = 0
	HSNAUplinkSoft         int64 = 1
	HSNAUplinkNotavailable int64 = 2
)

var hSNAUplinkConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0, 1, 2},
	ExtValues:  nil,
}

type HSNAUplink struct {
	Value int64
}

func (ie *HSNAUplink) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, hSNAUplinkConstraints)
}

func (ie *HSNAUplink) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(hSNAUplinkConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
