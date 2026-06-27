package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var cSIRSNeighbourListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofCSIRSneighbourCells)),
}

type CSIRSNeighbourList struct {
	Value []*CSIRSNeighbourItem
}

func (ie *CSIRSNeighbourList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(cSIRSNeighbourListConstraints)
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

func (ie *CSIRSNeighbourList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(cSIRSNeighbourListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*CSIRSNeighbourItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(CSIRSNeighbourItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
