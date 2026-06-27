package ies

import (
	"github.com/lvdund/asn1go/per"
)

var a2XPC5QoSParametersConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "a2XPC5QoSFlowList"},
		{Name: "aA2XPC5LinkAggregateBitRates", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type A2XPC5QoSParameters struct {
	A2XPC5QoSFlowList            A2XPC5QoSFlowList
	AA2XPC5LinkAggregateBitRates *BitRate
	IEExtensions                 []byte
}

func (ie *A2XPC5QoSParameters) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(a2XPC5QoSParametersConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.AA2XPC5LinkAggregateBitRates != nil, false}); err != nil {
		return err
	}
	if err := ie.A2XPC5QoSFlowList.Encode(e); err != nil {
		return err
	}
	if ie.AA2XPC5LinkAggregateBitRates != nil {
		if err := ie.AA2XPC5LinkAggregateBitRates.Encode(e); err != nil {
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

func (ie *A2XPC5QoSParameters) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(a2XPC5QoSParametersConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.A2XPC5QoSFlowList.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.AA2XPC5LinkAggregateBitRates = new(BitRate)
		if err := ie.AA2XPC5LinkAggregateBitRates.Decode(d); err != nil {
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
