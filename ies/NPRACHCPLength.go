package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	NPRACHCPLengthUs66dot7  int64 = 0
	NPRACHCPLengthUs266dot7 int64 = 1
)

var nPRACHCPLengthConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type NPRACHCPLength struct {
	Value int64
}

func (ie *NPRACHCPLength) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, nPRACHCPLengthConstraints)
}

func (ie *NPRACHCPLength) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(nPRACHCPLengthConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
