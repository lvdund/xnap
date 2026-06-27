package ies

import (
	"fmt"
	"io"

	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

// SNodeModificationRequired ::= SEQUENCE {
//	protocolIEs  ProtocolIE-Container {{SNodeModificationRequired-IEs}},
//	...
// }

type SNodeModificationRequired struct {
	MNGRANnodeUEXnAPID                  NGRANnodeUEXnAPID
	SNGRANnodeUEXnAPID                  NGRANnodeUEXnAPID
	Cause                               Cause
	PDCPChangeIndication                *PDCPChangeIndication
	PDUSessionToBeModifiedSNModRequired *PDUSessionToBeModifiedSNModRequired
	PDUSessionToBeReleasedSNModRequired *PDUSessionToBeReleasedSNModRequired
	SNToMNContainer                     []byte
	SpareDRBIDs                         *DRBList
	RequiredNumberOfDRBIDs              *DRBNumber
	LocationInformationSN               *TargetCGI
	MRDCResourceCoordinationInfo        *MRDCResourceCoordinationInfo
	RRCConfigIndication                 *RRCConfigIndication
	AvailableFastMCGRecoveryViaSRB3     *AvailableFastMCGRecoveryViaSRB3
	ReleaseFastMCGRecoveryViaSRB3       *ReleaseFastMCGRecoveryViaSRB3
	SCGIndicator                        *SCGIndicator
	SCGUEHistoryInformation             *SCGUEHistoryInformation
	SCGActivationRequest                *SCGActivationRequest
	CPACInformationModRequired          *CPACInformationModRequired
	SCGreconfigNotification             *SCGreconfigNotification
	SPRAvailability                     *SPRAvailability
	QMCCoordinationRequest              *QMCCoordinationRequest
	SCPACRequest                        *SCPACRequest
	PDUSessionsListToBeReleasedUPError  *PDUSessionsListToBeReleasedUPError
}

func (msg *SNodeModificationRequired) toIEs() []XnAPMessageIE {
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
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_Cause)},
		Criticality: Criticality{Value: CriticalityIgnore},
		Value:       &msg.Cause,
	})
	if msg.PDCPChangeIndication != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_PDCPChangeIndication)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.PDCPChangeIndication,
		})
	}
	if msg.PDUSessionToBeModifiedSNModRequired != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_PDUSessionToBeModifiedSNModRequired)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.PDUSessionToBeModifiedSNModRequired,
		})
	}
	if msg.PDUSessionToBeReleasedSNModRequired != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_PDUSessionToBeReleasedSNModRequired)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.PDUSessionToBeReleasedSNModRequired,
		})
	}
	if len(msg.SNToMNContainer) > 0 {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SNToMNContainer)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       &octetStringIE{Value: msg.SNToMNContainer},
		})
	}
	if msg.SpareDRBIDs != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SpareDRBIDs)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.SpareDRBIDs,
		})
	}
	if msg.RequiredNumberOfDRBIDs != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_RequiredNumberOfDRBIDs)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.RequiredNumberOfDRBIDs,
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
	if msg.SCGIndicator != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SCGIndicator)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.SCGIndicator,
		})
	}
	if msg.SCGUEHistoryInformation != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SCGUEHistoryInformation)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.SCGUEHistoryInformation,
		})
	}
	if msg.SCGActivationRequest != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SCGActivationRequest)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.SCGActivationRequest,
		})
	}
	if msg.CPACInformationModRequired != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_CPACInformationModRequired)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.CPACInformationModRequired,
		})
	}
	if msg.SCGreconfigNotification != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SCGreconfigNotification)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.SCGreconfigNotification,
		})
	}
	if msg.SPRAvailability != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SPRAvailability)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.SPRAvailability,
		})
	}
	if msg.QMCCoordinationRequest != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_QMCCoordinationRequest)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.QMCCoordinationRequest,
		})
	}
	if msg.SCPACRequest != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SCPACRequest)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.SCPACRequest,
		})
	}
	if msg.PDUSessionsListToBeReleasedUPError != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_PDUSessionsListToBeReleasedUPError)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.PDUSessionsListToBeReleasedUPError,
		})
	}
	return ies
}

