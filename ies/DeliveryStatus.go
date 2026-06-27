package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var deliveryStatusConstraints = per.IntegerConstraints{
	Extensible: true,
	LowerBound: common.Ptr(int64(0)),
	UpperBound: common.Ptr(int64(4095)),
}

type DeliveryStatus struct {
	Value int64
}

func (ie *DeliveryStatus) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, deliveryStatusConstraints)
}

func (ie *DeliveryStatus) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(deliveryStatusConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
