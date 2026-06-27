package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var gTPTLAsConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofGTPTLAs)),
}

type GTPTLAs struct {
	Value []*GTPTLAItem
}

func (ie *GTPTLAs) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(gTPTLAsConstraints)
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

func (ie *GTPTLAs) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(gTPTLAsConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*GTPTLAItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(GTPTLAItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
