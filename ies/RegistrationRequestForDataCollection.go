package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	RegistrationRequestForDataCollectionStart int64 = 0
	RegistrationRequestForDataCollectionStop  int64 = 1
)

var registrationRequestForDataCollectionConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type RegistrationRequestForDataCollection struct {
	Value int64
}

func (ie *RegistrationRequestForDataCollection) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, registrationRequestForDataCollectionConstraints)
}

func (ie *RegistrationRequestForDataCollection) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(registrationRequestForDataCollectionConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
