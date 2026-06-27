package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var sNToMNQMCCoordResponseListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofUEAppLayerMeas)),
}

type SNToMNQMCCoordResponseList struct {
	Value []*SNToMNQMCCoordResponseListItem
}

func (ie *SNToMNQMCCoordResponseList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(sNToMNQMCCoordResponseListConstraints)
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

func (ie *SNToMNQMCCoordResponseList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(sNToMNQMCCoordResponseListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*SNToMNQMCCoordResponseListItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(SNToMNQMCCoordResponseListItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
