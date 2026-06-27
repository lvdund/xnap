package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	RegistrationRequestStart int64 = 0
	RegistrationRequestStop  int64 = 1
	RegistrationRequestAdd   int64 = 2
)

var registrationRequestConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2},
	ExtValues:  nil,
}

type RegistrationRequest struct {
	Value int64
}

func (ie *RegistrationRequest) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, registrationRequestConstraints)
}

func (ie *RegistrationRequest) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(registrationRequestConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
