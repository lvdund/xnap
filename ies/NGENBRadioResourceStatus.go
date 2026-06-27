package ies

import (
	"github.com/lvdund/asn1go/per"
)

var nGENBRadioResourceStatusConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "dL-GBR-PRB-usage"},
		{Name: "uL-GBR-PRB-usage"},
		{Name: "dL-non-GBR-PRB-usage"},
		{Name: "uL-non-GBR-PRB-usage"},
		{Name: "dL-Total-PRB-usage"},
		{Name: "uL-Total-PRB-usage"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type NGENBRadioResourceStatus struct {
	DLGBRPRBUsage    DLGBRPRBUsage
	ULGBRPRBUsage    ULGBRPRBUsage
	DLNonGBRPRBUsage DLNonGBRPRBUsage
	ULNonGBRPRBUsage ULNonGBRPRBUsage
	DLTotalPRBUsage  DLTotalPRBUsage
	ULTotalPRBUsage  ULTotalPRBUsage
	IEExtensions     []byte
}

func (ie *NGENBRadioResourceStatus) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(nGENBRadioResourceStatusConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.DLGBRPRBUsage.Encode(e); err != nil {
		return err
	}
	if err := ie.ULGBRPRBUsage.Encode(e); err != nil {
		return err
	}
	if err := ie.DLNonGBRPRBUsage.Encode(e); err != nil {
		return err
	}
	if err := ie.ULNonGBRPRBUsage.Encode(e); err != nil {
		return err
	}
	if err := ie.DLTotalPRBUsage.Encode(e); err != nil {
		return err
	}
	if err := ie.ULTotalPRBUsage.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *NGENBRadioResourceStatus) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(nGENBRadioResourceStatusConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.DLGBRPRBUsage.Decode(d); err != nil {
		return err
	}
	if err := ie.ULGBRPRBUsage.Decode(d); err != nil {
		return err
	}
	if err := ie.DLNonGBRPRBUsage.Decode(d); err != nil {
		return err
	}
	if err := ie.ULNonGBRPRBUsage.Decode(d); err != nil {
		return err
	}
	if err := ie.DLTotalPRBUsage.Decode(d); err != nil {
		return err
	}
	if err := ie.ULTotalPRBUsage.Decode(d); err != nil {
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
