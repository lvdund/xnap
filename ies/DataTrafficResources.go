package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var dataTrafficResourcesConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(6)),
	Max:        common.Ptr(int64(17600)),
}

type DataTrafficResources struct {
	Value per.BitString
}

func (ie *DataTrafficResources) Encode(e *per.Encoder) error {
	return e.EncodeBitString(ie.Value, dataTrafficResourcesConstraints)
}

func (ie *DataTrafficResources) Decode(d *per.Decoder) error {
	val, err := d.DecodeBitString(dataTrafficResourcesConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
