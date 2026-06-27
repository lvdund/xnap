package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	ECNMarkingAtRANRequestUl   int64 = 0
	ECNMarkingAtRANRequestDl   int64 = 1
	ECNMarkingAtRANRequestBoth int64 = 2
	ECNMarkingAtRANRequestStop int64 = 3
)

var eCNMarkingAtRANRequestConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2, 3},
	ExtValues:  nil,
}

type ECNMarkingAtRANRequest struct {
	Value int64
}

func (ie *ECNMarkingAtRANRequest) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, eCNMarkingAtRANRequestConstraints)
}

func (ie *ECNMarkingAtRANRequest) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(eCNMarkingAtRANRequestConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
