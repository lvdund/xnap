package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var sSBAreaRadioResourceStatusListItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sSBIndex"},
		{Name: "ssb-Area-DL-GBR-PRB-usage"},
		{Name: "ssb-Area-UL-GBR-PRB-usage"},
		{Name: "ssb-Area-dL-non-GBR-PRB-usage"},
		{Name: "ssb-Area-uL-non-GBR-PRB-usage"},
		{Name: "ssb-Area-dL-Total-PRB-usage"},
		{Name: "ssb-Area-uL-Total-PRB-usage"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type SSBAreaRadioResourceStatusListItem struct {
	SSBIndex                int64
	SsbAreaDLGBRPRBUsage    DLGBRPRBUsage
	SsbAreaULGBRPRBUsage    ULGBRPRBUsage
	SsbAreaDLNonGBRPRBUsage DLNonGBRPRBUsage
	SsbAreaULNonGBRPRBUsage ULNonGBRPRBUsage
	SsbAreaDLTotalPRBUsage  DLTotalPRBUsage
	SsbAreaULTotalPRBUsage  ULTotalPRBUsage
	IEExtensions            []byte
}

func (ie *SSBAreaRadioResourceStatusListItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sSBAreaRadioResourceStatusListItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := e.EncodeInteger(ie.SSBIndex, per.IntegerConstraints{
		Extensible: false,
		LowerBound: common.Ptr(int64(0)),
		UpperBound: common.Ptr(int64(63)),
	}); err != nil {
		return err
	}
	if err := ie.SsbAreaDLGBRPRBUsage.Encode(e); err != nil {
		return err
	}
	if err := ie.SsbAreaULGBRPRBUsage.Encode(e); err != nil {
		return err
	}
	if err := ie.SsbAreaDLNonGBRPRBUsage.Encode(e); err != nil {
		return err
	}
	if err := ie.SsbAreaULNonGBRPRBUsage.Encode(e); err != nil {
		return err
	}
	if err := ie.SsbAreaDLTotalPRBUsage.Encode(e); err != nil {
		return err
	}
	if err := ie.SsbAreaULTotalPRBUsage.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *SSBAreaRadioResourceStatusListItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sSBAreaRadioResourceStatusListItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	{
		val, err := d.DecodeInteger(per.IntegerConstraints{
			Extensible: false,
			LowerBound: common.Ptr(int64(0)),
			UpperBound: common.Ptr(int64(63)),
		})
		if err != nil {
			return err
		}
		ie.SSBIndex = val
	}
	if err := ie.SsbAreaDLGBRPRBUsage.Decode(d); err != nil {
		return err
	}
	if err := ie.SsbAreaULGBRPRBUsage.Decode(d); err != nil {
		return err
	}
	if err := ie.SsbAreaDLNonGBRPRBUsage.Decode(d); err != nil {
		return err
	}
	if err := ie.SsbAreaULNonGBRPRBUsage.Decode(d); err != nil {
		return err
	}
	if err := ie.SsbAreaDLTotalPRBUsage.Decode(d); err != nil {
		return err
	}
	if err := ie.SsbAreaULTotalPRBUsage.Decode(d); err != nil {
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
