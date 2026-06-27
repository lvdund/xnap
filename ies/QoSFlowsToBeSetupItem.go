package ies

import (
	"github.com/lvdund/asn1go/per"
)

var qoSFlowsToBeSetupItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "qfi"},
		{Name: "qosFlowLevelQoSParameters"},
		{Name: "e-RAB-ID", Optional: true},
		{Name: "iE-Extension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type QoSFlowsToBeSetupItem struct {
	Qfi                       QoSFlowIdentifier
	QosFlowLevelQoSParameters QoSFlowLevelQoSParameters
	ERABID                    *ERABID
	IEExtensions              []byte
}

func (ie *QoSFlowsToBeSetupItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(qoSFlowsToBeSetupItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.ERABID != nil, false}); err != nil {
		return err
	}
	if err := ie.Qfi.Encode(e); err != nil {
		return err
	}
	if err := ie.QosFlowLevelQoSParameters.Encode(e); err != nil {
		return err
	}
	if ie.ERABID != nil {
		if err := ie.ERABID.Encode(e); err != nil {
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

func (ie *QoSFlowsToBeSetupItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(qoSFlowsToBeSetupItemConstraints)
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
		ie.ERABID = new(ERABID)
		if err := ie.ERABID.Decode(d); err != nil {
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
