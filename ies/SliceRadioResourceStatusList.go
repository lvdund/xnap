package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var sliceRadioResourceStatusListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofBPLMNs)),
}

type SliceRadioResourceStatusList struct {
	Value []*SliceRadioResourceStatusItem
}

func (ie *SliceRadioResourceStatusList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(sliceRadioResourceStatusListConstraints)
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

func (ie *SliceRadioResourceStatusList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(sliceRadioResourceStatusListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*SliceRadioResourceStatusItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(SliceRadioResourceStatusItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
