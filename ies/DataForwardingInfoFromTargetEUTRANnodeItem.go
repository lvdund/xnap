package ies

import (
	"github.com/lvdund/asn1go/per"
)

var dataForwardingInfoFromTargetEUTRANnodeItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "dlForwardingUPTNLInformation"},
		{Name: "qosFlowsToBeForwarded-List"},
		{Name: "iE-Extension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type DataForwardingInfoFromTargetEUTRANnodeItem struct {
	DlForwardingUPTNLInformation UPTransportLayerInformation
	QosFlowsToBeForwardedList    QoSFlowsToBeForwardedList
	IEExtensions                 []byte
}

func (ie *DataForwardingInfoFromTargetEUTRANnodeItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dataForwardingInfoFromTargetEUTRANnodeItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.DlForwardingUPTNLInformation.Encode(e); err != nil {
		return err
	}
	if err := ie.QosFlowsToBeForwardedList.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *DataForwardingInfoFromTargetEUTRANnodeItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dataForwardingInfoFromTargetEUTRANnodeItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.DlForwardingUPTNLInformation.Decode(d); err != nil {
		return err
	}
	if err := ie.QosFlowsToBeForwardedList.Decode(d); err != nil {
		return err
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
