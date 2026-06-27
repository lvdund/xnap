package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var mBSAreaSessionIDConstraints = per.IntegerConstraints{
	Extensible: true,
	LowerBound: common.Ptr(int64(0)),
	UpperBound: common.Ptr(int64(65535)),
}

type MBSAreaSessionID struct {
	Value int64
}

func (ie *MBSAreaSessionID) Encode(e *per.Encoder) error {
	return e.EncodeInteger(ie.Value, mBSAreaSessionIDConstraints)
}

func (ie *MBSAreaSessionID) Decode(d *per.Decoder) error {
	val, err := d.DecodeInteger(mBSAreaSessionIDConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
