package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	MBSCommServiceTypeMulticast int64 = 0
	MBSCommServiceTypeBroadcast int64 = 1
)

var mBSCommServiceTypeConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type MBSCommServiceType struct {
	Value int64
}

func (ie *MBSCommServiceType) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, mBSCommServiceTypeConstraints)
}

func (ie *MBSCommServiceType) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(mBSCommServiceTypeConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
