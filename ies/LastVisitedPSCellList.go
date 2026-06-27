package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var lastVisitedPSCellListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofPSCellsPerSN)),
}

type LastVisitedPSCellList struct {
	Value []*LastVisitedPSCellListItem
}

func (ie *LastVisitedPSCellList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(lastVisitedPSCellListConstraints)
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

func (ie *LastVisitedPSCellList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(lastVisitedPSCellListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*LastVisitedPSCellListItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(LastVisitedPSCellListItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
