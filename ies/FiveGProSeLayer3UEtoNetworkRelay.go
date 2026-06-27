package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	FiveGProSeLayer3UEtoNetworkRelayAuthorized    int64 = 0
	FiveGProSeLayer3UEtoNetworkRelayNotAuthorized int64 = 1
)

var fiveGProSeLayer3UEtoNetworkRelayConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type FiveGProSeLayer3UEtoNetworkRelay struct {
	Value int64
}

func (ie *FiveGProSeLayer3UEtoNetworkRelay) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, fiveGProSeLayer3UEtoNetworkRelayConstraints)
}

func (ie *FiveGProSeLayer3UEtoNetworkRelay) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(fiveGProSeLayer3UEtoNetworkRelayConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
