package ies

import (
	"github.com/lvdund/asn1go/per"
)

var alternativeQoSParaSetItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "alternativeQoSParaSetIndex"},
		{Name: "guaranteedFlowBitRateDL", Optional: true},
		{Name: "guaranteedFlowBitRateUL", Optional: true},
		{Name: "packetDelayBudget", Optional: true},
		{Name: "packetErrorRate", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type AlternativeQoSParaSetItem struct {
	AlternativeQoSParaSetIndex QoSParaSetIndex
	GuaranteedFlowBitRateDL    *BitRate
	GuaranteedFlowBitRateUL    *BitRate
	PacketDelayBudget          *PacketDelayBudget
	PacketErrorRate            *PacketErrorRate
	IEExtensions               []byte
}

func (ie *AlternativeQoSParaSetItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(alternativeQoSParaSetItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.GuaranteedFlowBitRateDL != nil, ie.GuaranteedFlowBitRateUL != nil, ie.PacketDelayBudget != nil, ie.PacketErrorRate != nil, false}); err != nil {
		return err
	}
	if err := ie.AlternativeQoSParaSetIndex.Encode(e); err != nil {
		return err
	}
	if ie.GuaranteedFlowBitRateDL != nil {
		if err := ie.GuaranteedFlowBitRateDL.Encode(e); err != nil {
			return err
		}
	}
	if ie.GuaranteedFlowBitRateUL != nil {
		if err := ie.GuaranteedFlowBitRateUL.Encode(e); err != nil {
			return err
		}
	}
	if ie.PacketDelayBudget != nil {
		if err := ie.PacketDelayBudget.Encode(e); err != nil {
			return err
		}
	}
	if ie.PacketErrorRate != nil {
		if err := ie.PacketErrorRate.Encode(e); err != nil {
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

func (ie *AlternativeQoSParaSetItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(alternativeQoSParaSetItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.AlternativeQoSParaSetIndex.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.GuaranteedFlowBitRateDL = new(BitRate)
		if err := ie.GuaranteedFlowBitRateDL.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(2) {
		ie.GuaranteedFlowBitRateUL = new(BitRate)
		if err := ie.GuaranteedFlowBitRateUL.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(3) {
		ie.PacketDelayBudget = new(PacketDelayBudget)
		if err := ie.PacketDelayBudget.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(4) {
		ie.PacketErrorRate = new(PacketErrorRate)
		if err := ie.PacketErrorRate.Decode(d); err != nil {
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
