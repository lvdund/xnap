package ies

import (
	"github.com/lvdund/asn1go/per"
)

var pDUSessionResourcesToBeSetupItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "pduSessionId"},
		{Name: "s-NSSAI"},
		{Name: "pduSessionAMBR", Optional: true},
		{Name: "uL-NG-U-TNLatUPF"},
		{Name: "source-DL-NG-U-TNL-Information", Optional: true},
		{Name: "securityIndication", Optional: true},
		{Name: "pduSessionType"},
		{Name: "pduSessionNetworkInstance", Optional: true},
		{Name: "qosFlowsToBeSetup-List"},
		{Name: "dataforwardinginfofromSource", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type PDUSessionResourcesToBeSetupItem struct {
	PduSessionId                 PDUSessionID
	SNSSAI                       SNSSAI
	PduSessionAMBR               *PDUSessionAggregateMaximumBitRate
	ULNGUTNLatUPF                UPTransportLayerInformation
	SourceDLNGUTNLInformation    *UPTransportLayerInformation
	SecurityIndication           *SecurityIndication
	PduSessionType               PDUSessionType
	PduSessionNetworkInstance    *PDUSessionNetworkInstance
	QosFlowsToBeSetupList        QoSFlowsToBeSetupList
	DataforwardinginfofromSource *DataforwardingandOffloadingInfofromSource
	IEExtensions                 []byte
}

func (ie *PDUSessionResourcesToBeSetupItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDUSessionResourcesToBeSetupItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.PduSessionAMBR != nil, ie.SourceDLNGUTNLInformation != nil, ie.SecurityIndication != nil, ie.PduSessionNetworkInstance != nil, ie.DataforwardinginfofromSource != nil, false}); err != nil {
		return err
	}
	if err := ie.PduSessionId.Encode(e); err != nil {
		return err
	}
	if err := ie.SNSSAI.Encode(e); err != nil {
		return err
	}
	if ie.PduSessionAMBR != nil {
		if err := ie.PduSessionAMBR.Encode(e); err != nil {
			return err
		}
	}
	if err := ie.ULNGUTNLatUPF.Encode(e); err != nil {
		return err
	}
	if ie.SourceDLNGUTNLInformation != nil {
		if err := ie.SourceDLNGUTNLInformation.Encode(e); err != nil {
			return err
		}
	}
	if ie.SecurityIndication != nil {
		if err := ie.SecurityIndication.Encode(e); err != nil {
			return err
		}
	}
	if err := ie.PduSessionType.Encode(e); err != nil {
		return err
	}
	if ie.PduSessionNetworkInstance != nil {
		if err := ie.PduSessionNetworkInstance.Encode(e); err != nil {
			return err
		}
	}
	if err := ie.QosFlowsToBeSetupList.Encode(e); err != nil {
		return err
	}
	if ie.DataforwardinginfofromSource != nil {
		if err := ie.DataforwardinginfofromSource.Encode(e); err != nil {
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

func (ie *PDUSessionResourcesToBeSetupItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDUSessionResourcesToBeSetupItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.PduSessionId.Decode(d); err != nil {
		return err
	}
	if err := ie.SNSSAI.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(2) {
		ie.PduSessionAMBR = new(PDUSessionAggregateMaximumBitRate)
		if err := ie.PduSessionAMBR.Decode(d); err != nil {
			return err
		}
	}
	if err := ie.ULNGUTNLatUPF.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(4) {
		ie.SourceDLNGUTNLInformation = new(UPTransportLayerInformation)
		if err := ie.SourceDLNGUTNLInformation.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(5) {
		ie.SecurityIndication = new(SecurityIndication)
		if err := ie.SecurityIndication.Decode(d); err != nil {
			return err
		}
	}
	if err := ie.PduSessionType.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(7) {
		ie.PduSessionNetworkInstance = new(PDUSessionNetworkInstance)
		if err := ie.PduSessionNetworkInstance.Decode(d); err != nil {
			return err
		}
	}
	if err := ie.QosFlowsToBeSetupList.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(9) {
		ie.DataforwardinginfofromSource = new(DataforwardingandOffloadingInfofromSource)
		if err := ie.DataforwardinginfofromSource.Decode(d); err != nil {
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
