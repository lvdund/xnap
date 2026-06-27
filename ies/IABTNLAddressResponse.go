package ies

import (
	"github.com/lvdund/asn1go/per"
)

var iABTNLAddressResponseConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "iABAllocatedTNLAddress-List"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type IABTNLAddressResponse struct {
	IABAllocatedTNLAddressList IABAllocatedTNLAddressList
	IEExtensions               []byte
}

func (ie *IABTNLAddressResponse) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(iABTNLAddressResponseConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.IABAllocatedTNLAddressList.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *IABTNLAddressResponse) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(iABTNLAddressResponseConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.IABAllocatedTNLAddressList.Decode(d); err != nil {
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
