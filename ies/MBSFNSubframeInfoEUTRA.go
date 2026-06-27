package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var mBSFNSubframeInfoEUTRAConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofMBSFNEUTRA)),
}

type MBSFNSubframeInfoEUTRA struct {
	Value []*MBSFNSubframeInfoEUTRAItem
}

func (ie *MBSFNSubframeInfoEUTRA) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(mBSFNSubframeInfoEUTRAConstraints)
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

func (ie *MBSFNSubframeInfoEUTRA) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(mBSFNSubframeInfoEUTRAConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*MBSFNSubframeInfoEUTRAItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(MBSFNSubframeInfoEUTRAItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
