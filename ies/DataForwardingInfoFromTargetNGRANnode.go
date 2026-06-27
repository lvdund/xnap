package ies

import (
	"github.com/lvdund/asn1go/per"
)

var dataForwardingInfoFromTargetNGRANnodeConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "qosFlowsAcceptedForDataForwarding-List"},
		{Name: "pduSessionLevelDLDataForwardingInfo", Optional: true},
		{Name: "pduSessionLevelULDataForwardingInfo", Optional: true},
		{Name: "dataForwardingResponseDRBItemList", Optional: true},
		{Name: "iE-Extension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type DataForwardingInfoFromTargetNGRANnode struct {
	QosFlowsAcceptedForDataForwardingList QoSFLowsAcceptedToBeForwardedList
	PduSessionLevelDLDataForwardingInfo   *UPTransportLayerInformation
	PduSessionLevelULDataForwardingInfo   *UPTransportLayerInformation
	DataForwardingResponseDRBItemList     *DataForwardingResponseDRBItemList
	IEExtensions                          []byte
}

func (ie *DataForwardingInfoFromTargetNGRANnode) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dataForwardingInfoFromTargetNGRANnodeConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.PduSessionLevelDLDataForwardingInfo != nil, ie.PduSessionLevelULDataForwardingInfo != nil, ie.DataForwardingResponseDRBItemList != nil, false}); err != nil {
		return err
	}
	if err := ie.QosFlowsAcceptedForDataForwardingList.Encode(e); err != nil {
		return err
	}
	if ie.PduSessionLevelDLDataForwardingInfo != nil {
		if err := ie.PduSessionLevelDLDataForwardingInfo.Encode(e); err != nil {
			return err
		}
	}
	if ie.PduSessionLevelULDataForwardingInfo != nil {
		if err := ie.PduSessionLevelULDataForwardingInfo.Encode(e); err != nil {
			return err
		}
	}
	if ie.DataForwardingResponseDRBItemList != nil {
		if err := ie.DataForwardingResponseDRBItemList.Encode(e); err != nil {
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

func (ie *DataForwardingInfoFromTargetNGRANnode) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dataForwardingInfoFromTargetNGRANnodeConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.QosFlowsAcceptedForDataForwardingList.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.PduSessionLevelDLDataForwardingInfo = new(UPTransportLayerInformation)
		if err := ie.PduSessionLevelDLDataForwardingInfo.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(2) {
		ie.PduSessionLevelULDataForwardingInfo = new(UPTransportLayerInformation)
		if err := ie.PduSessionLevelULDataForwardingInfo.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(3) {
		ie.DataForwardingResponseDRBItemList = new(DataForwardingResponseDRBItemList)
		if err := ie.DataForwardingResponseDRBItemList.Decode(d); err != nil {
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
