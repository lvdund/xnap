package ies

import (
	"github.com/lvdund/asn1go/per"
)

var rSPPTransportQoSParametersConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "rSPPQoSFlowList"},
		{Name: "rSPPLinkAggregateBitRates", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type RSPPTransportQoSParameters struct {
	RSPPQoSFlowList           RSPPQoSFlowList
	RSPPLinkAggregateBitRates *BitRate
	IEExtensions              []byte
}

func (ie *RSPPTransportQoSParameters) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rSPPTransportQoSParametersConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.RSPPLinkAggregateBitRates != nil, false}); err != nil {
		return err
	}
	if err := ie.RSPPQoSFlowList.Encode(e); err != nil {
		return err
	}
	if ie.RSPPLinkAggregateBitRates != nil {
		if err := ie.RSPPLinkAggregateBitRates.Encode(e); err != nil {
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

func (ie *RSPPTransportQoSParameters) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rSPPTransportQoSParametersConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.RSPPQoSFlowList.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.RSPPLinkAggregateBitRates = new(BitRate)
		if err := ie.RSPPLinkAggregateBitRates.Decode(d); err != nil {
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
