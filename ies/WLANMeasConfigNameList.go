package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var wLANMeasConfigNameListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofWLANName)),
}

type WLANMeasConfigNameList struct {
	Value []*WLANName
}

func (ie *WLANMeasConfigNameList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(wLANMeasConfigNameListConstraints)
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

func (ie *WLANMeasConfigNameList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(wLANMeasConfigNameListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*WLANName, n)
	for i := range ie.Value {
		ie.Value[i] = new(WLANName)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
