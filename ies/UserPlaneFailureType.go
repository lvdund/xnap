package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	UserPlaneFailureTypeGtpUErrorIndicationReceived int64 = 0
	UserPlaneFailureTypeUpPathFailure               int64 = 1
)

var userPlaneFailureTypeConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type UserPlaneFailureType struct {
	Value int64
}

func (ie *UserPlaneFailureType) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, userPlaneFailureTypeConstraints)
}

func (ie *UserPlaneFailureType) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(userPlaneFailureTypeConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
