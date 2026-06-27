package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var hSNASlotConfigListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofHSNASlots)),
}

type HSNASlotConfigList struct {
	Value []*HSNASlotConfigItem
}

func (ie *HSNASlotConfigList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(hSNASlotConfigListConstraints)
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

func (ie *HSNASlotConfigList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(hSNASlotConfigListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*HSNASlotConfigItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(HSNASlotConfigItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
