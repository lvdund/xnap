package ies

import (
	"github.com/lvdund/asn1go/per"
)

var nPRACHConfigurationFDDConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "nprach-CP-length"},
		{Name: "anchorCarrier-NPRACHConfig"},
		{Name: "anchorCarrier-EDT-NPRACHConfig", Optional: true},
		{Name: "anchorCarrier-Format2-NPRACHConfig", Optional: true},
		{Name: "anchorCarrier-Format2-EDT-NPRACHConfig", Optional: true},
		{Name: "non-anchorCarrier-NPRACHConfig", Optional: true},
		{Name: "non-anchorCarrier-Format2-NPRACHConfig", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type NPRACHConfigurationFDD struct {
	NprachCPLength                      NPRACHCPLength
	AnchorCarrierNPRACHConfig           []byte
	AnchorCarrierEDTNPRACHConfig        []byte
	AnchorCarrierFormat2NPRACHConfig    []byte
	AnchorCarrierFormat2EDTNPRACHConfig []byte
	NonAnchorCarrierNPRACHConfig        []byte
	NonAnchorCarrierFormat2NPRACHConfig []byte
	IEExtensions                        []byte
}

func (ie *NPRACHConfigurationFDD) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(nPRACHConfigurationFDDConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{len(ie.AnchorCarrierEDTNPRACHConfig) > 0, len(ie.AnchorCarrierFormat2NPRACHConfig) > 0, len(ie.AnchorCarrierFormat2EDTNPRACHConfig) > 0, len(ie.NonAnchorCarrierNPRACHConfig) > 0, len(ie.NonAnchorCarrierFormat2NPRACHConfig) > 0, false}); err != nil {
		return err
	}
	if err := ie.NprachCPLength.Encode(e); err != nil {
		return err
	}
	if err := e.EncodeOctetString(ie.AnchorCarrierNPRACHConfig, per.SizeConstraints{
		Extensible: false,
		Min:        nil,
		Max:        nil,
	}); err != nil {
		return err
	}
	if len(ie.AnchorCarrierEDTNPRACHConfig) > 0 {
		if err := e.EncodeOctetString(ie.AnchorCarrierEDTNPRACHConfig, per.SizeConstraints{
			Extensible: false,
			Min:        nil,
			Max:        nil,
		}); err != nil {
			return err
		}
	}
	if len(ie.AnchorCarrierFormat2NPRACHConfig) > 0 {
		if err := e.EncodeOctetString(ie.AnchorCarrierFormat2NPRACHConfig, per.SizeConstraints{
			Extensible: false,
			Min:        nil,
			Max:        nil,
		}); err != nil {
			return err
		}
	}
	if len(ie.AnchorCarrierFormat2EDTNPRACHConfig) > 0 {
		if err := e.EncodeOctetString(ie.AnchorCarrierFormat2EDTNPRACHConfig, per.SizeConstraints{
			Extensible: false,
			Min:        nil,
			Max:        nil,
		}); err != nil {
			return err
		}
	}
	if len(ie.NonAnchorCarrierNPRACHConfig) > 0 {
		if err := e.EncodeOctetString(ie.NonAnchorCarrierNPRACHConfig, per.SizeConstraints{
			Extensible: false,
			Min:        nil,
			Max:        nil,
		}); err != nil {
			return err
		}
	}
	if len(ie.NonAnchorCarrierFormat2NPRACHConfig) > 0 {
		if err := e.EncodeOctetString(ie.NonAnchorCarrierFormat2NPRACHConfig, per.SizeConstraints{
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

func (ie *NPRACHConfigurationFDD) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(nPRACHConfigurationFDDConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.NprachCPLength.Decode(d); err != nil {
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
		ie.AnchorCarrierNPRACHConfig = val
	}
	if seq.IsComponentPresent(2) {
		val, err := d.DecodeOctetString(per.SizeConstraints{
			Extensible: false,
			Min:        nil,
			Max:        nil,
		})
		if err != nil {
			return err
		}
		ie.AnchorCarrierEDTNPRACHConfig = val
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
		ie.AnchorCarrierFormat2NPRACHConfig = val
	}
	if seq.IsComponentPresent(4) {
		val, err := d.DecodeOctetString(per.SizeConstraints{
			Extensible: false,
			Min:        nil,
			Max:        nil,
		})
		if err != nil {
			return err
		}
		ie.AnchorCarrierFormat2EDTNPRACHConfig = val
	}
	if seq.IsComponentPresent(5) {
		val, err := d.DecodeOctetString(per.SizeConstraints{
			Extensible: false,
			Min:        nil,
			Max:        nil,
		})
		if err != nil {
			return err
		}
		ie.NonAnchorCarrierNPRACHConfig = val
	}
	if seq.IsComponentPresent(6) {
		val, err := d.DecodeOctetString(per.SizeConstraints{
			Extensible: false,
			Min:        nil,
			Max:        nil,
		})
		if err != nil {
			return err
		}
		ie.NonAnchorCarrierFormat2NPRACHConfig = val
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
