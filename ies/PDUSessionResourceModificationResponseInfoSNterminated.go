package ies

import (
	"github.com/lvdund/asn1go/per"
)

var pDUSessionResourceModificationResponseInfoSNterminatedConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "dL-NG-U-TNLatNG-RAN", Optional: true},
		{Name: "dRBsToBeSetup", Optional: true},
		{Name: "dataforwardinginfoTarget", Optional: true},
		{Name: "dRBsToBeModified", Optional: true},
		{Name: "dRBsToBeReleased", Optional: true},
		{Name: "dataforwardinginfofromSource", Optional: true},
		{Name: "qosFlowsNotAdmittedTBAdded", Optional: true},
		{Name: "qosFlowsReleased", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type PDUSessionResourceModificationResponseInfoSNterminated struct {
	DLNGUTNLatNGRAN              *UPTransportLayerInformation
	DRBsToBeSetup                *DRBsToBeSetupListSetupResponseSNterminated
	DataforwardinginfoTarget     *DataForwardingInfoFromTargetNGRANnode
	DRBsToBeModified             *DRBsToBeModifiedListModificationResponseSNterminated
	DRBsToBeReleased             *DRBListWithCause
	DataforwardinginfofromSource *DataforwardingandOffloadingInfofromSource
	QosFlowsNotAdmittedTBAdded   *QoSFlowsListWithCause
	QosFlowsReleased             *QoSFlowsListWithCause
	IEExtensions                 []byte
}

func (ie *PDUSessionResourceModificationResponseInfoSNterminated) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDUSessionResourceModificationResponseInfoSNterminatedConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.DLNGUTNLatNGRAN != nil, ie.DRBsToBeSetup != nil, ie.DataforwardinginfoTarget != nil, ie.DRBsToBeModified != nil, ie.DRBsToBeReleased != nil, ie.DataforwardinginfofromSource != nil, ie.QosFlowsNotAdmittedTBAdded != nil, ie.QosFlowsReleased != nil, false}); err != nil {
		return err
	}
	if ie.DLNGUTNLatNGRAN != nil {
		if err := ie.DLNGUTNLatNGRAN.Encode(e); err != nil {
			return err
		}
	}
	if ie.DRBsToBeSetup != nil {
		if err := ie.DRBsToBeSetup.Encode(e); err != nil {
			return err
		}
	}
	if ie.DataforwardinginfoTarget != nil {
		if err := ie.DataforwardinginfoTarget.Encode(e); err != nil {
			return err
		}
	}
	if ie.DRBsToBeModified != nil {
		if err := ie.DRBsToBeModified.Encode(e); err != nil {
			return err
		}
	}
	if ie.DRBsToBeReleased != nil {
		if err := ie.DRBsToBeReleased.Encode(e); err != nil {
			return err
		}
	}
	if ie.DataforwardinginfofromSource != nil {
		if err := ie.DataforwardinginfofromSource.Encode(e); err != nil {
			return err
		}
	}
	if ie.QosFlowsNotAdmittedTBAdded != nil {
		if err := ie.QosFlowsNotAdmittedTBAdded.Encode(e); err != nil {
			return err
		}
	}
	if ie.QosFlowsReleased != nil {
		if err := ie.QosFlowsReleased.Encode(e); err != nil {
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

func (ie *PDUSessionResourceModificationResponseInfoSNterminated) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDUSessionResourceModificationResponseInfoSNterminatedConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if seq.IsComponentPresent(0) {
		ie.DLNGUTNLatNGRAN = new(UPTransportLayerInformation)
		if err := ie.DLNGUTNLatNGRAN.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(1) {
		ie.DRBsToBeSetup = new(DRBsToBeSetupListSetupResponseSNterminated)
		if err := ie.DRBsToBeSetup.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(2) {
		ie.DataforwardinginfoTarget = new(DataForwardingInfoFromTargetNGRANnode)
		if err := ie.DataforwardinginfoTarget.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(3) {
		ie.DRBsToBeModified = new(DRBsToBeModifiedListModificationResponseSNterminated)
		if err := ie.DRBsToBeModified.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(4) {
		ie.DRBsToBeReleased = new(DRBListWithCause)
		if err := ie.DRBsToBeReleased.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(5) {
		ie.DataforwardinginfofromSource = new(DataforwardingandOffloadingInfofromSource)
		if err := ie.DataforwardinginfofromSource.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(6) {
		ie.QosFlowsNotAdmittedTBAdded = new(QoSFlowsListWithCause)
		if err := ie.QosFlowsNotAdmittedTBAdded.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(7) {
		ie.QosFlowsReleased = new(QoSFlowsListWithCause)
		if err := ie.QosFlowsReleased.Decode(d); err != nil {
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
