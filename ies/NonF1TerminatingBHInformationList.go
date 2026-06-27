package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var nonF1TerminatingBHInformationListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofBHInfo)),
}

type NonF1TerminatingBHInformationList struct {
	Value []*NonF1TerminatingBHInformationItem
}

func (ie *NonF1TerminatingBHInformationList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(nonF1TerminatingBHInformationListConstraints)
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

func (ie *NonF1TerminatingBHInformationList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(nonF1TerminatingBHInformationListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*NonF1TerminatingBHInformationItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(NonF1TerminatingBHInformationItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
