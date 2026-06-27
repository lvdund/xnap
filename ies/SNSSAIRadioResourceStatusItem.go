package ies

import (
	"github.com/lvdund/asn1go/per"
)

var sNSSAIRadioResourceStatusItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sNSSAI"},
		{Name: "slice-DL-GBR-PRB-Usage"},
		{Name: "slice-UL-GBR-PRB-Usage"},
		{Name: "slice-DL-non-GBR-PRB-Usage"},
		{Name: "slice-UL-non-GBR-PRB-Usage"},
		{Name: "slice-DL-Total-PRB-Allocation"},
		{Name: "slice-UL-Total-PRB-Allocation"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type SNSSAIRadioResourceStatusItem struct {
	SNSSAI                    SNSSAI
	SliceDLGBRPRBUsage        SliceDLGBRPRBUsage
	SliceULGBRPRBUsage        SliceULGBRPRBUsage
	SliceDLNonGBRPRBUsage     SliceDLNonGBRPRBUsage
	SliceULNonGBRPRBUsage     SliceULNonGBRPRBUsage
	SliceDLTotalPRBAllocation SliceDLTotalPRBAllocation
	SliceULTotalPRBAllocation SliceULTotalPRBAllocation
	IEExtensions              []byte
}

func (ie *SNSSAIRadioResourceStatusItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sNSSAIRadioResourceStatusItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.SNSSAI.Encode(e); err != nil {
		return err
	}
	if err := ie.SliceDLGBRPRBUsage.Encode(e); err != nil {
		return err
	}
	if err := ie.SliceULGBRPRBUsage.Encode(e); err != nil {
		return err
	}
	if err := ie.SliceDLNonGBRPRBUsage.Encode(e); err != nil {
		return err
	}
	if err := ie.SliceULNonGBRPRBUsage.Encode(e); err != nil {
		return err
	}
	if err := ie.SliceDLTotalPRBAllocation.Encode(e); err != nil {
		return err
	}
	if err := ie.SliceULTotalPRBAllocation.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *SNSSAIRadioResourceStatusItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sNSSAIRadioResourceStatusItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.SNSSAI.Decode(d); err != nil {
		return err
	}
	if err := ie.SliceDLGBRPRBUsage.Decode(d); err != nil {
		return err
	}
	if err := ie.SliceULGBRPRBUsage.Decode(d); err != nil {
		return err
	}
	if err := ie.SliceDLNonGBRPRBUsage.Decode(d); err != nil {
		return err
	}
	if err := ie.SliceULNonGBRPRBUsage.Decode(d); err != nil {
		return err
	}
	if err := ie.SliceDLTotalPRBAllocation.Decode(d); err != nil {
		return err
	}
	if err := ie.SliceULTotalPRBAllocation.Decode(d); err != nil {
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
