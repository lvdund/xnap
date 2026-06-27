package ies

import (
	"fmt"
	"io"

	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

// SNodeAdditionRequestAcknowledge ::= SEQUENCE {
//	protocolIEs  ProtocolIE-Container {{SNodeAdditionRequestAcknowledge-IEs}},
//	...
// }

type SNodeAdditionRequestAcknowledge struct {
	MNGRANnodeUEXnAPID                           NGRANnodeUEXnAPID
	SNGRANnodeUEXnAPID                           NGRANnodeUEXnAPID
	PDUSessionAdmittedAddedAddReqAck             PDUSessionAdmittedAddedAddReqAck
	PDUSessionNotAdmittedAddReqAck               *PDUSessionNotAdmittedAddReqAck
	SNToMNContainer                              []byte
	AdmittedSplitSRB                             *SplitSRBsTypes
	RRCConfigIndication                          *RRCConfigIndication
	CriticalityDiagnostics                       *CriticalityDiagnostics
	LocationInformationSN                        *TargetCGI
	MRDCResourceCoordinationInfo                 *MRDCResourceCoordinationInfo
	AvailableFastMCGRecoveryViaSRB3              *AvailableFastMCGRecoveryViaSRB3
	DirectForwardingPathAvailability             *DirectForwardingPathAvailability
	SCGActivationStatus                          *SCGActivationStatus
	CPAInformationAck                            *CPAInformationAck
	SNMobilityInformation                        *SNMobilityInformation
	QMCCoordinationResponse                      *QMCCoordinationResponse
	CHOinformationAddReqAck                      *CHOinformationAddReqAck
	DirectForwardingPathAvailabilityWithSourceMN *DirectForwardingPathAvailabilityWithSourceMN
}

func (msg *SNodeAdditionRequestAcknowledge) toIEs() []XnAPMessageIE {
	ies := make([]XnAPMessageIE, 0)
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_MNGRANnodeUEXnAPID)},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.MNGRANnodeUEXnAPID,
	})
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SNGRANnodeUEXnAPID)},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.SNGRANnodeUEXnAPID,
	})
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_PDUSessionAdmittedAddedAddReqAck)},
		Criticality: Criticality{Value: CriticalityIgnore},
		Value:       &msg.PDUSessionAdmittedAddedAddReqAck,
	})
	if msg.PDUSessionNotAdmittedAddReqAck != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_PDUSessionNotAdmittedAddReqAck)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.PDUSessionNotAdmittedAddReqAck,
		})
	}
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SNToMNContainer)},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &octetStringIE{Value: msg.SNToMNContainer},
	})
	if msg.AdmittedSplitSRB != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_AdmittedSplitSRB)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.AdmittedSplitSRB,
		})
	}
	if msg.RRCConfigIndication != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_RRCConfigIndication)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.RRCConfigIndication,
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
	if msg.AvailableFastMCGRecoveryViaSRB3 != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_AvailableFastMCGRecoveryViaSRB3)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.AvailableFastMCGRecoveryViaSRB3,
		})
	}
	if msg.DirectForwardingPathAvailability != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_DirectForwardingPathAvailability)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.DirectForwardingPathAvailability,
		})
	}
	if msg.SCGActivationStatus != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SCGActivationStatus)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.SCGActivationStatus,
		})
	}
	if msg.CPAInformationAck != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_CPAInformationAck)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.CPAInformationAck,
		})
	}
	if msg.SNMobilityInformation != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SNMobilityInformation)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.SNMobilityInformation,
		})
	}
	if msg.QMCCoordinationResponse != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_QMCCoordinationResponse)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.QMCCoordinationResponse,
		})
	}
	if msg.CHOinformationAddReqAck != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_CHOinformationAddReqAck)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.CHOinformationAddReqAck,
		})
	}
	if msg.DirectForwardingPathAvailabilityWithSourceMN != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_DirectForwardingPathAvailabilityWithSourceMN)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.DirectForwardingPathAvailabilityWithSourceMN,
		})
	}
	return ies
}

