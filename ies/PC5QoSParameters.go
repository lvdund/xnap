package ies

import (
	"github.com/lvdund/asn1go/per"
)

var pC5QoSParametersConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "pc5QoSFlowList"},
		{Name: "pc5LinkAggregateBitRates", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type PC5QoSParameters struct {
	Pc5QoSFlowList           PC5QoSFlowList
	Pc5LinkAggregateBitRates *BitRate
	IEExtensions             []byte
}

func (ie *PC5QoSParameters) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pC5QoSParametersConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.Pc5LinkAggregateBitRates != nil, false}); err != nil {
		return err
	}
	if err := ie.Pc5QoSFlowList.Encode(e); err != nil {
		return err
	}
	if ie.Pc5LinkAggregateBitRates != nil {
		if err := ie.Pc5LinkAggregateBitRates.Encode(e); err != nil {
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

func (ie *PC5QoSParameters) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pC5QoSParametersConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.Pc5QoSFlowList.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.Pc5LinkAggregateBitRates = new(BitRate)
		if err := ie.Pc5LinkAggregateBitRates.Decode(d); err != nil {
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
