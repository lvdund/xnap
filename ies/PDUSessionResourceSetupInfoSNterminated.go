package ies

import (
	"github.com/lvdund/asn1go/per"
)

var pDUSessionResourceSetupInfoSNterminatedConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "uL-NG-U-TNLatUPF"},
		{Name: "pduSessionType"},
		{Name: "pduSessionNetworkInstance", Optional: true},
		{Name: "qosFlowsToBeSetup-List"},
		{Name: "dataforwardinginfofromSource", Optional: true},
		{Name: "securityIndication", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type PDUSessionResourceSetupInfoSNterminated struct {
	ULNGUTNLatUPF                UPTransportLayerInformation
	PduSessionType               PDUSessionType
	PduSessionNetworkInstance    *PDUSessionNetworkInstance
	QosFlowsToBeSetupList        QoSFlowsToBeSetupListSetupSNterminated
	DataforwardinginfofromSource *DataforwardingandOffloadingInfofromSource
	SecurityIndication           *SecurityIndication
	IEExtensions                 []byte
}

func (ie *PDUSessionResourceSetupInfoSNterminated) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDUSessionResourceSetupInfoSNterminatedConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.PduSessionNetworkInstance != nil, ie.DataforwardinginfofromSource != nil, ie.SecurityIndication != nil, false}); err != nil {
		return err
	}
	if err := ie.ULNGUTNLatUPF.Encode(e); err != nil {
		return err
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
	if ie.SecurityIndication != nil {
		if err := ie.SecurityIndication.Encode(e); err != nil {
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

func (ie *PDUSessionResourceSetupInfoSNterminated) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDUSessionResourceSetupInfoSNterminatedConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.ULNGUTNLatUPF.Decode(d); err != nil {
		return err
	}
	if err := ie.PduSessionType.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(2) {
		ie.PduSessionNetworkInstance = new(PDUSessionNetworkInstance)
		if err := ie.PduSessionNetworkInstance.Decode(d); err != nil {
			return err
		}
	}
	if err := ie.QosFlowsToBeSetupList.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(4) {
		ie.DataforwardinginfofromSource = new(DataforwardingandOffloadingInfofromSource)
		if err := ie.DataforwardinginfofromSource.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(5) {
		ie.SecurityIndication = new(SecurityIndication)
		if err := ie.SecurityIndication.Decode(d); err != nil {
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
