package ies

import (
	"fmt"
	"io"

	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

// SNodeModificationRequest ::= SEQUENCE {
//	protocolIEs  ProtocolIE-Container {{SNodeModificationRequest-IEs}},
//	...
// }

type SNodeModificationRequest struct {
	MNGRANnodeUEXnAPID                     NGRANnodeUEXnAPID
	SNGRANnodeUEXnAPID                     NGRANnodeUEXnAPID
	Cause                                  Cause
	PDCPChangeIndication                   *PDCPChangeIndication
	SelectedPLMN                           *PLMNIdentity
	MobilityRestrictionList                *MobilityRestrictionList
	SCGConfigurationQuery                  *SCGConfigurationQuery
	UEContextInfoSNModRequest              *UEContextInfoSNModRequest
	MNToSNContainer                        []byte
	RequestedSplitSRB                      *SplitSRBsTypes
	RequestedSplitSRBrelease               *SplitSRBsTypes
	DesiredActNotificationLevel            *DesiredActNotificationLevel
	AdditionalDRBIDs                       *DRBList
	SNGRANnodeMaxIPDataRateUL              *BitRate
	SNGRANnodeMaxIPDataRateDL              *BitRate
	LocationInformationSNReporting         *LocationInformationSNReporting
	MRDCResourceCoordinationInfo           *MRDCResourceCoordinationInfo
	PCellID                                *GlobalNGRANCellID
	NEDCTDMPattern                         *NEDCTDMPattern
	RequestedFastMCGRecoveryViaSRB3        *RequestedFastMCGRecoveryViaSRB3
	RequestedFastMCGRecoveryViaSRB3Release *RequestedFastMCGRecoveryViaSRB3Release
	SNTriggered                            *SNTriggered
	TargetNodeID                           *GlobalNGRANNodeID
	PSCellHistoryInformationRetrieve       *PSCellHistoryInformationRetrieve
	UEHistoryInformationFromTheUE          *UEHistoryInformationFromTheUE
	CHOinformationModReq                   *CHOinformationModReq
	SCGActivationRequest                   *SCGActivationRequest
	CPAInformationModReq                   *CPAInformationModReq
	CPCInformationUpdate                   *CPCInformationUpdate
	SNGRANnodeUESliceMBR                   *UESliceMaximumBitRateList
	ManagementBasedMDTPLMNModificationList *MDTPLMNModificationList
	SelectedNID                            *NID
	QMCCoordinationRequest                 *QMCCoordinationRequest
	SrcSNToTgtSNQMCInfoInquiry             *SrcSNToTgtSNQMCInfoInquiry
	IABAuthorizationStatus                 *IABAuthorizationStatus
}

func (msg *SNodeModificationRequest) toIEs() []XnAPMessageIE {
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
	if msg.SCGConfigurationQuery != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SCGConfigurationQuery)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.SCGConfigurationQuery,
		})
	}
	if msg.UEContextInfoSNModRequest != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_UEContextInfoSNModRequest)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.UEContextInfoSNModRequest,
		})
	}
	if len(msg.MNToSNContainer) > 0 {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_MNToSNContainer)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       &octetStringIE{Value: msg.MNToSNContainer},
		})
	}
	if msg.RequestedSplitSRB != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_RequestedSplitSRB)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.RequestedSplitSRB,
		})
	}
	if msg.RequestedSplitSRBrelease != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_RequestedSplitSRBrelease)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.RequestedSplitSRBrelease,
		})
	}
	if msg.DesiredActNotificationLevel != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_DesiredActNotificationLevel)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.DesiredActNotificationLevel,
		})
	}
	if msg.AdditionalDRBIDs != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_AdditionalDRBIDs)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.AdditionalDRBIDs,
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
	if msg.PCellID != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_PCellID)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.PCellID,
		})
	}
	if msg.NEDCTDMPattern != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_NEDCTDMPattern)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.NEDCTDMPattern,
		})
	}
	if msg.RequestedFastMCGRecoveryViaSRB3 != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_RequestedFastMCGRecoveryViaSRB3)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.RequestedFastMCGRecoveryViaSRB3,
		})
	}
	if msg.RequestedFastMCGRecoveryViaSRB3Release != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_RequestedFastMCGRecoveryViaSRB3Release)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.RequestedFastMCGRecoveryViaSRB3Release,
		})
	}
	if msg.SNTriggered != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SNTriggered)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.SNTriggered,
		})
	}
	if msg.TargetNodeID != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_TargetNodeID)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.TargetNodeID,
		})
	}
	if msg.PSCellHistoryInformationRetrieve != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_PSCellHistoryInformationRetrieve)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.PSCellHistoryInformationRetrieve,
		})
	}
	if msg.UEHistoryInformationFromTheUE != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_UEHistoryInformationFromTheUE)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.UEHistoryInformationFromTheUE,
		})
	}
	if msg.CHOinformationModReq != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_CHOinformationModReq)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.CHOinformationModReq,
		})
	}
	if msg.SCGActivationRequest != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SCGActivationRequest)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.SCGActivationRequest,
		})
	}
	if msg.CPAInformationModReq != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_CPAInformationModReq)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.CPAInformationModReq,
		})
	}
	if msg.CPCInformationUpdate != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_CPCInformationUpdate)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.CPCInformationUpdate,
		})
	}
	if msg.SNGRANnodeUESliceMBR != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SNGRANnodeUESliceMBR)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.SNGRANnodeUESliceMBR,
		})
	}
	if msg.ManagementBasedMDTPLMNModificationList != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_ManagementBasedMDTPLMNModificationList)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.ManagementBasedMDTPLMNModificationList,
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
	if msg.SrcSNToTgtSNQMCInfoInquiry != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SrcSNToTgtSNQMCInfoInquiry)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.SrcSNToTgtSNQMCInfoInquiry,
		})
	}
	if msg.IABAuthorizationStatus != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_IABAuthorizationStatus)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.IABAuthorizationStatus,
		})
	}
	return ies
}

