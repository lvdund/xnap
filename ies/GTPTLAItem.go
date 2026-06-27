package ies

import (
	"github.com/lvdund/asn1go/per"
)

var gTPTLAItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "gTPTransportLayerAddresses"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type GTPTLAItem struct {
	GTPTransportLayerAddresses TransportLayerAddress
	IEExtensions               []byte
}

func (ie *GTPTLAItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(gTPTLAItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.GTPTransportLayerAddresses.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *GTPTLAItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(gTPTLAItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.GTPTransportLayerAddresses.Decode(d); err != nil {
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
