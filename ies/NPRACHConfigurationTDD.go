package ies

import (
	"github.com/lvdund/asn1go/per"
)

var nPRACHConfigurationTDDConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "nprach-preambleFormat"},
		{Name: "anchorCarrier-NPRACHConfigTDD"},
		{Name: "non-anchorCarrierFequencyConfiglist", Optional: true},
		{Name: "non-anchorCarrier-NPRACHConfigTDD", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type NPRACHConfigurationTDD struct {
	NprachPreambleFormat               NPRACHPreambleFormat
	AnchorCarrierNPRACHConfigTDD       []byte
	NonAnchorCarrierFequencyConfiglist *NonAnchorCarrierFrequencylist
	NonAnchorCarrierNPRACHConfigTDD    []byte
	IEExtensions                       []byte
}

func (ie *NPRACHConfigurationTDD) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(nPRACHConfigurationTDDConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.NonAnchorCarrierFequencyConfiglist != nil, len(ie.NonAnchorCarrierNPRACHConfigTDD) > 0, false}); err != nil {
		return err
	}
	if err := ie.NprachPreambleFormat.Encode(e); err != nil {
		return err
	}
	if err := e.EncodeOctetString(ie.AnchorCarrierNPRACHConfigTDD, per.SizeConstraints{
		Extensible: false,
		Min:        nil,
		Max:        nil,
	}); err != nil {
		return err
	}
	if ie.NonAnchorCarrierFequencyConfiglist != nil {
		if err := ie.NonAnchorCarrierFequencyConfiglist.Encode(e); err != nil {
			return err
		}
	}
	if len(ie.NonAnchorCarrierNPRACHConfigTDD) > 0 {
		if err := e.EncodeOctetString(ie.NonAnchorCarrierNPRACHConfigTDD, per.SizeConstraints{
			Extensible: false,
			Min:        nil,
			Max:        nil,
		}); err != nil {
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

func (ie *NPRACHConfigurationTDD) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(nPRACHConfigurationTDDConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.NprachPreambleFormat.Decode(d); err != nil {
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
		ie.AnchorCarrierNPRACHConfigTDD = val
	}
	if seq.IsComponentPresent(2) {
		ie.NonAnchorCarrierFequencyConfiglist = new(NonAnchorCarrierFrequencylist)
		if err := ie.NonAnchorCarrierFequencyConfiglist.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(3) {
		val, err := d.DecodeOctetString(per.SizeConstraints{
			Extensible: false,
			Min:        nil,
			Max:        nil,
		})
		if err != nil {
			return err
		}
		ie.NonAnchorCarrierNPRACHConfigTDD = val
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
