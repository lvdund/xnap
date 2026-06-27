package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var dUFSlotConfigListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofDUFSlots)),
}

type DUFSlotConfigList struct {
	Value []*DUFSlotConfigItem
}

func (ie *DUFSlotConfigList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(dUFSlotConfigListConstraints)
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

func (ie *DUFSlotConfigList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(dUFSlotConfigListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*DUFSlotConfigItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(DUFSlotConfigItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
