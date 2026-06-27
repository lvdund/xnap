package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var neighbourNGRANNodeListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(0)),
	Max:        common.Ptr(int64(common.MaxnoofNeighbour_NG_RAN_Nodes)),
}

type NeighbourNGRANNodeList struct {
	Value []*NeighbourNGRANNodeItem
}

func (ie *NeighbourNGRANNodeList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(neighbourNGRANNodeListConstraints)
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

func (ie *NeighbourNGRANNodeList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(neighbourNGRANNodeListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*NeighbourNGRANNodeItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(NeighbourNGRANNodeItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
