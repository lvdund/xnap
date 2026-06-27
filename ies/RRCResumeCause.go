package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	RRCResumeCauseRnaUpdate int64 = 0
)

var rRCResumeCauseConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type RRCResumeCause struct {
	Value int64
}

func (ie *RRCResumeCause) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, rRCResumeCauseConstraints)
}

func (ie *RRCResumeCause) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(rRCResumeCauseConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
