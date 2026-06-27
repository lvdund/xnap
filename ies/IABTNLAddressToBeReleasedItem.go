package ies

import (
	"github.com/lvdund/asn1go/per"
)

var iABTNLAddressToBeReleasedItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "iabTNLAddress"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type IABTNLAddressToBeReleasedItem struct {
	IabTNLAddress IABTNLAddress
	IEExtensions  []byte
}

func (ie *IABTNLAddressToBeReleasedItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(iABTNLAddressToBeReleasedItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.IabTNLAddress.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *IABTNLAddressToBeReleasedItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(iABTNLAddressToBeReleasedItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.IabTNLAddress.Decode(d); err != nil {
		return err
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
