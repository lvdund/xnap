package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var transportLayerAddressConstraints = per.SizeConstraints{
	Extensible: true,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(160)),
}

type TransportLayerAddress struct {
	Value per.BitString
}

func (ie *TransportLayerAddress) Encode(e *per.Encoder) error {
	return e.EncodeBitString(ie.Value, transportLayerAddressConstraints)
}

func (ie *TransportLayerAddress) Decode(d *per.Decoder) error {
	val, err := d.DecodeBitString(transportLayerAddressConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
