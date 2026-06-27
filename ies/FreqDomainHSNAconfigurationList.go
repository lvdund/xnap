package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var freqDomainHSNAconfigurationListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofHSNASlots)),
}

type FreqDomainHSNAconfigurationList struct {
	Value []*FreqDomainHSNAconfigurationListItem
}

func (ie *FreqDomainHSNAconfigurationList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(freqDomainHSNAconfigurationListConstraints)
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

func (ie *FreqDomainHSNAconfigurationList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(freqDomainHSNAconfigurationListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*FreqDomainHSNAconfigurationListItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(FreqDomainHSNAconfigurationListItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
