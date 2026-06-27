package ies

import (
	"github.com/lvdund/asn1go/per"
)

var extendedRATRestrictionInformationConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "primaryRATRestriction"},
		{Name: "secondaryRATRestriction"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type ExtendedRATRestrictionInformation struct {
	PrimaryRATRestriction   per.BitString
	SecondaryRATRestriction per.BitString
	IEExtensions            []byte
}

func (ie *ExtendedRATRestrictionInformation) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(extendedRATRestrictionInformationConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := e.EncodeBitString(ie.PrimaryRATRestriction, per.SizeConstraints{
		Extensible: false,
		Min:        nil,
		Max:        nil,
	}); err != nil {
		return err
	}
	if err := e.EncodeBitString(ie.SecondaryRATRestriction, per.SizeConstraints{
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

func (ie *ExtendedRATRestrictionInformation) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(extendedRATRestrictionInformationConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	{
		val, err := d.DecodeBitString(per.SizeConstraints{
			Extensible: false,
			Min:        nil,
			Max:        nil,
		})
		if err != nil {
			return err
		}
		ie.PrimaryRATRestriction = val
	}
	{
		val, err := d.DecodeBitString(per.SizeConstraints{
			Extensible: false,
			Min:        nil,
			Max:        nil,
		})
		if err != nil {
			return err
		}
		ie.SecondaryRATRestriction = val
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
