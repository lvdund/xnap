package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	HSNADownlinkHard         int64 = 0
	HSNADownlinkSoft         int64 = 1
	HSNADownlinkNotavailable int64 = 2
)

var hSNADownlinkConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0, 1, 2},
	ExtValues:  nil,
}

type HSNADownlink struct {
	Value int64
}

func (ie *HSNADownlink) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, hSNADownlinkConstraints)
}

func (ie *HSNADownlink) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(hSNADownlinkConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
