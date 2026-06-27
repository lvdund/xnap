package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var nRFrequencyBandListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofNRCellBands)),
}

type NRFrequencyBandList struct {
	Value []*NRFrequencyBandItem
}

func (ie *NRFrequencyBandList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(nRFrequencyBandListConstraints)
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

func (ie *NRFrequencyBandList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(nRFrequencyBandListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*NRFrequencyBandItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(NRFrequencyBandItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
