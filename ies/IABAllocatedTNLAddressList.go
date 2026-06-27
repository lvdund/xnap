package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var iABAllocatedTNLAddressListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofTLAsIAB)),
}

type IABAllocatedTNLAddressList struct {
	Value []*IABAllocatedTNLAddressItem
}

func (ie *IABAllocatedTNLAddressList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(iABAllocatedTNLAddressListConstraints)
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

func (ie *IABAllocatedTNLAddressList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(iABAllocatedTNLAddressListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*IABAllocatedTNLAddressItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(IABAllocatedTNLAddressItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