func (msg *SNodeModificationRequest) Encode(w io.Writer) error {
	inner := per.NewEncoder(per.APER)
	if err := encodeBody(msg, inner); err != nil {
		return err
	}

	outer := per.NewEncoder(per.APER)

	if err := choiceXnMsg(outer, common.InitiatingMessage); err != nil {
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

func (msg *SNodeModificationRequest) Decode(wire []byte) error {
	d := per.NewDecoder(wire, per.APER)

	seq := d.NewSequenceDecoder(per.SequenceConstraints{Extensible: true})
	if err := seq.DecodeExtensionBit(); err != nil {
		return fmt.Errorf("decode SNodeModificationRequest extension bit: %w", err)
	}

	seqOf := d.NewSequenceOfDecoder(per.SizeConstraints{
		Extensible: false,
		Min:        common.Ptr(int64(0)),
		Max:        common.Ptr(int64(MaxProtocolIEs)),
	})
	n, err := seqOf.DecodeLength()
	if err != nil {
		return fmt.Errorf("decode SNodeModificationRequest protocolIEs length: %w", err)
	}

	seen := make(map[int64]bool, n)
	for i := range n {
		id, err := d.DecodeInteger(protocolIEIDConstraints)
		if err != nil {
			return fmt.Errorf("decode SNodeModificationRequest IE[%d] id: %w", i, err)
		}
		seen[id] = true

		var criticality Criticality
		if err := criticality.Decode(d); err != nil {
			return fmt.Errorf("decode SNodeModificationRequest IE[%d] criticality: %w", i, err)
		}

		payload, err := d.DecodeOpenType()
		if err != nil {
			return fmt.Errorf("decode SNodeModificationRequest IE[%d] value: %w", i, err)
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
		case common.ProtocolIEID_SelectedPLMN:
			v := &PLMNIdentity{}
			err = v.Decode(ies)
			msg.SelectedPLMN = v
		case common.ProtocolIEID_MobilityRestrictionList:
			v := &MobilityRestrictionList{}
			err = v.Decode(ies)
			msg.MobilityRestrictionList = v
		case common.ProtocolIEID_SCGConfigurationQuery:
			v := &SCGConfigurationQuery{}
			err = v.Decode(ies)
			msg.SCGConfigurationQuery = v
		case common.ProtocolIEID_UEContextInfoSNModRequest:
			v := &UEContextInfoSNModRequest{}
			err = v.Decode(ies)
			msg.UEContextInfoSNModRequest = v
		case common.ProtocolIEID_MNToSNContainer:
			v := octetStringIE{}
			err = v.Decode(ies)
			msg.MNToSNContainer = v.Value
		case common.ProtocolIEID_RequestedSplitSRB:
			v := &SplitSRBsTypes{}
			err = v.Decode(ies)
			msg.RequestedSplitSRB = v
		case common.ProtocolIEID_RequestedSplitSRBrelease:
			v := &SplitSRBsTypes{}
			err = v.Decode(ies)
			msg.RequestedSplitSRBrelease = v
		case common.ProtocolIEID_DesiredActNotificationLevel:
			v := &DesiredActNotificationLevel{}
			err = v.Decode(ies)
			msg.DesiredActNotificationLevel = v
		case common.ProtocolIEID_AdditionalDRBIDs:
			v := &DRBList{}
			err = v.Decode(ies)
			msg.AdditionalDRBIDs = v
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
		case common.ProtocolIEID_PCellID:
			v := &GlobalNGRANCellID{}
			err = v.Decode(ies)
			msg.PCellID = v
		case common.ProtocolIEID_NEDCTDMPattern:
			v := &NEDCTDMPattern{}
			err = v.Decode(ies)
			msg.NEDCTDMPattern = v
		case common.ProtocolIEID_RequestedFastMCGRecoveryViaSRB3:
			v := &RequestedFastMCGRecoveryViaSRB3{}
			err = v.Decode(ies)
			msg.RequestedFastMCGRecoveryViaSRB3 = v
		case common.ProtocolIEID_RequestedFastMCGRecoveryViaSRB3Release:
			v := &RequestedFastMCGRecoveryViaSRB3Release{}
			err = v.Decode(ies)
			msg.RequestedFastMCGRecoveryViaSRB3Release = v
		case common.ProtocolIEID_SNTriggered:
			v := &SNTriggered{}
			err = v.Decode(ies)
			msg.SNTriggered = v
		case common.ProtocolIEID_TargetNodeID:
			v := &GlobalNGRANNodeID{}
			err = v.Decode(ies)
			msg.TargetNodeID = v
		case common.ProtocolIEID_PSCellHistoryInformationRetrieve:
			v := &PSCellHistoryInformationRetrieve{}
			err = v.Decode(ies)
			msg.PSCellHistoryInformationRetrieve = v
		case common.ProtocolIEID_UEHistoryInformationFromTheUE:
			v := &UEHistoryInformationFromTheUE{}
			err = v.Decode(ies)
			msg.UEHistoryInformationFromTheUE = v
		case common.ProtocolIEID_CHOinformationModReq:
			v := &CHOinformationModReq{}
			err = v.Decode(ies)
			msg.CHOinformationModReq = v
		case common.ProtocolIEID_SCGActivationRequest:
			v := &SCGActivationRequest{}
			err = v.Decode(ies)
			msg.SCGActivationRequest = v
		case common.ProtocolIEID_CPAInformationModReq:
			v := &CPAInformationModReq{}
			err = v.Decode(ies)
			msg.CPAInformationModReq = v
		case common.ProtocolIEID_CPCInformationUpdate:
			v := &CPCInformationUpdate{}
			err = v.Decode(ies)
			msg.CPCInformationUpdate = v
		case common.ProtocolIEID_SNGRANnodeUESliceMBR:
			v := &UESliceMaximumBitRateList{}
			err = v.Decode(ies)
			msg.SNGRANnodeUESliceMBR = v
		case common.ProtocolIEID_ManagementBasedMDTPLMNModificationList:
			v := &MDTPLMNModificationList{}
			err = v.Decode(ies)
			msg.ManagementBasedMDTPLMNModificationList = v
		case common.ProtocolIEID_SelectedNID:
			v := &NID{}
			err = v.Decode(ies)
			msg.SelectedNID = v
		case common.ProtocolIEID_QMCCoordinationRequest:
			v := &QMCCoordinationRequest{}
			err = v.Decode(ies)
			msg.QMCCoordinationRequest = v
		case common.ProtocolIEID_SrcSNToTgtSNQMCInfoInquiry:
			v := &SrcSNToTgtSNQMCInfoInquiry{}
			err = v.Decode(ies)
			msg.SrcSNToTgtSNQMCInfoInquiry = v
		case common.ProtocolIEID_IABAuthorizationStatus:
			v := &IABAuthorizationStatus{}
			err = v.Decode(ies)
			msg.IABAuthorizationStatus = v
		default:
			if criticality.Value == CriticalityReject {
				return fmt.Errorf("unsupported SNodeModificationRequest IE id=%d with reject criticality", id)
			}
		}
		if err != nil {
			return fmt.Errorf("decode SNodeModificationRequest IE id=%d: %w", id, err)
		}
	}

	if _, err := seq.DecodeExtensionAdditions(); err != nil {
		return fmt.Errorf("decode SNodeModificationRequest extensions: %w", err)
	}

	for _, id := range []int64{
		common.ProtocolIEID_MNGRANnodeUEXnAPID,
		common.ProtocolIEID_SNGRANnodeUEXnAPID,
		common.ProtocolIEID_Cause,
	} {
		if !seen[id] {
			return fmt.Errorf("missing mandatory SNodeModificationRequest IE id=%d", id)
		}
	}

	return nil
}
