package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var sNToMNQMCCoordRequestListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofUEAppLayerMeas)),
}

type SNToMNQMCCoordRequestList struct {
	Value []*SNToMNQMCCoordRequestListItem
}

func (ie *SNToMNQMCCoordRequestList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(sNToMNQMCCoordRequestListConstraints)
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

func (ie *SNToMNQMCCoordRequestList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(sNToMNQMCCoordRequestListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*SNToMNQMCCoordRequestListItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(SNToMNQMCCoordRequestListItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
