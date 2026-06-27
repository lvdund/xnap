package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var tNLAFailedToSetupListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofTNLAssociations)),
}

type TNLAFailedToSetupList struct {
	Value []*TNLAFailedToSetupItem
}

func (ie *TNLAFailedToSetupList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(tNLAFailedToSetupListConstraints)
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

func (ie *TNLAFailedToSetupList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(tNLAFailedToSetupListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*TNLAFailedToSetupItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(TNLAFailedToSetupItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
