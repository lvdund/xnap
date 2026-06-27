package ies

import (
	"github.com/lvdund/asn1go/per"
)

var tNLConfigurationInfoConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "extendedUPTransportLayerAddressesToAdd", Optional: true},
		{Name: "extendedUPTransportLayerAddressesToRemove", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type TNLConfigurationInfo struct {
	ExtendedUPTransportLayerAddressesToAdd    *ExtTLAs
	ExtendedUPTransportLayerAddressesToRemove *ExtTLAs
	IEExtensions                              []byte
}

func (ie *TNLConfigurationInfo) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(tNLConfigurationInfoConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.ExtendedUPTransportLayerAddressesToAdd != nil, ie.ExtendedUPTransportLayerAddressesToRemove != nil, false}); err != nil {
		return err
	}
	if ie.ExtendedUPTransportLayerAddressesToAdd != nil {
		if err := ie.ExtendedUPTransportLayerAddressesToAdd.Encode(e); err != nil {
			return err
		}
	}
	if ie.ExtendedUPTransportLayerAddressesToRemove != nil {
		if err := ie.ExtendedUPTransportLayerAddressesToRemove.Encode(e); err != nil {
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

func (ie *TNLConfigurationInfo) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(tNLConfigurationInfoConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if seq.IsComponentPresent(0) {
		ie.ExtendedUPTransportLayerAddressesToAdd = new(ExtTLAs)
		if err := ie.ExtendedUPTransportLayerAddressesToAdd.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(1) {
		ie.ExtendedUPTransportLayerAddressesToRemove = new(ExtTLAs)
		if err := ie.ExtendedUPTransportLayerAddressesToRemove.Decode(d); err != nil {
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
