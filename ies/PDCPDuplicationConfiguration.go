package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	PDCPDuplicationConfigurationConfigured   int64 = 0
	PDCPDuplicationConfigurationDeConfigured int64 = 1
)

var pDCPDuplicationConfigurationConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type PDCPDuplicationConfiguration struct {
	Value int64
}

func (ie *PDCPDuplicationConfiguration) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, pDCPDuplicationConfigurationConstraints)
}

func (ie *PDCPDuplicationConfiguration) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(pDCPDuplicationConfigurationConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
