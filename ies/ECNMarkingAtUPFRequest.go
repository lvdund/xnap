package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	ECNMarkingAtUPFRequestUl   int64 = 0
	ECNMarkingAtUPFRequestDl   int64 = 1
	ECNMarkingAtUPFRequestBoth int64 = 2
	ECNMarkingAtUPFRequestStop int64 = 3
)

var eCNMarkingAtUPFRequestConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2, 3},
	ExtValues:  nil,
}

type ECNMarkingAtUPFRequest struct {
	Value int64
}

func (ie *ECNMarkingAtUPFRequest) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, eCNMarkingAtUPFRequestConstraints)
}

func (ie *ECNMarkingAtUPFRequest) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(eCNMarkingAtUPFRequestConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
