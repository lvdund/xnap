package ies

import (
	"github.com/lvdund/asn1go/per"
)

var fiveGProSePC5QoSFlowItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "fiveGproSepQI"},
		{Name: "fiveGproSepc5FlowBitRates", Optional: true},
		{Name: "fiveGproSerange", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type FiveGProSePC5QoSFlowItem struct {
	FiveGproSepQI             FiveQI
	FiveGproSepc5FlowBitRates *FiveGProSePC5FlowBitRates
	FiveGproSerange           *Range
	IEExtensions              []byte
}

func (ie *FiveGProSePC5QoSFlowItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(fiveGProSePC5QoSFlowItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.FiveGproSepc5FlowBitRates != nil, ie.FiveGproSerange != nil, false}); err != nil {
		return err
	}
	if err := ie.FiveGproSepQI.Encode(e); err != nil {
		return err
	}
	if ie.FiveGproSepc5FlowBitRates != nil {
		if err := ie.FiveGproSepc5FlowBitRates.Encode(e); err != nil {
			return err
		}
	}
	if ie.FiveGproSerange != nil {
		if err := ie.FiveGproSerange.Encode(e); err != nil {
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

func (ie *FiveGProSePC5QoSFlowItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(fiveGProSePC5QoSFlowItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.FiveGproSepQI.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.FiveGproSepc5FlowBitRates = new(FiveGProSePC5FlowBitRates)
		if err := ie.FiveGproSepc5FlowBitRates.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(2) {
		ie.FiveGproSerange = new(Range)
		if err := ie.FiveGproSerange.Decode(d); err != nil {
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
