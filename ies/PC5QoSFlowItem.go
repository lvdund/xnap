package ies

import (
	"github.com/lvdund/asn1go/per"
)

var pC5QoSFlowItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "pQI"},
		{Name: "pc5FlowBitRates", Optional: true},
		{Name: "range", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type PC5QoSFlowItem struct {
	PQI             FiveQI
	Pc5FlowBitRates *PC5FlowBitRates
	Range           *Range
	IEExtensions    []byte
}

func (ie *PC5QoSFlowItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pC5QoSFlowItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.Pc5FlowBitRates != nil, ie.Range != nil, false}); err != nil {
		return err
	}
	if err := ie.PQI.Encode(e); err != nil {
		return err
	}
	if ie.Pc5FlowBitRates != nil {
		if err := ie.Pc5FlowBitRates.Encode(e); err != nil {
			return err
		}
	}
	if ie.Range != nil {
		if err := ie.Range.Encode(e); err != nil {
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

func (ie *PC5QoSFlowItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pC5QoSFlowItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.PQI.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.Pc5FlowBitRates = new(PC5FlowBitRates)
		if err := ie.Pc5FlowBitRates.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(2) {
		ie.Range = new(Range)
		if err := ie.Range.Decode(d); err != nil {
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
