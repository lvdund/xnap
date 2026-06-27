package ies

import (
	"github.com/lvdund/asn1go/per"
)

var a2XPC5QoSFlowItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "a2XpQI"},
		{Name: "a2Xpc5FlowBitRates", Optional: true},
		{Name: "a2Xrange", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type A2XPC5QoSFlowItem struct {
	A2XpQI             FiveQI
	A2Xpc5FlowBitRates *A2XPC5FlowBitRates
	A2Xrange           *Range
	IEExtensions       []byte
}

func (ie *A2XPC5QoSFlowItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(a2XPC5QoSFlowItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.A2Xpc5FlowBitRates != nil, ie.A2Xrange != nil, false}); err != nil {
		return err
	}
	if err := ie.A2XpQI.Encode(e); err != nil {
		return err
	}
	if ie.A2Xpc5FlowBitRates != nil {
		if err := ie.A2Xpc5FlowBitRates.Encode(e); err != nil {
			return err
		}
	}
	if ie.A2Xrange != nil {
		if err := ie.A2Xrange.Encode(e); err != nil {
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

func (ie *A2XPC5QoSFlowItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(a2XPC5QoSFlowItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.A2XpQI.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.A2Xpc5FlowBitRates = new(A2XPC5FlowBitRates)
		if err := ie.A2Xpc5FlowBitRates.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(2) {
		ie.A2Xrange = new(Range)
		if err := ie.A2Xrange.Decode(d); err != nil {
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
