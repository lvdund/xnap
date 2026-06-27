package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var dRBsToBeSetupListSetupResponseSNterminatedConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofDRBs)),
}

type DRBsToBeSetupListSetupResponseSNterminated struct {
	Value []*DRBsToBeSetupListSetupResponseSNterminatedItem
}

func (ie *DRBsToBeSetupListSetupResponseSNterminated) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(dRBsToBeSetupListSetupResponseSNterminatedConstraints)
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

func (ie *DRBsToBeSetupListSetupResponseSNterminated) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(dRBsToBeSetupListSetupResponseSNterminatedConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*DRBsToBeSetupListSetupResponseSNterminatedItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(DRBsToBeSetupListSetupResponseSNterminatedItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
