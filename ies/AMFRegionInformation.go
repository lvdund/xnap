package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var aMFRegionInformationConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofAMFRegions)),
}

type AMFRegionInformation struct {
	Value []*GlobalAMFRegionInformation
}

func (ie *AMFRegionInformation) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(aMFRegionInformationConstraints)
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

func (ie *AMFRegionInformation) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(aMFRegionInformationConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*GlobalAMFRegionInformation, n)
	for i := range ie.Value {
		ie.Value[i] = new(GlobalAMFRegionInformation)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
