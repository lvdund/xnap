package ies

import (
	"github.com/lvdund/asn1go/per"
)

var pC5FlowBitRatesConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "guaranteedFlowBitRate"},
		{Name: "maximumFlowBitRate"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type PC5FlowBitRates struct {
	GuaranteedFlowBitRate BitRate
	MaximumFlowBitRate    BitRate
	IEExtensions          []byte
}

func (ie *PC5FlowBitRates) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pC5FlowBitRatesConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.GuaranteedFlowBitRate.Encode(e); err != nil {
		return err
	}
	if err := ie.MaximumFlowBitRate.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *PC5FlowBitRates) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pC5FlowBitRatesConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.GuaranteedFlowBitRate.Decode(d); err != nil {
		return err
	}
	if err := ie.MaximumFlowBitRate.Decode(d); err != nil {
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
