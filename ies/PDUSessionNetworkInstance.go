package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var pDUSessionNetworkInstanceConstraints = per.IntegerConstraints{
	Extensible: true,
	LowerBound: common.Ptr(int64(1)),
	UpperBound: common.Ptr(int64(256)),
}

type PDUSessionNetworkInstance struct {
	Value int64
}

func (ie *PDUSessionNetworkInstance) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, pDUSessionNetworkInstanceConstraints)
}

func (ie *PDUSessionNetworkInstance) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(pDUSessionNetworkInstanceConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
