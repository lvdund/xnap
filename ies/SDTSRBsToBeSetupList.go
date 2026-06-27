package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var sDTSRBsToBeSetupListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofSRBs)),
}

type SDTSRBsToBeSetupList struct {
	Value []*SDTSRBsToBeSetupListItem
}

func (ie *SDTSRBsToBeSetupList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(sDTSRBsToBeSetupListConstraints)
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

func (ie *SDTSRBsToBeSetupList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(sDTSRBsToBeSetupListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*SDTSRBsToBeSetupListItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(SDTSRBsToBeSetupListItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
