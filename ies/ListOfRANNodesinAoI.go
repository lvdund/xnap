package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var listOfRANNodesinAoIConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofRANNodesinAoI)),
}

type ListOfRANNodesinAoI struct {
	Value []*GlobalNGRANNodesinAoIItem
}

func (ie *ListOfRANNodesinAoI) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(listOfRANNodesinAoIConstraints)
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

func (ie *ListOfRANNodesinAoI) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(listOfRANNodesinAoIConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*GlobalNGRANNodesinAoIItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(GlobalNGRANNodesinAoIItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
