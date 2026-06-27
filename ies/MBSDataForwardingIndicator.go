package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	MBSDataForwardingIndicatorMbsOnly int64 = 0
)

var mBSDataForwardingIndicatorConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type MBSDataForwardingIndicator struct {
	Value int64
}

func (ie *MBSDataForwardingIndicator) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, mBSDataForwardingIndicatorConstraints)
}

func (ie *MBSDataForwardingIndicator) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(mBSDataForwardingIndicatorConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
