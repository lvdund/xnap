package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var mBSSessionInformationListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofMBSSessions)),
}

type MBSSessionInformationList struct {
	Value []*MBSSessionInformationItem
}

func (ie *MBSSessionInformationList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(mBSSessionInformationListConstraints)
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

func (ie *MBSSessionInformationList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(mBSSessionInformationListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*MBSSessionInformationItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(MBSSessionInformationItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
