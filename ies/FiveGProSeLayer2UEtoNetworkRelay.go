package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	FiveGProSeLayer2UEtoNetworkRelayAuthorized    int64 = 0
	FiveGProSeLayer2UEtoNetworkRelayNotAuthorized int64 = 1
)

var fiveGProSeLayer2UEtoNetworkRelayConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type FiveGProSeLayer2UEtoNetworkRelay struct {
	Value int64
}

func (ie *FiveGProSeLayer2UEtoNetworkRelay) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, fiveGProSeLayer2UEtoNetworkRelayConstraints)
}

func (ie *FiveGProSeLayer2UEtoNetworkRelay) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(fiveGProSeLayer2UEtoNetworkRelayConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
