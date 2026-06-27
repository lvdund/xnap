package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var areaOfInterestInformationConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofAoIs)),
}

type AreaOfInterestInformation struct {
	Value []*AreaOfInterestItem
}

func (ie *AreaOfInterestInformation) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(areaOfInterestInformationConstraints)
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

func (ie *AreaOfInterestInformation) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(areaOfInterestInformationConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*AreaOfInterestItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(AreaOfInterestItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
