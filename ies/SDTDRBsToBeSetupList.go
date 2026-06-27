package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var sDTDRBsToBeSetupListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofDRBs)),
}

type SDTDRBsToBeSetupList struct {
	Value []*SDTDRBsToBeSetupListItem
}

func (ie *SDTDRBsToBeSetupList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(sDTDRBsToBeSetupListConstraints)
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

func (ie *SDTDRBsToBeSetupList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(sDTDRBsToBeSetupListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*SDTDRBsToBeSetupListItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(SDTDRBsToBeSetupListItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
