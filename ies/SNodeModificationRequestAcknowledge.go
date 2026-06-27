package ies

import (
	"fmt"
	"io"

	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

// SNodeModificationRequestAcknowledge ::= SEQUENCE {
//	protocolIEs  ProtocolIE-Container {{SNodeModificationRequestAcknowledge-IEs}},
//	...
// }

type SNodeModificationRequestAcknowledge struct {
	MNGRANnodeUEXnAPID                    NGRANnodeUEXnAPID
	SNGRANnodeUEXnAPID                    NGRANnodeUEXnAPID
	PDUSessionAdmittedSNModResponse       *PDUSessionAdmittedSNModResponse
	PDUSessionNotAdmittedSNModResponse    *PDUSessionNotAdmittedSNModResponse
	SNToMNContainer                       []byte
	AdmittedSplitSRB                      *SplitSRBsTypes
	AdmittedSplitSRBrelease               *SplitSRBsTypes
	CriticalityDiagnostics                *CriticalityDiagnostics
	LocationInformationSN                 *TargetCGI
	MRDCResourceCoordinationInfo          *MRDCResourceCoordinationInfo
	PDUSessionDataForwardingSNModResponse *PDUSessionDataForwardingSNModResponse
	RRCConfigIndication                   *RRCConfigIndication
	AvailableFastMCGRecoveryViaSRB3       *AvailableFastMCGRecoveryViaSRB3
	ReleaseFastMCGRecoveryViaSRB3         *ReleaseFastMCGRecoveryViaSRB3
	DirectForwardingPathAvailability      *DirectForwardingPathAvailability
	SCGUEHistoryInformation               *SCGUEHistoryInformation
	SCGActivationStatus                   *SCGActivationStatus
	CPAInformationModReqAck               *CPAInformationModReqAck
	QMCCoordinationResponse               *QMCCoordinationResponse
	SourceSNToTargetSNQMCInfo             *QMCConfigInfo
}

func (msg *SNodeModificationRequestAcknowledge) toIEs() []XnAPMessageIE {
	ies := make([]XnAPMessageIE, 0)
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_MNGRANnodeUEXnAPID)},
		Criticality: Criticality{Value: CriticalityIgnore},
		Value:       &msg.MNGRANnodeUEXnAPID,
	})
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SNGRANnodeUEXnAPID)},
		Criticality: Criticality{Value: CriticalityIgnore},
		Value:       &msg.SNGRANnodeUEXnAPID,
	})
	if msg.PDUSessionAdmittedSNModResponse != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_PDUSessionAdmittedSNModResponse)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.PDUSessionAdmittedSNModResponse,
		})
	}
	if msg.PDUSessionNotAdmittedSNModResponse != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_PDUSessionNotAdmittedSNModResponse)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.PDUSessionNotAdmittedSNModResponse,
		})
	}
	if len(msg.SNToMNContainer) > 0 {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SNToMNContainer)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       &octetStringIE{Value: msg.SNToMNContainer},
		})
	}
	if msg.AdmittedSplitSRB != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_AdmittedSplitSRB)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.AdmittedSplitSRB,
		})
	}
	if msg.AdmittedSplitSRBrelease != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_AdmittedSplitSRBrelease)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.AdmittedSplitSRBrelease,
		})
	}
	if msg.CriticalityDiagnostics != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_CriticalityDiagnostics)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.CriticalityDiagnostics,
		})
	}
	if msg.LocationInformationSN != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_LocationInformationSN)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.LocationInformationSN,
		})
	}
	if msg.MRDCResourceCoordinationInfo != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_MRDCResourceCoordinationInfo)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.MRDCResourceCoordinationInfo,
		})
	}
	if msg.PDUSessionDataForwardingSNModResponse != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_PDUSessionDataForwardingSNModResponse)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.PDUSessionDataForwardingSNModResponse,
		})
	}
	if msg.RRCConfigIndication != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_RRCConfigIndication)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.RRCConfigIndication,
		})
	}
	if msg.AvailableFastMCGRecoveryViaSRB3 != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_AvailableFastMCGRecoveryViaSRB3)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.AvailableFastMCGRecoveryViaSRB3,
		})
	}
	if msg.ReleaseFastMCGRecoveryViaSRB3 != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_ReleaseFastMCGRecoveryViaSRB3)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.ReleaseFastMCGRecoveryViaSRB3,
		})
	}
	if msg.DirectForwardingPathAvailability != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_DirectForwardingPathAvailability)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.DirectForwardingPathAvailability,
		})
	}
	if msg.SCGUEHistoryInformation != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SCGUEHistoryInformation)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.SCGUEHistoryInformation,
		})
	}
	if msg.SCGActivationStatus != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SCGActivationStatus)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.SCGActivationStatus,
		})
	}
	if msg.CPAInformationModReqAck != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_CPAInformationModReqAck)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.CPAInformationModReqAck,
		})
	}
	if msg.QMCCoordinationResponse != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_QMCCoordinationResponse)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.QMCCoordinationResponse,
		})
	}
	if msg.SourceSNToTargetSNQMCInfo != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SourceSNToTargetSNQMCInfo)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.SourceSNToTargetSNQMCInfo,
		})
	}
	return ies
}

