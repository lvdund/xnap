package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	AerialUESubscriptionInformationAllowed    int64 = 0
	AerialUESubscriptionInformationNotAllowed int64 = 1
)

var aerialUESubscriptionInformationConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type AerialUESubscriptionInformation struct {
	Value int64
}

func (ie *AerialUESubscriptionInformation) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, aerialUESubscriptionInformationConstraints)
}

func (ie *AerialUESubscriptionInformation) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(aerialUESubscriptionInformationConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
