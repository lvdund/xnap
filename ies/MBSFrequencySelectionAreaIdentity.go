package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var mBSFrequencySelectionAreaIdentityConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(3)),
	Max:        common.Ptr(int64(3)),
}

type MBSFrequencySelectionAreaIdentity struct {
	Value []byte
}

func (ie *MBSFrequencySelectionAreaIdentity) Encode(e *per.Encoder) error {
	return e.EncodeOctetString(ie.Value, mBSFrequencySelectionAreaIdentityConstraints)
}

func (ie *MBSFrequencySelectionAreaIdentity) Decode(d *per.Decoder) error {
	val, err := d.DecodeOctetString(mBSFrequencySelectionAreaIdentityConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