func (msg *SNodeModificationRequestAcknowledge) Encode(w io.Writer) error {
	inner := per.NewEncoder(per.APER)
	if err := encodeBody(msg, inner); err != nil {
		return err
	}

	outer := per.NewEncoder(per.APER)

	if err := choiceXnMsg(outer, common.SuccessfulOutcome); err != nil {
		return err
	}
	if err := procedureCodeXnMsg(outer, common.ProcedureCode_MNGRANnodeinitiatedSNGRANnodeModificationPreparation); err != nil {
		return err
	}
	if err := criticalityXnMsg(outer, CriticalityReject); err != nil {
		return err
	}

	if err := outer.EncodeOpenType(inner.Bytes()); err != nil {
		return err
	}

	_, err := w.Write(outer.Bytes())
	return err
}

func (msg *SNodeModificationRequestAcknowledge) Decode(wire []byte) error {
	d := per.NewDecoder(wire, per.APER)

	seq := d.NewSequenceDecoder(per.SequenceConstraints{Extensible: true})
	if err := seq.DecodeExtensionBit(); err != nil {
		return fmt.Errorf("decode SNodeModificationRequestAcknowledge extension bit: %w", err)
	}

	seqOf := d.NewSequenceOfDecoder(per.SizeConstraints{
		Extensible: false,
		Min:        common.Ptr(int64(0)),
		Max:        common.Ptr(int64(MaxProtocolIEs)),
	})
	n, err := seqOf.DecodeLength()
	if err != nil {
		return fmt.Errorf("decode SNodeModificationRequestAcknowledge protocolIEs length: %w", err)
	}

	seen := make(map[int64]bool, n)
	for i := range n {
		id, err := d.DecodeInteger(protocolIEIDConstraints)
		if err != nil {
			return fmt.Errorf("decode SNodeModificationRequestAcknowledge IE[%d] id: %w", i, err)
		}
		seen[id] = true

		var criticality Criticality
		if err := criticality.Decode(d); err != nil {
			return fmt.Errorf("decode SNodeModificationRequestAcknowledge IE[%d] criticality: %w", i, err)
		}

		payload, err := d.DecodeOpenType()
		if err != nil {
			return fmt.Errorf("decode SNodeModificationRequestAcknowledge IE[%d] value: %w", i, err)
		}
		ies := per.NewDecoder(payload, per.APER)

		switch id {
		case common.ProtocolIEID_MNGRANnodeUEXnAPID:
			err = msg.MNGRANnodeUEXnAPID.Decode(ies)
		case common.ProtocolIEID_SNGRANnodeUEXnAPID:
			err = msg.SNGRANnodeUEXnAPID.Decode(ies)
		case common.ProtocolIEID_PDUSessionAdmittedSNModResponse:
			v := &PDUSessionAdmittedSNModResponse{}
			err = v.Decode(ies)
			msg.PDUSessionAdmittedSNModResponse = v
		case common.ProtocolIEID_PDUSessionNotAdmittedSNModResponse:
			v := &PDUSessionNotAdmittedSNModResponse{}
			err = v.Decode(ies)
			msg.PDUSessionNotAdmittedSNModResponse = v
		case common.ProtocolIEID_SNToMNContainer:
			v := octetStringIE{}
			err = v.Decode(ies)
			msg.SNToMNContainer = v.Value
		case common.ProtocolIEID_AdmittedSplitSRB:
			v := &SplitSRBsTypes{}
			err = v.Decode(ies)
			msg.AdmittedSplitSRB = v
		case common.ProtocolIEID_AdmittedSplitSRBrelease:
			v := &SplitSRBsTypes{}
			err = v.Decode(ies)
			msg.AdmittedSplitSRBrelease = v
		case common.ProtocolIEID_CriticalityDiagnostics:
			v := &CriticalityDiagnostics{}
			err = v.Decode(ies)
			msg.CriticalityDiagnostics = v
		case common.ProtocolIEID_LocationInformationSN:
			v := &TargetCGI{}
			err = v.Decode(ies)
			msg.LocationInformationSN = v
		case common.ProtocolIEID_MRDCResourceCoordinationInfo:
			v := &MRDCResourceCoordinationInfo{}
			err = v.Decode(ies)
			msg.MRDCResourceCoordinationInfo = v
		case common.ProtocolIEID_PDUSessionDataForwardingSNModResponse:
			v := &PDUSessionDataForwardingSNModResponse{}
			err = v.Decode(ies)
			msg.PDUSessionDataForwardingSNModResponse = v
		case common.ProtocolIEID_RRCConfigIndication:
			v := &RRCConfigIndication{}
			err = v.Decode(ies)
			msg.RRCConfigIndication = v
		case common.ProtocolIEID_AvailableFastMCGRecoveryViaSRB3:
			v := &AvailableFastMCGRecoveryViaSRB3{}
			err = v.Decode(ies)
			msg.AvailableFastMCGRecoveryViaSRB3 = v
		case common.ProtocolIEID_ReleaseFastMCGRecoveryViaSRB3:
			v := &ReleaseFastMCGRecoveryViaSRB3{}
			err = v.Decode(ies)
			msg.ReleaseFastMCGRecoveryViaSRB3 = v
		case common.ProtocolIEID_DirectForwardingPathAvailability:
			v := &DirectForwardingPathAvailability{}
			err = v.Decode(ies)
			msg.DirectForwardingPathAvailability = v
		case common.ProtocolIEID_SCGUEHistoryInformation:
			v := &SCGUEHistoryInformation{}
			err = v.Decode(ies)
			msg.SCGUEHistoryInformation = v
		case common.ProtocolIEID_SCGActivationStatus:
			v := &SCGActivationStatus{}
			err = v.Decode(ies)
			msg.SCGActivationStatus = v
		case common.ProtocolIEID_CPAInformationModReqAck:
			v := &CPAInformationModReqAck{}
			err = v.Decode(ies)
			msg.CPAInformationModReqAck = v
		case common.ProtocolIEID_QMCCoordinationResponse:
			v := &QMCCoordinationResponse{}
			err = v.Decode(ies)
			msg.QMCCoordinationResponse = v
		case common.ProtocolIEID_SourceSNToTargetSNQMCInfo:
			v := &QMCConfigInfo{}
			err = v.Decode(ies)
			msg.SourceSNToTargetSNQMCInfo = v
		default:
			if criticality.Value == CriticalityReject {
				return fmt.Errorf("unsupported SNodeModificationRequestAcknowledge IE id=%d with reject criticality", id)
			}
		}
		if err != nil {
			return fmt.Errorf("decode SNodeModificationRequestAcknowledge IE id=%d: %w", id, err)
		}
	}

	if _, err := seq.DecodeExtensionAdditions(); err != nil {
		return fmt.Errorf("decode SNodeModificationRequestAcknowledge extensions: %w", err)
	}

	for _, id := range []int64{
		common.ProtocolIEID_MNGRANnodeUEXnAPID,
		common.ProtocolIEID_SNGRANnodeUEXnAPID,
	} {
		if !seen[id] {
			return fmt.Errorf("missing mandatory SNodeModificationRequestAcknowledge IE id=%d", id)
		}
	}

	return nil
}
