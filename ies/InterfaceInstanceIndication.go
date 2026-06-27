package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var interfaceInstanceIndicationConstraints = per.IntegerConstraints{
	Extensible: true,
	LowerBound: common.Ptr(int64(0)),
	UpperBound: common.Ptr(int64(255)),
}

type InterfaceInstanceIndication struct {
	Value int64
}

func (ie *InterfaceInstanceIndication) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, interfaceInstanceIndicationConstraints)
}

func (ie *InterfaceInstanceIndication) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(interfaceInstanceIndicationConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
