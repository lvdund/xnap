package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var rANAreaIDListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofRANAreasinRNA)),
}

type RANAreaIDList struct {
	Value []*RANAreaID
}

func (ie *RANAreaIDList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(rANAreaIDListConstraints)
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

func (ie *RANAreaIDList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(rANAreaIDListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*RANAreaID, n)
	for i := range ie.Value {
		ie.Value[i] = new(RANAreaID)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
