package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	IABAuthorizationStatusAuthorized    int64 = 0
	IABAuthorizationStatusNotAuthorized int64 = 1
)

var iABAuthorizationStatusConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type IABAuthorizationStatus struct {
	Value int64
}

func (ie *IABAuthorizationStatus) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, iABAuthorizationStatusConstraints)
}

func (ie *IABAuthorizationStatus) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(iABAuthorizationStatusConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
