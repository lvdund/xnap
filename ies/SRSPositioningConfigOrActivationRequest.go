package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	SRSPositioningConfigOrActivationRequestTrue int64 = 0
)

var sRSPositioningConfigOrActivationRequestConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type SRSPositioningConfigOrActivationRequest struct {
	Value int64
}

func (ie *SRSPositioningConfigOrActivationRequest) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, sRSPositioningConfigOrActivationRequestConstraints)
}

func (ie *SRSPositioningConfigOrActivationRequest) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(sRSPositioningConfigOrActivationRequestConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
