package ies

import (
	"github.com/lvdund/asn1go/per"
)

var fiveGProSePC5FlowBitRatesConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "fiveGproSeguaranteedFlowBitRate"},
		{Name: "fiveGproSemaximumFlowBitRate"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type FiveGProSePC5FlowBitRates struct {
	FiveGproSeguaranteedFlowBitRate BitRate
	FiveGproSemaximumFlowBitRate    BitRate
	IEExtensions                    []byte
}

func (ie *FiveGProSePC5FlowBitRates) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(fiveGProSePC5FlowBitRatesConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.FiveGproSeguaranteedFlowBitRate.Encode(e); err != nil {
		return err
	}
	if err := ie.FiveGproSemaximumFlowBitRate.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *FiveGProSePC5FlowBitRates) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(fiveGProSePC5FlowBitRatesConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.FiveGproSeguaranteedFlowBitRate.Decode(d); err != nil {
		return err
	}
	if err := ie.FiveGproSemaximumFlowBitRate.Decode(d); err != nil {
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
