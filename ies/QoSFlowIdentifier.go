package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var qoSFlowIdentifierConstraints = per.IntegerConstraints{
	Extensible: true,
	LowerBound: common.Ptr(int64(0)),
	UpperBound: common.Ptr(int64(63)),
}

type QoSFlowIdentifier struct {
	Value int64
}

func (ie *QoSFlowIdentifier) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, qoSFlowIdentifierConstraints)
}

func (ie *QoSFlowIdentifier) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(qoSFlowIdentifierConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
