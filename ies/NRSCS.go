package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	NRSCSScs15  int64 = 0
	NRSCSScs30  int64 = 1
	NRSCSScs60  int64 = 2
	NRSCSScs120 int64 = 3
	NRSCSScs480 int64 = 4
	NRSCSScs960 int64 = 5
)

var nRSCSConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2, 3},
	ExtValues:  []int64{4, 5},
}

type NRSCS struct {
	Value int64
}

func (ie *NRSCS) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, nRSCSConstraints)
}

func (ie *NRSCS) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(nRSCSConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
