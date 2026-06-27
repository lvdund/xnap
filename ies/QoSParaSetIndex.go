package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var qoSParaSetIndexConstraints = per.IntegerConstraints{
	Extensible: true,
	LowerBound: common.Ptr(int64(1)),
	UpperBound: common.Ptr(int64(8)),
}

type QoSParaSetIndex struct {
	Value int64
}

func (ie *QoSParaSetIndex) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, qoSParaSetIndexConstraints)
}

func (ie *QoSParaSetIndex) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(qoSParaSetIndexConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
