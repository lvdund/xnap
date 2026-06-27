package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var mBSSessionInformationResponseListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofMBSSessions)),
}

type MBSSessionInformationResponseList struct {
	Value []*MBSSessionInformationResponseItem
}

func (ie *MBSSessionInformationResponseList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(mBSSessionInformationResponseListConstraints)
	if err := seqOf.EncodeLength(int64(len(ie.Value))); err != nil {
		return err
	}
	for _, item := range ie.Value {
		if err := item.Encode(e); err != nil {
			return err
		}
	}
	return nil
}

func (ie *MBSSessionInformationResponseList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(mBSSessionInformationResponseListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*MBSSessionInformationResponseItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(MBSSessionInformationResponseItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
