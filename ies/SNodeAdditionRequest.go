package ies

import (
	"fmt"
	"io"

	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

// SNodeAdditionRequest ::= SEQUENCE {
//	protocolIEs  ProtocolIE-Container {{SNodeAdditionRequest-IEs}},
//	...
// }

type SNodeAdditionRequest struct {
	MNGRANnodeUEXnAPID               NGRANnodeUEXnAPID
	UESecurityCapabilities           UESecurityCapabilities
	SNgRANnodeSecurityKey            SNGRANnodeSecurityKey
	SNGRANnodeUEAMBR                 UEAggregateMaximumBitRate
	SelectedPLMN                     *PLMNIdentity
	MobilityRestrictionList          *MobilityRestrictionList
	IndexToRatFrequSelectionPriority *RFSPIndex
	PDUSessionToBeAddedAddReq        PDUSessionToBeAddedAddReq
	MNToSNContainer                  []byte
	SNGRANnodeUEXnAPID               *NGRANnodeUEXnAPID
	ExpectedUEBehaviour              *ExpectedUEBehaviour
	RequestedSplitSRB                *SplitSRBsTypes
	PCellID                          *GlobalNGRANCellID
	DesiredActNotificationLevel      *DesiredActNotificationLevel
	AvailableDRBIDs                  *DRBList
	SNGRANnodeMaxIPDataRateUL        *BitRate
	SNGRANnodeMaxIPDataRateDL        *BitRate
	LocationInformationSNReporting   *LocationInformationSNReporting
	MRDCResourceCoordinationInfo     *MRDCResourceCoordinationInfo
	MaskedIMEISV                     *MaskedIMEISV
	NEDCTDMPattern                   *NEDCTDMPattern
	SNGRANnodeAdditionTriggerInd     *SNGRANnodeAdditionTriggerInd
	TraceActivation                  *TraceActivation
	RequestedFastMCGRecoveryViaSRB3  *RequestedFastMCGRecoveryViaSRB3
	UERadioCapabilityID              *UERadioCapabilityID
	SourceNGRANNodeID                *GlobalNGRANNodeID
	ManagementBasedMDTPLMNList       *MDTPLMNList
	UEHistoryInformation             *UEHistoryInformation
	UEHistoryInformationFromTheUE    *UEHistoryInformationFromTheUE
	PSCellChangeHistory              *PSCellChangeHistory
	IABNodeIndication                *IABNodeIndication
	NoPDUSessionIndication           *NoPDUSessionIndication
	CHOinformationAddReq             *CHOinformationAddReq
	SCGActivationRequest             *SCGActivationRequest
	CPAInformationRequest            *CPAInformationRequest
	SNGRANnodeUESliceMBR             *UESliceMaximumBitRateList
	F1TerminatingIABDonorIndicator   *F1TerminatingIABDonorIndicator
	SelectedNID                      *NID
	QMCCoordinationRequest           *QMCCoordinationRequest
	SourceSNToTargetSNQMCInfo        *QMCConfigInfo
	IABAuthorizationStatus           *IABAuthorizationStatus
	SourceMNGRANnodeID               *GlobalNGRANNodeID
}

func (msg *SNodeAdditionRequest) toIEs() []XnAPMessageIE {
	ies := make([]XnAPMessageIE, 0)
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_MNGRANnodeUEXnAPID)},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.MNGRANnodeUEXnAPID,
	})
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_UESecurityCapabilities)},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.UESecurityCapabilities,
	})
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SNgRANnodeSecurityKey)},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.SNgRANnodeSecurityKey,
	})
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SNGRANnodeUEAMBR)},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.SNGRANnodeUEAMBR,
	})
	if msg.SelectedPLMN != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SelectedPLMN)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.SelectedPLMN,
		})
	}
	if msg.MobilityRestrictionList != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_MobilityRestrictionList)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.MobilityRestrictionList,
		})
	}
	if msg.IndexToRatFrequSelectionPriority != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_IndexToRatFrequSelectionPriority)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.IndexToRatFrequSelectionPriority,
		})
	}
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_PDUSessionToBeAddedAddReq)},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.PDUSessionToBeAddedAddReq,
	})
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_MNToSNContainer)},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &octetStringIE{Value: msg.MNToSNContainer},
	})
	if msg.SNGRANnodeUEXnAPID != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SNGRANnodeUEXnAPID)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.SNGRANnodeUEXnAPID,
		})
	}
	if msg.ExpectedUEBehaviour != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_ExpectedUEBehaviour)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.ExpectedUEBehaviour,
		})
	}
	if msg.RequestedSplitSRB != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_RequestedSplitSRB)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.RequestedSplitSRB,
		})
	}
	if msg.PCellID != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_PCellID)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.PCellID,
		})
	}
	if msg.DesiredActNotificationLevel != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_DesiredActNotificationLevel)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.DesiredActNotificationLevel,
		})
	}
	if msg.AvailableDRBIDs != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_AvailableDRBIDs)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.AvailableDRBIDs,
		})
	}
	if msg.SNGRANnodeMaxIPDataRateUL != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SNGRANnodeMaxIPDataRateUL)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.SNGRANnodeMaxIPDataRateUL,
		})
	}
	if msg.SNGRANnodeMaxIPDataRateDL != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SNGRANnodeMaxIPDataRateDL)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.SNGRANnodeMaxIPDataRateDL,
		})
	}
	if msg.LocationInformationSNReporting != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_LocationInformationSNReporting)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.LocationInformationSNReporting,
		})
	}
	if msg.MRDCResourceCoordinationInfo != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_MRDCResourceCoordinationInfo)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.MRDCResourceCoordinationInfo,
		})
	}
	if msg.MaskedIMEISV != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_MaskedIMEISV)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.MaskedIMEISV,
		})
	}
	if msg.NEDCTDMPattern != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_NEDCTDMPattern)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.NEDCTDMPattern,
		})
	}
	if msg.SNGRANnodeAdditionTriggerInd != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SNGRANnodeAdditionTriggerInd)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.SNGRANnodeAdditionTriggerInd,
		})
	}
	if msg.TraceActivation != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_TraceActivation)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.TraceActivation,
		})
	}
	if msg.RequestedFastMCGRecoveryViaSRB3 != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_RequestedFastMCGRecoveryViaSRB3)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.RequestedFastMCGRecoveryViaSRB3,
		})
	}
	if msg.UERadioCapabilityID != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_UERadioCapabilityID)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.UERadioCapabilityID,
		})
	}
	if msg.SourceNGRANNodeID != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SourceNGRANNodeID)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.SourceNGRANNodeID,
		})
	}
	if msg.ManagementBasedMDTPLMNList != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_ManagementBasedMDTPLMNList)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.ManagementBasedMDTPLMNList,
		})
	}
	if msg.UEHistoryInformation != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_UEHistoryInformation)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.UEHistoryInformation,
		})
	}
	if msg.UEHistoryInformationFromTheUE != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_UEHistoryInformationFromTheUE)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.UEHistoryInformationFromTheUE,
		})
	}
	if msg.PSCellChangeHistory != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_PSCellChangeHistory)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.PSCellChangeHistory,
		})
	}
	if msg.IABNodeIndication != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_IABNodeIndication)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.IABNodeIndication,
		})
	}
	if msg.NoPDUSessionIndication != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_NoPDUSessionIndication)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.NoPDUSessionIndication,
		})
	}
	if msg.CHOinformationAddReq != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_CHOinformationAddReq)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.CHOinformationAddReq,
		})
	}
	if msg.SCGActivationRequest != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SCGActivationRequest)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.SCGActivationRequest,
		})
	}
	if msg.CPAInformationRequest != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_CPAInformationRequest)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.CPAInformationRequest,
		})
	}
	if msg.SNGRANnodeUESliceMBR != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SNGRANnodeUESliceMBR)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.SNGRANnodeUESliceMBR,
		})
	}
	if msg.F1TerminatingIABDonorIndicator != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_F1TerminatingIABDonorIndicator)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.F1TerminatingIABDonorIndicator,
		})
	}
	if msg.SelectedNID != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SelectedNID)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.SelectedNID,
		})
	}
	if msg.QMCCoordinationRequest != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_QMCCoordinationRequest)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.QMCCoordinationRequest,
		})
	}
	if msg.SourceSNToTargetSNQMCInfo != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SourceSNToTargetSNQMCInfo)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.SourceSNToTargetSNQMCInfo,
		})
	}
	if msg.IABAuthorizationStatus != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_IABAuthorizationStatus)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.IABAuthorizationStatus,
		})
	}
	if msg.SourceMNGRANnodeID != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SourceMNGRANnodeID)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.SourceMNGRANnodeID,
		})
	}
	return ies
}

