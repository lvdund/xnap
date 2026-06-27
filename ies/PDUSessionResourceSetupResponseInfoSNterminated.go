package ies

import (
	"github.com/lvdund/asn1go/per"
)

var pDUSessionResourceSetupResponseInfoSNterminatedConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "dL-NG-U-TNLatNG-RAN"},
		{Name: "dRBsToBeSetup", Optional: true},
		{Name: "dataforwardinginfoTarget", Optional: true},
		{Name: "qosFlowsNotAdmittedList", Optional: true},
		{Name: "securityResult", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type PDUSessionResourceSetupResponseInfoSNterminated struct {
	DLNGUTNLatNGRAN          UPTransportLayerInformation
	DRBsToBeSetup            *DRBsToBeSetupListSetupResponseSNterminated
	DataforwardinginfoTarget *DataForwardingInfoFromTargetNGRANnode
	QosFlowsNotAdmittedList  *QoSFlowsListWithCause
	SecurityResult           *SecurityResult
	IEExtensions             []byte
}

func (ie *PDUSessionResourceSetupResponseInfoSNterminated) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDUSessionResourceSetupResponseInfoSNterminatedConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.DRBsToBeSetup != nil, ie.DataforwardinginfoTarget != nil, ie.QosFlowsNotAdmittedList != nil, ie.SecurityResult != nil, false}); err != nil {
		return err
	}
	if err := ie.DLNGUTNLatNGRAN.Encode(e); err != nil {
		return err
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
	if ie.QosFlowsNotAdmittedList != nil {
		if err := ie.QosFlowsNotAdmittedList.Encode(e); err != nil {
			return err
		}
	}
	if ie.SecurityResult != nil {
		if err := ie.SecurityResult.Encode(e); err != nil {
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

func (ie *PDUSessionResourceSetupResponseInfoSNterminated) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDUSessionResourceSetupResponseInfoSNterminatedConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.DLNGUTNLatNGRAN.Decode(d); err != nil {
		return err
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
		ie.QosFlowsNotAdmittedList = new(QoSFlowsListWithCause)
		if err := ie.QosFlowsNotAdmittedList.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(4) {
		ie.SecurityResult = new(SecurityResult)
		if err := ie.SecurityResult.Decode(d); err != nil {
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
