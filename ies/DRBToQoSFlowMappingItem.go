package ies

import (
	"github.com/lvdund/asn1go/per"
)

var dRBToQoSFlowMappingItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "drb-ID"},
		{Name: "qosFlows-List"},
		{Name: "rLC-Mode", Optional: true},
		{Name: "iE-Extension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type DRBToQoSFlowMappingItem struct {
	DrbID        DRBID
	QosFlowsList QoSFlowsList
	RLCMode      *RLCMode
	IEExtensions []byte
}

func (ie *DRBToQoSFlowMappingItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dRBToQoSFlowMappingItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.RLCMode != nil, false}); err != nil {
		return err
	}
	if err := ie.DrbID.Encode(e); err != nil {
		return err
	}
	if err := ie.QosFlowsList.Encode(e); err != nil {
		return err
	}
	if ie.RLCMode != nil {
		if err := ie.RLCMode.Encode(e); err != nil {
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

func (ie *DRBToQoSFlowMappingItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dRBToQoSFlowMappingItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.DrbID.Decode(d); err != nil {
		return err
	}
	if err := ie.QosFlowsList.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(2) {
		ie.RLCMode = new(RLCMode)
		if err := ie.RLCMode.Decode(d); err != nil {
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
