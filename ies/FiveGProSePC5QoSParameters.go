package ies

import (
	"github.com/lvdund/asn1go/per"
)

var fiveGProSePC5QoSParametersConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "fiveGProSepc5QoSFlowList"},
		{Name: "fiveGproSepc5LinkAggregateBitRates", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type FiveGProSePC5QoSParameters struct {
	FiveGProSepc5QoSFlowList           FiveGProSePC5QoSFlowList
	FiveGproSepc5LinkAggregateBitRates *BitRate
	IEExtensions                       []byte
}

func (ie *FiveGProSePC5QoSParameters) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(fiveGProSePC5QoSParametersConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.FiveGproSepc5LinkAggregateBitRates != nil, false}); err != nil {
		return err
	}
	if err := ie.FiveGProSepc5QoSFlowList.Encode(e); err != nil {
		return err
	}
	if ie.FiveGproSepc5LinkAggregateBitRates != nil {
		if err := ie.FiveGproSepc5LinkAggregateBitRates.Encode(e); err != nil {
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

func (ie *FiveGProSePC5QoSParameters) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(fiveGProSePC5QoSParametersConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.FiveGProSepc5QoSFlowList.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.FiveGproSepc5LinkAggregateBitRates = new(BitRate)
		if err := ie.FiveGproSepc5LinkAggregateBitRates.Decode(d); err != nil {
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