func (msg *SNodeAdditionRequest) Encode(w io.Writer) error {
	inner := per.NewEncoder(per.APER)
	if err := encodeBody(msg, inner); err != nil {
		return err
	}

	outer := per.NewEncoder(per.APER)

	if err := choiceXnMsg(outer, common.InitiatingMessage); err != nil {
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

func (msg *SNodeAdditionRequest) Decode(wire []byte) error {
	d := per.NewDecoder(wire, per.APER)

	seq := d.NewSequenceDecoder(per.SequenceConstraints{Extensible: true})
	if err := seq.DecodeExtensionBit(); err != nil {
		return fmt.Errorf("decode SNodeAdditionRequest extension bit: %w", err)
	}

	seqOf := d.NewSequenceOfDecoder(per.SizeConstraints{
		Extensible: false,
		Min:        common.Ptr(int64(0)),
		Max:        common.Ptr(int64(MaxProtocolIEs)),
	})
	n, err := seqOf.DecodeLength()
	if err != nil {
		return fmt.Errorf("decode SNodeAdditionRequest protocolIEs length: %w", err)
	}

	seen := make(map[int64]bool, n)
	for i := range n {
		id, err := d.DecodeInteger(protocolIEIDConstraints)
		if err != nil {
			return fmt.Errorf("decode SNodeAdditionRequest IE[%d] id: %w", i, err)
		}
		seen[id] = true

		var criticality Criticality
		if err := criticality.Decode(d); err != nil {
			return fmt.Errorf("decode SNodeAdditionRequest IE[%d] criticality: %w", i, err)
		}

		payload, err := d.DecodeOpenType()
		if err != nil {
			return fmt.Errorf("decode SNodeAdditionRequest IE[%d] value: %w", i, err)
		}
		ies := per.NewDecoder(payload, per.APER)

		switch id {
		case common.ProtocolIEID_MNGRANnodeUEXnAPID:
			err = msg.MNGRANnodeUEXnAPID.Decode(ies)
		case common.ProtocolIEID_UESecurityCapabilities:
			err = msg.UESecurityCapabilities.Decode(ies)
		case common.ProtocolIEID_SNgRANnodeSecurityKey:
			err = msg.SNgRANnodeSecurityKey.Decode(ies)
		case common.ProtocolIEID_SNGRANnodeUEAMBR:
			err = msg.SNGRANnodeUEAMBR.Decode(ies)
		case common.ProtocolIEID_SelectedPLMN:
			v := &PLMNIdentity{}
			err = v.Decode(ies)
			msg.SelectedPLMN = v
		case common.ProtocolIEID_MobilityRestrictionList:
			v := &MobilityRestrictionList{}
			err = v.Decode(ies)
			msg.MobilityRestrictionList = v
		case common.ProtocolIEID_IndexToRatFrequSelectionPriority:
			v := &RFSPIndex{}
			err = v.Decode(ies)
			msg.IndexToRatFrequSelectionPriority = v
		case common.ProtocolIEID_PDUSessionToBeAddedAddReq:
			err = msg.PDUSessionToBeAddedAddReq.Decode(ies)
		case common.ProtocolIEID_MNToSNContainer:
			v := octetStringIE{}
			err = v.Decode(ies)
			msg.MNToSNContainer = v.Value
		case common.ProtocolIEID_SNGRANnodeUEXnAPID:
			v := &NGRANnodeUEXnAPID{}
			err = v.Decode(ies)
			msg.SNGRANnodeUEXnAPID = v
		case common.ProtocolIEID_ExpectedUEBehaviour:
			v := &ExpectedUEBehaviour{}
			err = v.Decode(ies)
			msg.ExpectedUEBehaviour = v
		case common.ProtocolIEID_RequestedSplitSRB:
			v := &SplitSRBsTypes{}
			err = v.Decode(ies)
			msg.RequestedSplitSRB = v
		case common.ProtocolIEID_PCellID:
			v := &GlobalNGRANCellID{}
			err = v.Decode(ies)
			msg.PCellID = v
		case common.ProtocolIEID_DesiredActNotificationLevel:
			v := &DesiredActNotificationLevel{}
			err = v.Decode(ies)
			msg.DesiredActNotificationLevel = v
		case common.ProtocolIEID_AvailableDRBIDs:
			v := &DRBList{}
			err = v.Decode(ies)
			msg.AvailableDRBIDs = v
		case common.ProtocolIEID_SNGRANnodeMaxIPDataRateUL:
			v := &BitRate{}
			err = v.Decode(ies)
			msg.SNGRANnodeMaxIPDataRateUL = v
		case common.ProtocolIEID_SNGRANnodeMaxIPDataRateDL:
			v := &BitRate{}
			err = v.Decode(ies)
			msg.SNGRANnodeMaxIPDataRateDL = v
		case common.ProtocolIEID_LocationInformationSNReporting:
			v := &LocationInformationSNReporting{}
			err = v.Decode(ies)
			msg.LocationInformationSNReporting = v
		case common.ProtocolIEID_MRDCResourceCoordinationInfo:
			v := &MRDCResourceCoordinationInfo{}
			err = v.Decode(ies)
			msg.MRDCResourceCoordinationInfo = v
		case common.ProtocolIEID_MaskedIMEISV:
			v := &MaskedIMEISV{}
			err = v.Decode(ies)
			msg.MaskedIMEISV = v
		case common.ProtocolIEID_NEDCTDMPattern:
			v := &NEDCTDMPattern{}
			err = v.Decode(ies)
			msg.NEDCTDMPattern = v
		case common.ProtocolIEID_SNGRANnodeAdditionTriggerInd:
			v := &SNGRANnodeAdditionTriggerInd{}
			err = v.Decode(ies)
			msg.SNGRANnodeAdditionTriggerInd = v
		case common.ProtocolIEID_TraceActivation:
			v := &TraceActivation{}
			err = v.Decode(ies)
			msg.TraceActivation = v
		case common.ProtocolIEID_RequestedFastMCGRecoveryViaSRB3:
			v := &RequestedFastMCGRecoveryViaSRB3{}
			err = v.Decode(ies)
			msg.RequestedFastMCGRecoveryViaSRB3 = v
		case common.ProtocolIEID_UERadioCapabilityID:
			v := &UERadioCapabilityID{}
			err = v.Decode(ies)
			msg.UERadioCapabilityID = v
		case common.ProtocolIEID_SourceNGRANNodeID:
			v := &GlobalNGRANNodeID{}
			err = v.Decode(ies)
			msg.SourceNGRANNodeID = v
		case common.ProtocolIEID_ManagementBasedMDTPLMNList:
			v := &MDTPLMNList{}
			err = v.Decode(ies)
			msg.ManagementBasedMDTPLMNList = v
		case common.ProtocolIEID_UEHistoryInformation:
			v := &UEHistoryInformation{}
			err = v.Decode(ies)
			msg.UEHistoryInformation = v
		case common.ProtocolIEID_UEHistoryInformationFromTheUE:
			v := &UEHistoryInformationFromTheUE{}
			err = v.Decode(ies)
			msg.UEHistoryInformationFromTheUE = v
		case common.ProtocolIEID_PSCellChangeHistory:
			v := &PSCellChangeHistory{}
			err = v.Decode(ies)
			msg.PSCellChangeHistory = v
		case common.ProtocolIEID_IABNodeIndication:
			v := &IABNodeIndication{}
			err = v.Decode(ies)
			msg.IABNodeIndication = v
		case common.ProtocolIEID_NoPDUSessionIndication:
			v := &NoPDUSessionIndication{}
			err = v.Decode(ies)
			msg.NoPDUSessionIndication = v
		case common.ProtocolIEID_CHOinformationAddReq:
			v := &CHOinformationAddReq{}
			err = v.Decode(ies)
			msg.CHOinformationAddReq = v
		case common.ProtocolIEID_SCGActivationRequest:
			v := &SCGActivationRequest{}
			err = v.Decode(ies)
			msg.SCGActivationRequest = v
		case common.ProtocolIEID_CPAInformationRequest:
			v := &CPAInformationRequest{}
			err = v.Decode(ies)
			msg.CPAInformationRequest = v
		case common.ProtocolIEID_SNGRANnodeUESliceMBR:
			v := &UESliceMaximumBitRateList{}
			err = v.Decode(ies)
			msg.SNGRANnodeUESliceMBR = v
		case common.ProtocolIEID_F1TerminatingIABDonorIndicator:
			v := &F1TerminatingIABDonorIndicator{}
			err = v.Decode(ies)
			msg.F1TerminatingIABDonorIndicator = v
		case common.ProtocolIEID_SelectedNID:
			v := &NID{}
			err = v.Decode(ies)
			msg.SelectedNID = v
		case common.ProtocolIEID_QMCCoordinationRequest:
			v := &QMCCoordinationRequest{}
			err = v.Decode(ies)
			msg.QMCCoordinationRequest = v
		case common.ProtocolIEID_SourceSNToTargetSNQMCInfo:
			v := &QMCConfigInfo{}
			err = v.Decode(ies)
			msg.SourceSNToTargetSNQMCInfo = v
		case common.ProtocolIEID_IABAuthorizationStatus:
			v := &IABAuthorizationStatus{}
			err = v.Decode(ies)
			msg.IABAuthorizationStatus = v
		case common.ProtocolIEID_SourceMNGRANnodeID:
			v := &GlobalNGRANNodeID{}
			err = v.Decode(ies)
			msg.SourceMNGRANnodeID = v
		default:
			if criticality.Value == CriticalityReject {
				return fmt.Errorf("unsupported SNodeAdditionRequest IE id=%d with reject criticality", id)
			}
		}
		if err != nil {
			return fmt.Errorf("decode SNodeAdditionRequest IE id=%d: %w", id, err)
		}
	}

	if _, err := seq.DecodeExtensionAdditions(); err != nil {
		return fmt.Errorf("decode SNodeAdditionRequest extensions: %w", err)
	}

	for _, id := range []int64{
		common.ProtocolIEID_MNGRANnodeUEXnAPID,
		common.ProtocolIEID_UESecurityCapabilities,
		common.ProtocolIEID_SNgRANnodeSecurityKey,
		common.ProtocolIEID_SNGRANnodeUEAMBR,
		common.ProtocolIEID_PDUSessionToBeAddedAddReq,
		common.ProtocolIEID_MNToSNContainer,
	} {
		if !seen[id] {
			return fmt.Errorf("missing mandatory SNodeAdditionRequest IE id=%d", id)
		}
	}

	return nil
}
