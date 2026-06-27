package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var containerAppLayerMeasConfigConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(8000)),
}

type ContainerAppLayerMeasConfig struct {
	Value []byte
}

func (ie *ContainerAppLayerMeasConfig) Encode(e *per.Encoder) error {
	return e.EncodeOctetString(ie.Value, containerAppLayerMeasConfigConstraints)
}

func (ie *ContainerAppLayerMeasConfig) Decode(d *per.Decoder) error {
	val, err := d.DecodeOctetString(containerAppLayerMeasConfigConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
