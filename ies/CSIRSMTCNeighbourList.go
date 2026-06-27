package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var cSIRSMTCNeighbourListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofCSIRSneighbourCellsInMTC)),
}

type CSIRSMTCNeighbourList struct {
	Value []*CSIRSMTCNeighbourItem
}

func (ie *CSIRSMTCNeighbourList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(cSIRSMTCNeighbourListConstraints)
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

func (ie *CSIRSMTCNeighbourList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(cSIRSMTCNeighbourListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*CSIRSMTCNeighbourItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(CSIRSMTCNeighbourItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
