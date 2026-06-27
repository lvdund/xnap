package ies

import (
	"github.com/lvdund/asn1go/per"
)

var pDUSessionResourceModificationInfoSNterminatedConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "uL-NG-U-TNLatUPF", Optional: true},
		{Name: "pduSessionNetworkInstance", Optional: true},
		{Name: "qosFlowsToBeSetup-List", Optional: true},
		{Name: "dataforwardinginfofromSource", Optional: true},
		{Name: "qosFlowsToBeModified-List", Optional: true},
		{Name: "qoSFlowsToBeReleased-List", Optional: true},
		{Name: "drbsToBeModifiedList", Optional: true},
		{Name: "dRBsToBeReleased", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type PDUSessionResourceModificationInfoSNterminated struct {
	ULNGUTNLatUPF                *UPTransportLayerInformation
	PduSessionNetworkInstance    *PDUSessionNetworkInstance
	QosFlowsToBeSetupList        *QoSFlowsToBeSetupListSetupSNterminated
	DataforwardinginfofromSource *DataforwardingandOffloadingInfofromSource
	QosFlowsToBeModifiedList     *QoSFlowsToBeSetupListModifiedSNterminated
	QoSFlowsToBeReleasedList     *QoSFlowsListWithCause
	DrbsToBeModifiedList         *DRBsToBeModifiedListModifiedSNterminated
	DRBsToBeReleased             *DRBListWithCause
	IEExtensions                 []byte
}

func (ie *PDUSessionResourceModificationInfoSNterminated) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDUSessionResourceModificationInfoSNterminatedConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.ULNGUTNLatUPF != nil, ie.PduSessionNetworkInstance != nil, ie.QosFlowsToBeSetupList != nil, ie.DataforwardinginfofromSource != nil, ie.QosFlowsToBeModifiedList != nil, ie.QoSFlowsToBeReleasedList != nil, ie.DrbsToBeModifiedList != nil, ie.DRBsToBeReleased != nil, false}); err != nil {
		return err
	}
	if ie.ULNGUTNLatUPF != nil {
		if err := ie.ULNGUTNLatUPF.Encode(e); err != nil {
			return err
		}
	}
	if ie.PduSessionNetworkInstance != nil {
		if err := ie.PduSessionNetworkInstance.Encode(e); err != nil {
			return err
		}
	}
	if ie.QosFlowsToBeSetupList != nil {
		if err := ie.QosFlowsToBeSetupList.Encode(e); err != nil {
			return err
		}
	}
	if ie.DataforwardinginfofromSource != nil {
		if err := ie.DataforwardinginfofromSource.Encode(e); err != nil {
			return err
		}
	}
	if ie.QosFlowsToBeModifiedList != nil {
		if err := ie.QosFlowsToBeModifiedList.Encode(e); err != nil {
			return err
		}
	}
	if ie.QoSFlowsToBeReleasedList != nil {
		if err := ie.QoSFlowsToBeReleasedList.Encode(e); err != nil {
			return err
		}
	}
	if ie.DrbsToBeModifiedList != nil {
		if err := ie.DrbsToBeModifiedList.Encode(e); err != nil {
			return err
		}
	}
	if ie.DRBsToBeReleased != nil {
		if err := ie.DRBsToBeReleased.Encode(e); err != nil {
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

func (ie *PDUSessionResourceModificationInfoSNterminated) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDUSessionResourceModificationInfoSNterminatedConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if seq.IsComponentPresent(0) {
		ie.ULNGUTNLatUPF = new(UPTransportLayerInformation)
		if err := ie.ULNGUTNLatUPF.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(1) {
		ie.PduSessionNetworkInstance = new(PDUSessionNetworkInstance)
		if err := ie.PduSessionNetworkInstance.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(2) {
		ie.QosFlowsToBeSetupList = new(QoSFlowsToBeSetupListSetupSNterminated)
		if err := ie.QosFlowsToBeSetupList.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(3) {
		ie.DataforwardinginfofromSource = new(DataforwardingandOffloadingInfofromSource)
		if err := ie.DataforwardinginfofromSource.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(4) {
		ie.QosFlowsToBeModifiedList = new(QoSFlowsToBeSetupListModifiedSNterminated)
		if err := ie.QosFlowsToBeModifiedList.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(5) {
		ie.QoSFlowsToBeReleasedList = new(QoSFlowsListWithCause)
		if err := ie.QoSFlowsToBeReleasedList.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(6) {
		ie.DrbsToBeModifiedList = new(DRBsToBeModifiedListModifiedSNterminated)
		if err := ie.DrbsToBeModifiedList.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(7) {
		ie.DRBsToBeReleased = new(DRBListWithCause)
		if err := ie.DRBsToBeReleased.Decode(d); err != nil {
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
