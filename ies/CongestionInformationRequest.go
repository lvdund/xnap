package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	CongestionInformationRequestUl   int64 = 0
	CongestionInformationRequestDl   int64 = 1
	CongestionInformationRequestBoth int64 = 2
	CongestionInformationRequestStop int64 = 3
)

var congestionInformationRequestConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2, 3},
	ExtValues:  nil,
}

type CongestionInformationRequest struct {
	Value int64
}

func (ie *CongestionInformationRequest) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, congestionInformationRequestConstraints)
}

func (ie *CongestionInformationRequest) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(congestionInformationRequestConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
