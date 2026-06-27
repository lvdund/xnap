package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var dLLBTFailureInformationListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofLBTFailureInformation)),
}

type DLLBTFailureInformationList struct {
	Value []*DLLBTFailureInformationListItem
}

func (ie *DLLBTFailureInformationList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(dLLBTFailureInformationListConstraints)
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

func (ie *DLLBTFailureInformationList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(dLLBTFailureInformationListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*DLLBTFailureInformationListItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(DLLBTFailureInformationListItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
