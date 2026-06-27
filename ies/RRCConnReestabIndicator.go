package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	RRCConnReestabIndicatorReconfigurationFailure int64 = 0
	RRCConnReestabIndicatorHandoverFailure        int64 = 1
	RRCConnReestabIndicatorOtherFailure           int64 = 2
)

var rRCConnReestabIndicatorConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2},
	ExtValues:  nil,
}

type RRCConnReestabIndicator struct {
	Value int64
}

func (ie *RRCConnReestabIndicator) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, rRCConnReestabIndicatorConstraints)
}

func (ie *RRCConnReestabIndicator) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(rRCConnReestabIndicatorConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
