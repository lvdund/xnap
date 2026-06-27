package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var extTLAsConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofExtTLAs)),
}

type ExtTLAs struct {
	Value []*ExtTLAItem
}

func (ie *ExtTLAs) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(extTLAsConstraints)
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

func (ie *ExtTLAs) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(extTLAsConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*ExtTLAItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(ExtTLAItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
