package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var nGRANCellIdentityListinRANPagingAreaConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofCellsinRNA)),
}

type NGRANCellIdentityListinRANPagingArea struct {
	Value []*NGRANCellIdentity
}

func (ie *NGRANCellIdentityListinRANPagingArea) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(nGRANCellIdentityListinRANPagingAreaConstraints)
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

func (ie *NGRANCellIdentityListinRANPagingArea) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(nGRANCellIdentityListinRANPagingAreaConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*NGRANCellIdentity, n)
	for i := range ie.Value {
		ie.Value[i] = new(NGRANCellIdentity)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
