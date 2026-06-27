package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var nGRANnode2SSBOffsetsModificationRangeConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofSSBAreas)),
}

type NGRANnode2SSBOffsetsModificationRange struct {
	Value []*SSBOffsetModificationRange
}

func (ie *NGRANnode2SSBOffsetsModificationRange) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(nGRANnode2SSBOffsetsModificationRangeConstraints)
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

func (ie *NGRANnode2SSBOffsetsModificationRange) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(nGRANnode2SSBOffsetsModificationRangeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*SSBOffsetModificationRange, n)
	for i := range ie.Value {
		ie.Value[i] = new(SSBOffsetModificationRange)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
