package ies

import (
	"github.com/lvdund/asn1go/per"
)

var qoSFlowsToBeSetupListSetupSNterminatedItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "qfi"},
		{Name: "qosFlowLevelQoSParameters"},
		{Name: "offeredGBRQoSFlowInfo", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type QoSFlowsToBeSetupListSetupSNterminatedItem struct {
	Qfi                       QoSFlowIdentifier
	QosFlowLevelQoSParameters QoSFlowLevelQoSParameters
	OfferedGBRQoSFlowInfo     *GBRQoSFlowInfo
	IEExtensions              []byte
}

func (ie *QoSFlowsToBeSetupListSetupSNterminatedItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(qoSFlowsToBeSetupListSetupSNterminatedItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.OfferedGBRQoSFlowInfo != nil, false}); err != nil {
		return err
	}
	if err := ie.Qfi.Encode(e); err != nil {
		return err
	}
	if err := ie.QosFlowLevelQoSParameters.Encode(e); err != nil {
		return err
	}
	if ie.OfferedGBRQoSFlowInfo != nil {
		if err := ie.OfferedGBRQoSFlowInfo.Encode(e); err != nil {
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

func (ie *QoSFlowsToBeSetupListSetupSNterminatedItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(qoSFlowsToBeSetupListSetupSNterminatedItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.Qfi.Decode(d); err != nil {
		return err
	}
	if err := ie.QosFlowLevelQoSParameters.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(2) {
		ie.OfferedGBRQoSFlowInfo = new(GBRQoSFlowInfo)
		if err := ie.OfferedGBRQoSFlowInfo.Decode(d); err != nil {
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
