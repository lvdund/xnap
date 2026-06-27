package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var noofRRCConnectionsConstraints = per.IntegerConstraints{
	Extensible: true,
	LowerBound: common.Ptr(int64(1)),
	UpperBound: common.Ptr(int64(65536)),
}

type NoofRRCConnections struct {
	Value int64
}

func (ie *NoofRRCConnections) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, noofRRCConnectionsConstraints)
}

func (ie *NoofRRCConnections) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(noofRRCConnectionsConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
