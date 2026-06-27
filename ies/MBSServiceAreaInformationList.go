package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var mBSServiceAreaInformationListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofMBSServiceAreaInformation)),
}

type MBSServiceAreaInformationList struct {
	Value []*MBSServiceAreaInformationItem
}

func (ie *MBSServiceAreaInformationList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(mBSServiceAreaInformationListConstraints)
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

func (ie *MBSServiceAreaInformationList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(mBSServiceAreaInformationListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*MBSServiceAreaInformationItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(MBSServiceAreaInformationItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
