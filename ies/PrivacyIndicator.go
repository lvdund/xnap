package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	PrivacyIndicatorImmediateMDT int64 = 0
	PrivacyIndicatorLoggedMDT    int64 = 1
)

var privacyIndicatorConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type PrivacyIndicator struct {
	Value int64
}

func (ie *PrivacyIndicator) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, privacyIndicatorConstraints)
}

func (ie *PrivacyIndicator) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(privacyIndicatorConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
