package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var sSBAreaRadioResourceStatusListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofSSBAreas)),
}

type SSBAreaRadioResourceStatusList struct {
	Value []*SSBAreaRadioResourceStatusListItem
}

func (ie *SSBAreaRadioResourceStatusList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(sSBAreaRadioResourceStatusListConstraints)
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

func (ie *SSBAreaRadioResourceStatusList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(sSBAreaRadioResourceStatusListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*SSBAreaRadioResourceStatusListItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(SSBAreaRadioResourceStatusListItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
