package ies

import (
	"github.com/lvdund/asn1go/per"
)

var qoSFlowsSetupMappedtoDRBModRqdSNterminatedItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "qoSFlowIdentifier"},
		{Name: "mCGRequestedGBRQoSFlowInfo", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type QoSFlowsSetupMappedtoDRBModRqdSNterminatedItem struct {
	QoSFlowIdentifier          QoSFlowIdentifier
	MCGRequestedGBRQoSFlowInfo *GBRQoSFlowInfo
	IEExtensions               []byte
}

func (ie *QoSFlowsSetupMappedtoDRBModRqdSNterminatedItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(qoSFlowsSetupMappedtoDRBModRqdSNterminatedItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.MCGRequestedGBRQoSFlowInfo != nil, false}); err != nil {
		return err
	}
	if err := ie.QoSFlowIdentifier.Encode(e); err != nil {
		return err
	}
	if ie.MCGRequestedGBRQoSFlowInfo != nil {
		if err := ie.MCGRequestedGBRQoSFlowInfo.Encode(e); err != nil {
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

func (ie *QoSFlowsSetupMappedtoDRBModRqdSNterminatedItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(qoSFlowsSetupMappedtoDRBModRqdSNterminatedItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.QoSFlowIdentifier.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.MCGRequestedGBRQoSFlowInfo = new(GBRQoSFlowInfo)
		if err := ie.MCGRequestedGBRQoSFlowInfo.Decode(d); err != nil {
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
