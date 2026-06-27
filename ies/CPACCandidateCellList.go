package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var cPACCandidateCellListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofPSCellsinCPAC)),
}

type CPACCandidateCellList struct {
	Value []*CPACCandidateCellItem
}

func (ie *CPACCandidateCellList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(cPACCandidateCellListConstraints)
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

func (ie *CPACCandidateCellList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(cPACCandidateCellListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*CPACCandidateCellItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(CPACCandidateCellItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
