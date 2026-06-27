package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	MobileIABAuthorizationStatusAuthorized    int64 = 0
	MobileIABAuthorizationStatusNotAuthorized int64 = 1
)

var mobileIABAuthorizationStatusConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type MobileIABAuthorizationStatus struct {
	Value int64
}

func (ie *MobileIABAuthorizationStatus) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, mobileIABAuthorizationStatusConstraints)
}

func (ie *MobileIABAuthorizationStatus) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(mobileIABAuthorizationStatusConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
