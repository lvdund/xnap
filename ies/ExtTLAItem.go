package ies

import (
	"github.com/lvdund/asn1go/per"
)

var extTLAItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "iPsecTLA", Optional: true},
		{Name: "gTPTransportLayerAddresses", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type ExtTLAItem struct {
	IPsecTLA                   *TransportLayerAddress
	GTPTransportLayerAddresses *GTPTLAs
	IEExtensions               []byte
}

func (ie *ExtTLAItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(extTLAItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.IPsecTLA != nil, ie.GTPTransportLayerAddresses != nil, false}); err != nil {
		return err
	}
	if ie.IPsecTLA != nil {
		if err := ie.IPsecTLA.Encode(e); err != nil {
			return err
		}
	}
	if ie.GTPTransportLayerAddresses != nil {
		if err := ie.GTPTransportLayerAddresses.Encode(e); err != nil {
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

func (ie *ExtTLAItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(extTLAItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if seq.IsComponentPresent(0) {
		ie.IPsecTLA = new(TransportLayerAddress)
		if err := ie.IPsecTLA.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(1) {
		ie.GTPTransportLayerAddresses = new(GTPTLAs)
		if err := ie.GTPTransportLayerAddresses.Decode(d); err != nil {
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
