package ies

import (
	"github.com/lvdund/asn1go/per"
)

var iABTNLAddressRequestConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "iABIPv4AddressesRequested"},
		{Name: "iABIPv6RequestType"},
		{Name: "iABTNLAddressToRemove-List"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type IABTNLAddressRequest struct {
	IABIPv4AddressesRequested IABTNLAddressesRequested
	IABIPv6RequestType        IABIPv6RequestType
	IABTNLAddressToRemoveList IABTNLAddressToRemoveList
	IEExtensions              []byte
}

func (ie *IABTNLAddressRequest) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(iABTNLAddressRequestConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.IABIPv4AddressesRequested.Encode(e); err != nil {
		return err
	}
	if err := ie.IABIPv6RequestType.Encode(e); err != nil {
		return err
	}
	if err := ie.IABTNLAddressToRemoveList.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *IABTNLAddressRequest) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(iABTNLAddressRequestConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.IABIPv4AddressesRequested.Decode(d); err != nil {
		return err
	}
	if err := ie.IABIPv6RequestType.Decode(d); err != nil {
		return err
	}
	if err := ie.IABTNLAddressToRemoveList.Decode(d); err != nil {
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
