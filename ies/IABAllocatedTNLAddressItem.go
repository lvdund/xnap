package ies

import (
	"github.com/lvdund/asn1go/per"
)

var iABAllocatedTNLAddressItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "iABTNLAddress"},
		{Name: "iABTNLAddressUsage", Optional: true},
		{Name: "associatedDonorDUAddress", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type IABAllocatedTNLAddressItem struct {
	IABTNLAddress            IABTNLAddress
	IABTNLAddressUsage       *IABTNLAddressUsage
	AssociatedDonorDUAddress *BAPAddress
	IEExtensions             []byte
}

func (ie *IABAllocatedTNLAddressItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(iABAllocatedTNLAddressItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.IABTNLAddressUsage != nil, ie.AssociatedDonorDUAddress != nil, false}); err != nil {
		return err
	}
	if err := ie.IABTNLAddress.Encode(e); err != nil {
		return err
	}
	if ie.IABTNLAddressUsage != nil {
		if err := ie.IABTNLAddressUsage.Encode(e); err != nil {
			return err
		}
	}
	if ie.AssociatedDonorDUAddress != nil {
		if err := ie.AssociatedDonorDUAddress.Encode(e); err != nil {
			return err
		}
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *IABAllocatedTNLAddressItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(iABAllocatedTNLAddressItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.IABTNLAddress.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.IABTNLAddressUsage = new(IABTNLAddressUsage)
		if err := ie.IABTNLAddressUsage.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(2) {
		ie.AssociatedDonorDUAddress = new(BAPAddress)
		if err := ie.AssociatedDonorDUAddress.Decode(d); err != nil {
			return err
		}
	}
	extBytes, err := seq.DecodeExtensionAdditions()
	if err != nil {
		return err
	}
	if len(extBytes) > 0 {
		ie.IEExtensions = extBytes[0]
	}
	return nil
}
