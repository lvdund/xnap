package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var aMFUENGAPIDConstraints = per.IntegerConstraints{
	Extensible: false,
	LowerBound: common.Ptr(int64(0)),
	UpperBound: common.Ptr(int64(1099511627775)),
}

type AMFUENGAPID struct {
	Value int64
}

func (ie *AMFUENGAPID) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, aMFUENGAPIDConstraints)
}

func (ie *AMFUENGAPID) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(aMFUENGAPIDConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