func (msg *SNodeAdditionRequestAcknowledge) Encode(w io.Writer) error {
	inner := per.NewEncoder(per.APER)
	if err := encodeBody(msg, inner); err != nil {
		return err
	}

	outer := per.NewEncoder(per.APER)

	if err := choiceXnMsg(outer, common.SuccessfulOutcome); err != nil {
		return err
	}
	if err := procedureCodeXnMsg(outer, common.ProcedureCode_SNGRANnodeAdditionPreparation); err != nil {
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

func (msg *SNodeAdditionRequestAcknowledge) Decode(wire []byte) error {
	d := per.NewDecoder(wire, per.APER)

	seq := d.NewSequenceDecoder(per.SequenceConstraints{Extensible: true})
	if err := seq.DecodeExtensionBit(); err != nil {
		return fmt.Errorf("decode SNodeAdditionRequestAcknowledge extension bit: %w", err)
	}

	seqOf := d.NewSequenceOfDecoder(per.SizeConstraints{
		Extensible: false,
		Min:        common.Ptr(int64(0)),
		Max:        common.Ptr(int64(MaxProtocolIEs)),
	})
	n, err := seqOf.DecodeLength()
	if err != nil {
		return fmt.Errorf("decode SNodeAdditionRequestAcknowledge protocolIEs length: %w", err)
	}

	seen := make(map[int64]bool, n)
	for i := range n {
		id, err := d.DecodeInteger(protocolIEIDConstraints)
		if err != nil {
			return fmt.Errorf("decode SNodeAdditionRequestAcknowledge IE[%d] id: %w", i, err)
		}
		seen[id] = true

		var criticality Criticality
		if err := criticality.Decode(d); err != nil {
			return fmt.Errorf("decode SNodeAdditionRequestAcknowledge IE[%d] criticality: %w", i, err)
		}

		payload, err := d.DecodeOpenType()
		if err != nil {
			return fmt.Errorf("decode SNodeAdditionRequestAcknowledge IE[%d] value: %w", i, err)
		}
		ies := per.NewDecoder(payload, per.APER)

		switch id {
		case common.ProtocolIEID_MNGRANnodeUEXnAPID:
			err = msg.MNGRANnodeUEXnAPID.Decode(ies)
		case common.ProtocolIEID_SNGRANnodeUEXnAPID:
			err = msg.SNGRANnodeUEXnAPID.Decode(ies)
		case common.ProtocolIEID_PDUSessionAdmittedAddedAddReqAck:
			err = msg.PDUSessionAdmittedAddedAddReqAck.Decode(ies)
		case common.ProtocolIEID_PDUSessionNotAdmittedAddReqAck:
			v := &PDUSessionNotAdmittedAddReqAck{}
			err = v.Decode(ies)
			msg.PDUSessionNotAdmittedAddReqAck = v
		case common.ProtocolIEID_SNToMNContainer:
			v := octetStringIE{}
			err = v.Decode(ies)
			msg.SNToMNContainer = v.Value
		case common.ProtocolIEID_AdmittedSplitSRB:
			v := &SplitSRBsTypes{}
			err = v.Decode(ies)
			msg.AdmittedSplitSRB = v
		case common.ProtocolIEID_RRCConfigIndication:
			v := &RRCConfigIndication{}
			err = v.Decode(ies)
			msg.RRCConfigIndication = v
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
		case common.ProtocolIEID_AvailableFastMCGRecoveryViaSRB3:
			v := &AvailableFastMCGRecoveryViaSRB3{}
			err = v.Decode(ies)
			msg.AvailableFastMCGRecoveryViaSRB3 = v
		case common.ProtocolIEID_DirectForwardingPathAvailability:
			v := &DirectForwardingPathAvailability{}
			err = v.Decode(ies)
			msg.DirectForwardingPathAvailability = v
		case common.ProtocolIEID_SCGActivationStatus:
			v := &SCGActivationStatus{}
			err = v.Decode(ies)
			msg.SCGActivationStatus = v
		case common.ProtocolIEID_CPAInformationAck:
			v := &CPAInformationAck{}
			err = v.Decode(ies)
			msg.CPAInformationAck = v
		case common.ProtocolIEID_SNMobilityInformation:
			v := &SNMobilityInformation{}
			err = v.Decode(ies)
			msg.SNMobilityInformation = v
		case common.ProtocolIEID_QMCCoordinationResponse:
			v := &QMCCoordinationResponse{}
			err = v.Decode(ies)
			msg.QMCCoordinationResponse = v
		case common.ProtocolIEID_CHOinformationAddReqAck:
			v := &CHOinformationAddReqAck{}
			err = v.Decode(ies)
			msg.CHOinformationAddReqAck = v
		case common.ProtocolIEID_DirectForwardingPathAvailabilityWithSourceMN:
			v := &DirectForwardingPathAvailabilityWithSourceMN{}
			err = v.Decode(ies)
			msg.DirectForwardingPathAvailabilityWithSourceMN = v
		default:
			if criticality.Value == CriticalityReject {
				return fmt.Errorf("unsupported SNodeAdditionRequestAcknowledge IE id=%d with reject criticality", id)
			}
		}
		if err != nil {
			return fmt.Errorf("decode SNodeAdditionRequestAcknowledge IE id=%d: %w", id, err)
		}
	}

	if _, err := seq.DecodeExtensionAdditions(); err != nil {
		return fmt.Errorf("decode SNodeAdditionRequestAcknowledge extensions: %w", err)
	}

	for _, id := range []int64{
		common.ProtocolIEID_MNGRANnodeUEXnAPID,
		common.ProtocolIEID_SNGRANnodeUEXnAPID,
		common.ProtocolIEID_PDUSessionAdmittedAddedAddReqAck,
		common.ProtocolIEID_SNToMNContainer,
	} {
		if !seen[id] {
			return fmt.Errorf("missing mandatory SNodeAdditionRequestAcknowledge IE id=%d", id)
		}
	}

	return nil
}
