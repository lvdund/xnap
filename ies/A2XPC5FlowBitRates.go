package ies

import (
	"github.com/lvdund/asn1go/per"
)

var a2XPC5FlowBitRatesConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "a2XguaranteedFlowBitRate"},
		{Name: "a2XmaximumFlowBitRate"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type A2XPC5FlowBitRates struct {
	A2XguaranteedFlowBitRate BitRate
	A2XmaximumFlowBitRate    BitRate
	IEExtensions             []byte
}

func (ie *A2XPC5FlowBitRates) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(a2XPC5FlowBitRatesConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.A2XguaranteedFlowBitRate.Encode(e); err != nil {
		return err
	}
	if err := ie.A2XmaximumFlowBitRate.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *A2XPC5FlowBitRates) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(a2XPC5FlowBitRatesConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.A2XguaranteedFlowBitRate.Decode(d); err != nil {
		return err
	}
	if err := ie.A2XmaximumFlowBitRate.Decode(d); err != nil {
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