func (msg *SNodeModificationRequired) Encode(w io.Writer) error {
	inner := per.NewEncoder(per.APER)
	if err := encodeBody(msg, inner); err != nil {
		return err
	}

	outer := per.NewEncoder(per.APER)

	if err := choiceXnMsg(outer, common.InitiatingMessage); err != nil {
		return err
	}
	if err := procedureCodeXnMsg(outer, common.ProcedureCode_SNGRANnodeinitiatedSNGRANnodeModificationPreparation); err != nil {
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

func (msg *SNodeModificationRequired) Decode(wire []byte) error {
	d := per.NewDecoder(wire, per.APER)

	seq := d.NewSequenceDecoder(per.SequenceConstraints{Extensible: true})
	if err := seq.DecodeExtensionBit(); err != nil {
		return fmt.Errorf("decode SNodeModificationRequired extension bit: %w", err)
	}

	seqOf := d.NewSequenceOfDecoder(per.SizeConstraints{
		Extensible: false,
		Min:        common.Ptr(int64(0)),
		Max:        common.Ptr(int64(MaxProtocolIEs)),
	})
	n, err := seqOf.DecodeLength()
	if err != nil {
		return fmt.Errorf("decode SNodeModificationRequired protocolIEs length: %w", err)
	}

	seen := make(map[int64]bool, n)
	for i := range n {
		id, err := d.DecodeInteger(protocolIEIDConstraints)
		if err != nil {
			return fmt.Errorf("decode SNodeModificationRequired IE[%d] id: %w", i, err)
		}
		seen[id] = true

		var criticality Criticality
		if err := criticality.Decode(d); err != nil {
			return fmt.Errorf("decode SNodeModificationRequired IE[%d] criticality: %w", i, err)
		}

		payload, err := d.DecodeOpenType()
		if err != nil {
			return fmt.Errorf("decode SNodeModificationRequired IE[%d] value: %w", i, err)
		}
		ies := per.NewDecoder(payload, per.APER)

		switch id {
		case common.ProtocolIEID_MNGRANnodeUEXnAPID:
			err = msg.MNGRANnodeUEXnAPID.Decode(ies)
		case common.ProtocolIEID_SNGRANnodeUEXnAPID:
			err = msg.SNGRANnodeUEXnAPID.Decode(ies)
		case common.ProtocolIEID_Cause:
			err = msg.Cause.Decode(ies)
		case common.ProtocolIEID_PDCPChangeIndication:
			v := &PDCPChangeIndication{}
			err = v.Decode(ies)
			msg.PDCPChangeIndication = v
		case common.ProtocolIEID_PDUSessionToBeModifiedSNModRequired:
			v := &PDUSessionToBeModifiedSNModRequired{}
			err = v.Decode(ies)
			msg.PDUSessionToBeModifiedSNModRequired = v
		case common.ProtocolIEID_PDUSessionToBeReleasedSNModRequired:
			v := &PDUSessionToBeReleasedSNModRequired{}
			err = v.Decode(ies)
			msg.PDUSessionToBeReleasedSNModRequired = v
		case common.ProtocolIEID_SNToMNContainer:
			v := octetStringIE{}
			err = v.Decode(ies)
			msg.SNToMNContainer = v.Value
		case common.ProtocolIEID_SpareDRBIDs:
			v := &DRBList{}
			err = v.Decode(ies)
			msg.SpareDRBIDs = v
		case common.ProtocolIEID_RequiredNumberOfDRBIDs:
			v := &DRBNumber{}
			err = v.Decode(ies)
			msg.RequiredNumberOfDRBIDs = v
		case common.ProtocolIEID_LocationInformationSN:
			v := &TargetCGI{}
			err = v.Decode(ies)
			msg.LocationInformationSN = v
		case common.ProtocolIEID_MRDCResourceCoordinationInfo:
			v := &MRDCResourceCoordinationInfo{}
			err = v.Decode(ies)
			msg.MRDCResourceCoordinationInfo = v
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
		case common.ProtocolIEID_SCGIndicator:
			v := &SCGIndicator{}
			err = v.Decode(ies)
			msg.SCGIndicator = v
		case common.ProtocolIEID_SCGUEHistoryInformation:
			v := &SCGUEHistoryInformation{}
			err = v.Decode(ies)
			msg.SCGUEHistoryInformation = v
		case common.ProtocolIEID_SCGActivationRequest:
			v := &SCGActivationRequest{}
			err = v.Decode(ies)
			msg.SCGActivationRequest = v
		case common.ProtocolIEID_CPACInformationModRequired:
			v := &CPACInformationModRequired{}
			err = v.Decode(ies)
			msg.CPACInformationModRequired = v
		case common.ProtocolIEID_SCGreconfigNotification:
			v := &SCGreconfigNotification{}
			err = v.Decode(ies)
			msg.SCGreconfigNotification = v
		case common.ProtocolIEID_SPRAvailability:
			v := &SPRAvailability{}
			err = v.Decode(ies)
			msg.SPRAvailability = v
		case common.ProtocolIEID_QMCCoordinationRequest:
			v := &QMCCoordinationRequest{}
			err = v.Decode(ies)
			msg.QMCCoordinationRequest = v
		case common.ProtocolIEID_SCPACRequest:
			v := &SCPACRequest{}
			err = v.Decode(ies)
			msg.SCPACRequest = v
		case common.ProtocolIEID_PDUSessionsListToBeReleasedUPError:
			v := &PDUSessionsListToBeReleasedUPError{}
			err = v.Decode(ies)
			msg.PDUSessionsListToBeReleasedUPError = v
		default:
			if criticality.Value == CriticalityReject {
				return fmt.Errorf("unsupported SNodeModificationRequired IE id=%d with reject criticality", id)
			}
		}
		if err != nil {
			return fmt.Errorf("decode SNodeModificationRequired IE id=%d: %w", id, err)
		}
	}

	if _, err := seq.DecodeExtensionAdditions(); err != nil {
		return fmt.Errorf("decode SNodeModificationRequired extensions: %w", err)
	}

	for _, id := range []int64{
		common.ProtocolIEID_MNGRANnodeUEXnAPID,
		common.ProtocolIEID_SNGRANnodeUEXnAPID,
		common.ProtocolIEID_Cause,
	} {
		if !seen[id] {
			return fmt.Errorf("missing mandatory SNodeModificationRequired IE id=%d", id)
		}
	}

	return nil
}
