package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var qoSParaSetNotifyIndexConstraints = per.IntegerConstraints{
	Extensible: true,
	LowerBound: common.Ptr(int64(0)),
	UpperBound: common.Ptr(int64(8)),
}

type QoSParaSetNotifyIndex struct {
	Value int64
}

func (ie *QoSParaSetNotifyIndex) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, qoSParaSetNotifyIndexConstraints)
}

func (ie *QoSParaSetNotifyIndex) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(qoSParaSetNotifyIndexConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
