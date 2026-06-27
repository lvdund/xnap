package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var cPCInformationUpdatePSCellsListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofPSCellCandidates)),
}

type CPCInformationUpdatePSCellsList struct {
	Value []*CPCInformationUpdatePSCellsItem
}

func (ie *CPCInformationUpdatePSCellsList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(cPCInformationUpdatePSCellsListConstraints)
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

func (ie *CPCInformationUpdatePSCellsList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(cPCInformationUpdatePSCellsListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*CPCInformationUpdatePSCellsItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(CPCInformationUpdatePSCellsItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
