package ies

import (
	"github.com/lvdund/asn1go/per"
)

var nonAnchorCarrierFrequencylistItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "non-anchorCarrierFrquency"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type NonAnchorCarrierFrequencylistItem struct {
	NonAnchorCarrierFrquency []byte
	IEExtensions             []byte
}

func (ie *NonAnchorCarrierFrequencylistItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(nonAnchorCarrierFrequencylistItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := e.EncodeOctetString(ie.NonAnchorCarrierFrquency, per.SizeConstraints{
		Extensible: false,
		Min:        nil,
		Max:        nil,
	}); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *NonAnchorCarrierFrequencylistItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(nonAnchorCarrierFrequencylistItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	{
		val, err := d.DecodeOctetString(per.SizeConstraints{
			Extensible: false,
			Min:        nil,
			Max:        nil,
		})
		if err != nil {
			return err
		}
		ie.NonAnchorCarrierFrquency = val
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
