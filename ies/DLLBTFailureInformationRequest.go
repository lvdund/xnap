package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	DLLBTFailureInformationRequestInquiry int64 = 0
)

var dLLBTFailureInformationRequestConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type DLLBTFailureInformationRequest struct {
	Value int64
}

func (ie *DLLBTFailureInformationRequest) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, dLLBTFailureInformationRequestConstraints)
}

func (ie *DLLBTFailureInformationRequest) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(dLLBTFailureInformationRequestConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
