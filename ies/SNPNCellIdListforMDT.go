package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var sNPNCellIdListforMDTConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofCellIDforMDT)),
}

type SNPNCellIdListforMDT struct {
	Value []*SNPNCellIdforMDTItem
}

func (ie *SNPNCellIdListforMDT) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(sNPNCellIdListforMDTConstraints)
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

func (ie *SNPNCellIdListforMDT) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(sNPNCellIdListforMDTConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*SNPNCellIdforMDTItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(SNPNCellIdforMDTItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
