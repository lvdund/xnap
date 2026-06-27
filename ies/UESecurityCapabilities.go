package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var uESecurityCapabilitiesConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "nr-EncyptionAlgorithms"},
		{Name: "nr-IntegrityProtectionAlgorithms"},
		{Name: "e-utra-EncyptionAlgorithms"},
		{Name: "e-utra-IntegrityProtectionAlgorithms"},
		{Name: "iE-Extension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type UESecurityCapabilities struct {
	NrEncyptionAlgorithms              per.BitString
	NrIntegrityProtectionAlgorithms    per.BitString
	EUtraEncyptionAlgorithms           per.BitString
	EUtraIntegrityProtectionAlgorithms per.BitString
	IEExtensions                       []byte
}

var uESecurityCapabilitiesBitStringSize = common.Ptr(int64(16))

func (ie *UESecurityCapabilities) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uESecurityCapabilitiesConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := e.EncodeBitString(ie.NrEncyptionAlgorithms, per.SizeConstraints{
		Extensible: true,
		Min:        uESecurityCapabilitiesBitStringSize,
		Max:        uESecurityCapabilitiesBitStringSize,
	}); err != nil {
		return err
	}
	if err := e.EncodeBitString(ie.NrIntegrityProtectionAlgorithms, per.SizeConstraints{
		Extensible: true,
		Min:        uESecurityCapabilitiesBitStringSize,
		Max:        uESecurityCapabilitiesBitStringSize,
	}); err != nil {
		return err
	}
	if err := e.EncodeBitString(ie.EUtraEncyptionAlgorithms, per.SizeConstraints{
		Extensible: true,
		Min:        uESecurityCapabilitiesBitStringSize,
		Max:        uESecurityCapabilitiesBitStringSize,
	}); err != nil {
		return err
	}
	if err := e.EncodeBitString(ie.EUtraIntegrityProtectionAlgorithms, per.SizeConstraints{
		Extensible: true,
		Min:        uESecurityCapabilitiesBitStringSize,
		Max:        uESecurityCapabilitiesBitStringSize,
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

func (ie *UESecurityCapabilities) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uESecurityCapabilitiesConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	{
		val, err := d.DecodeBitString(per.SizeConstraints{
			Extensible: true,
			Min:        uESecurityCapabilitiesBitStringSize,
			Max:        uESecurityCapabilitiesBitStringSize,
		})
		if err != nil {
			return err
		}
		ie.NrEncyptionAlgorithms = val
	}
	{
		val, err := d.DecodeBitString(per.SizeConstraints{
			Extensible: true,
			Min:        uESecurityCapabilitiesBitStringSize,
			Max:        uESecurityCapabilitiesBitStringSize,
		})
		if err != nil {
			return err
		}
		ie.NrIntegrityProtectionAlgorithms = val
	}
	{
		val, err := d.DecodeBitString(per.SizeConstraints{
			Extensible: true,
			Min:        uESecurityCapabilitiesBitStringSize,
			Max:        uESecurityCapabilitiesBitStringSize,
		})
		if err != nil {
			return err
		}
		ie.EUtraEncyptionAlgorithms = val
	}
	{
		val, err := d.DecodeBitString(per.SizeConstraints{
			Extensible: true,
			Min:        uESecurityCapabilitiesBitStringSize,
			Max:        uESecurityCapabilitiesBitStringSize,
		})
		if err != nil {
			return err
		}
		ie.EUtraIntegrityProtectionAlgorithms = val
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
