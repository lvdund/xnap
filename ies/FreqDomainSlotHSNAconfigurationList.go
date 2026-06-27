package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var freqDomainSlotHSNAconfigurationListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofHSNASlots)),
}

type FreqDomainSlotHSNAconfigurationList struct {
	Value []*FreqDomainSlotHSNAconfigurationListItem
}

func (ie *FreqDomainSlotHSNAconfigurationList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(freqDomainSlotHSNAconfigurationListConstraints)
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

func (ie *FreqDomainSlotHSNAconfigurationList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(freqDomainSlotHSNAconfigurationListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*FreqDomainSlotHSNAconfigurationListItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(FreqDomainSlotHSNAconfigurationListItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
