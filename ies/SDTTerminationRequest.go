package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	SDTTerminationRequestRadioLinkProblem      int64 = 0
	SDTTerminationRequestNormal                int64 = 1
	SDTTerminationRequestLargeSdtVolumeFromBSR int64 = 2
)

var sDTTerminationRequestConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1},
	ExtValues:  []int64{2},
}

type SDTTerminationRequest struct {
	Value int64
}

func (ie *SDTTerminationRequest) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, sDTTerminationRequestConstraints)
}

func (ie *SDTTerminationRequest) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(sDTTerminationRequestConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
