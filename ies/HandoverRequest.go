package ies

import (
	"fmt"
	"io"

	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

// HandoverRequest ::= SEQUENCE {
//	protocolIEs  ProtocolIE-Container {{HandoverRequest-IEs}},
//	...
// }

type HandoverRequest struct {
	SourceNGRANnodeUEXnAPID                  NGRANnodeUEXnAPID
	Cause                                    Cause
	TargetCellGlobalID                       TargetCGI
	GUAMI                                    GUAMI
	UEContextInfoHORequest                   UEContextInfoHORequest
	TraceActivation                          *TraceActivation
	MaskedIMEISV                             *MaskedIMEISV
	UEHistoryInformation                     UEHistoryInformation
	UEContextRefAtSNHORequest                *UEContextRefAtSNHORequest
	CHOinformationReq                        *CHOinformationReq
	NRV2XServicesAuthorized                  *NRV2XServicesAuthorized
	LTEV2XServicesAuthorized                 *LTEV2XServicesAuthorized
	PC5QoSParameters                         *PC5QoSParameters
	MobilityInformation                      *MobilityInformation
	UEHistoryInformationFromTheUE            *UEHistoryInformationFromTheUE
	IABNodeIndication                        *IABNodeIndication
	NoPDUSessionIndication                   *NoPDUSessionIndication
	TimeSynchronizationAssistanceInformation *TimeSynchronizationAssistanceInformation
	QMCConfigInfo                            *QMCConfigInfo
	FiveGProSeAuthorized                     *FiveGProSeAuthorized
	FiveGProSePC5QoSParameters               *FiveGProSePC5QoSParameters
	IABAuthorizationStatus                   *IABAuthorizationStatus
	DLLBTFailureInformationRequest           *DLLBTFailureInformationRequest
	AerialUESubscriptionInformation          *AerialUESubscriptionInformation
	NRA2XServicesAuthorized                  *NRA2XServicesAuthorized
	LTEA2XServicesAuthorized                 *LTEA2XServicesAuthorized
	A2XPC5QoSParameters                      *A2XPC5QoSParameters
	CellBasedUETrajectoryPrediction          *CellBasedUETrajectoryPrediction
	DataCollectionID                         *DataCollectionID
	CandidateRelayUEInfoList                 *CandidateRelayUEInfoList
	SourceSNToTargetSNQMCInfo                *QMCConfigInfo
	MobileIABAuthorizationStatus             *MobileIABAuthorizationStatus
	SLPositioningRangingServicesInfo         *SLPositioningRangingServicesInfo
}

func (msg *HandoverRequest) toIEs() []XnAPMessageIE {
	ies := make([]XnAPMessageIE, 0)
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SourceNGRANnodeUEXnAPID)},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.SourceNGRANnodeUEXnAPID,
	})
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_Cause)},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.Cause,
	})
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_TargetCellGlobalID)},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.TargetCellGlobalID,
	})
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_GUAMI)},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.GUAMI,
	})
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_UEContextInfoHORequest)},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.UEContextInfoHORequest,
	})
	if msg.TraceActivation != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_TraceActivation)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.TraceActivation,
		})
	}
	if msg.MaskedIMEISV != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_MaskedIMEISV)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.MaskedIMEISV,
		})
	}
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_UEHistoryInformation)},
		Criticality: Criticality{Value: CriticalityIgnore},
		Value:       &msg.UEHistoryInformation,
	})
	if msg.UEContextRefAtSNHORequest != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_UEContextRefAtSNHORequest)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.UEContextRefAtSNHORequest,
		})
	}
	if msg.CHOinformationReq != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_CHOinformationReq)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.CHOinformationReq,
		})
	}
	if msg.NRV2XServicesAuthorized != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_NRV2XServicesAuthorized)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.NRV2XServicesAuthorized,
		})
	}
	if msg.LTEV2XServicesAuthorized != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_LTEV2XServicesAuthorized)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.LTEV2XServicesAuthorized,
		})
	}
	if msg.PC5QoSParameters != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_PC5QoSParameters)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.PC5QoSParameters,
		})
	}
	if msg.MobilityInformation != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_MobilityInformation)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.MobilityInformation,
		})
	}
	if msg.UEHistoryInformationFromTheUE != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_UEHistoryInformationFromTheUE)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.UEHistoryInformationFromTheUE,
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
	if msg.TimeSynchronizationAssistanceInformation != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_TimeSynchronizationAssistanceInformation)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.TimeSynchronizationAssistanceInformation,
		})
	}
	if msg.QMCConfigInfo != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_QMCConfigInfo)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.QMCConfigInfo,
		})
	}
	if msg.FiveGProSeAuthorized != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_FiveGProSeAuthorized)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.FiveGProSeAuthorized,
		})
	}
	if msg.FiveGProSePC5QoSParameters != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_FiveGProSePC5QoSParameters)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.FiveGProSePC5QoSParameters,
		})
	}
	if msg.IABAuthorizationStatus != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_IABAuthorizationStatus)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.IABAuthorizationStatus,
		})
	}
	if msg.DLLBTFailureInformationRequest != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_DLLBTFailureInformationRequest)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.DLLBTFailureInformationRequest,
		})
	}
	if msg.AerialUESubscriptionInformation != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_AerialUESubscriptionInformation)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.AerialUESubscriptionInformation,
		})
	}
	if msg.NRA2XServicesAuthorized != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_NRA2XServicesAuthorized)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.NRA2XServicesAuthorized,
		})
	}
	if msg.LTEA2XServicesAuthorized != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_LTEA2XServicesAuthorized)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.LTEA2XServicesAuthorized,
		})
	}
	if msg.A2XPC5QoSParameters != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_A2XPC5QoSParameters)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.A2XPC5QoSParameters,
		})
	}
	if msg.CellBasedUETrajectoryPrediction != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_CellBasedUETrajectoryPrediction)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.CellBasedUETrajectoryPrediction,
		})
	}
	if msg.DataCollectionID != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_DataCollectionID)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.DataCollectionID,
		})
	}
	if msg.CandidateRelayUEInfoList != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_CandidateRelayUEInfoList)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.CandidateRelayUEInfoList,
		})
	}
	if msg.SourceSNToTargetSNQMCInfo != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SourceSNToTargetSNQMCInfo)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.SourceSNToTargetSNQMCInfo,
		})
	}
	if msg.MobileIABAuthorizationStatus != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_MobileIABAuthorizationStatus)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.MobileIABAuthorizationStatus,
		})
	}
	if msg.SLPositioningRangingServicesInfo != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SLPositioningRangingServicesInfo)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.SLPositioningRangingServicesInfo,
		})
	}
	return ies
}

func (msg *HandoverRequest) Encode(w io.Writer) error {
	inner := per.NewEncoder(per.APER)
	if err := encodeBody(msg, inner); err != nil {
		return err
	}

	outer := per.NewEncoder(per.APER)

	if err := choiceXnMsg(outer, common.InitiatingMessage); err != nil {
		return err
	}
	if err := procedureCodeXnMsg(outer, common.ProcedureCode_HandoverPreparation); err != nil {
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

func (msg *HandoverRequest) Decode(wire []byte) error {
	d := per.NewDecoder(wire, per.APER)

	seq := d.NewSequenceDecoder(per.SequenceConstraints{Extensible: true})
	if err := seq.DecodeExtensionBit(); err != nil {
		return fmt.Errorf("decode HandoverRequest extension bit: %w", err)
	}

	seqOf := d.NewSequenceOfDecoder(per.SizeConstraints{
		Extensible: false,
		Min:        common.Ptr(int64(0)),
		Max:        common.Ptr(int64(MaxProtocolIEs)),
	})
	n, err := seqOf.DecodeLength()
	if err != nil {
		return fmt.Errorf("decode HandoverRequest protocolIEs length: %w", err)
	}

	seen := make(map[int64]bool, n)
	for i := range n {
		id, err := d.DecodeInteger(protocolIEIDConstraints)
		if err != nil {
			return fmt.Errorf("decode HandoverRequest IE[%d] id: %w", i, err)
		}
		seen[id] = true

		var criticality Criticality
		if err := criticality.Decode(d); err != nil {
			return fmt.Errorf("decode HandoverRequest IE[%d] criticality: %w", i, err)
		}

		payload, err := d.DecodeOpenType()
		if err != nil {
			return fmt.Errorf("decode HandoverRequest IE[%d] value: %w", i, err)
		}
		ies := per.NewDecoder(payload, per.APER)

		switch id {
		case common.ProtocolIEID_SourceNGRANnodeUEXnAPID:
			err = msg.SourceNGRANnodeUEXnAPID.Decode(ies)
		case common.ProtocolIEID_Cause:
			err = msg.Cause.Decode(ies)
		case common.ProtocolIEID_TargetCellGlobalID:
			err = msg.TargetCellGlobalID.Decode(ies)
		case common.ProtocolIEID_GUAMI:
			err = msg.GUAMI.Decode(ies)
		case common.ProtocolIEID_UEContextInfoHORequest:
			err = msg.UEContextInfoHORequest.Decode(ies)
		case common.ProtocolIEID_TraceActivation:
			v := &TraceActivation{}
			err = v.Decode(ies)
			msg.TraceActivation = v
		case common.ProtocolIEID_MaskedIMEISV:
			v := &MaskedIMEISV{}
			err = v.Decode(ies)
			msg.MaskedIMEISV = v
		case common.ProtocolIEID_UEHistoryInformation:
			err = msg.UEHistoryInformation.Decode(ies)
		case common.ProtocolIEID_UEContextRefAtSNHORequest:
			v := &UEContextRefAtSNHORequest{}
			err = v.Decode(ies)
			msg.UEContextRefAtSNHORequest = v
		case common.ProtocolIEID_CHOinformationReq:
			v := &CHOinformationReq{}
			err = v.Decode(ies)
			msg.CHOinformationReq = v
		case common.ProtocolIEID_NRV2XServicesAuthorized:
			v := &NRV2XServicesAuthorized{}
			err = v.Decode(ies)
			msg.NRV2XServicesAuthorized = v
		case common.ProtocolIEID_LTEV2XServicesAuthorized:
			v := &LTEV2XServicesAuthorized{}
			err = v.Decode(ies)
			msg.LTEV2XServicesAuthorized = v
		case common.ProtocolIEID_PC5QoSParameters:
			v := &PC5QoSParameters{}
			err = v.Decode(ies)
			msg.PC5QoSParameters = v
		case common.ProtocolIEID_MobilityInformation:
			v := &MobilityInformation{}
			err = v.Decode(ies)
			msg.MobilityInformation = v
		case common.ProtocolIEID_UEHistoryInformationFromTheUE:
			v := &UEHistoryInformationFromTheUE{}
			err = v.Decode(ies)
			msg.UEHistoryInformationFromTheUE = v
		case common.ProtocolIEID_IABNodeIndication:
			v := &IABNodeIndication{}
			err = v.Decode(ies)
			msg.IABNodeIndication = v
		case common.ProtocolIEID_NoPDUSessionIndication:
			v := &NoPDUSessionIndication{}
			err = v.Decode(ies)
			msg.NoPDUSessionIndication = v
		case common.ProtocolIEID_TimeSynchronizationAssistanceInformation:
			v := &TimeSynchronizationAssistanceInformation{}
			err = v.Decode(ies)
			msg.TimeSynchronizationAssistanceInformation = v
		case common.ProtocolIEID_QMCConfigInfo:
			v := &QMCConfigInfo{}
			err = v.Decode(ies)
			msg.QMCConfigInfo = v
		case common.ProtocolIEID_FiveGProSeAuthorized:
			v := &FiveGProSeAuthorized{}
			err = v.Decode(ies)
			msg.FiveGProSeAuthorized = v
		case common.ProtocolIEID_FiveGProSePC5QoSParameters:
			v := &FiveGProSePC5QoSParameters{}
			err = v.Decode(ies)
			msg.FiveGProSePC5QoSParameters = v
		case common.ProtocolIEID_IABAuthorizationStatus:
			v := &IABAuthorizationStatus{}
			err = v.Decode(ies)
			msg.IABAuthorizationStatus = v
		case common.ProtocolIEID_DLLBTFailureInformationRequest:
			v := &DLLBTFailureInformationRequest{}
			err = v.Decode(ies)
			msg.DLLBTFailureInformationRequest = v
		case common.ProtocolIEID_AerialUESubscriptionInformation:
			v := &AerialUESubscriptionInformation{}
			err = v.Decode(ies)
			msg.AerialUESubscriptionInformation = v
		case common.ProtocolIEID_NRA2XServicesAuthorized:
			v := &NRA2XServicesAuthorized{}
			err = v.Decode(ies)
			msg.NRA2XServicesAuthorized = v
		case common.ProtocolIEID_LTEA2XServicesAuthorized:
			v := &LTEA2XServicesAuthorized{}
			err = v.Decode(ies)
			msg.LTEA2XServicesAuthorized = v
		case common.ProtocolIEID_A2XPC5QoSParameters:
			v := &A2XPC5QoSParameters{}
			err = v.Decode(ies)
			msg.A2XPC5QoSParameters = v
		case common.ProtocolIEID_CellBasedUETrajectoryPrediction:
			v := &CellBasedUETrajectoryPrediction{}
			err = v.Decode(ies)
			msg.CellBasedUETrajectoryPrediction = v
		case common.ProtocolIEID_DataCollectionID:
			v := &DataCollectionID{}
			err = v.Decode(ies)
			msg.DataCollectionID = v
		case common.ProtocolIEID_CandidateRelayUEInfoList:
			v := &CandidateRelayUEInfoList{}
			err = v.Decode(ies)
			msg.CandidateRelayUEInfoList = v
		case common.ProtocolIEID_SourceSNToTargetSNQMCInfo:
			v := &QMCConfigInfo{}
			err = v.Decode(ies)
			msg.SourceSNToTargetSNQMCInfo = v
		case common.ProtocolIEID_MobileIABAuthorizationStatus:
			v := &MobileIABAuthorizationStatus{}
			err = v.Decode(ies)
			msg.MobileIABAuthorizationStatus = v
		case common.ProtocolIEID_SLPositioningRangingServicesInfo:
			v := &SLPositioningRangingServicesInfo{}
			err = v.Decode(ies)
			msg.SLPositioningRangingServicesInfo = v
		default:
			if criticality.Value == CriticalityReject {
				return fmt.Errorf("unsupported HandoverRequest IE id=%d with reject criticality", id)
			}
		}
		if err != nil {
			return fmt.Errorf("decode HandoverRequest IE id=%d: %w", id, err)
		}
	}

	if _, err := seq.DecodeExtensionAdditions(); err != nil {
		return fmt.Errorf("decode HandoverRequest extensions: %w", err)
	}

	for _, id := range []int64{
		common.ProtocolIEID_SourceNGRANnodeUEXnAPID,
		common.ProtocolIEID_Cause,
		common.ProtocolIEID_TargetCellGlobalID,
		common.ProtocolIEID_GUAMI,
		common.ProtocolIEID_UEContextInfoHORequest,
		common.ProtocolIEID_UEHistoryInformation,
	} {
		if !seen[id] {
			return fmt.Errorf("missing mandatory HandoverRequest IE id=%d", id)
		}
	}

	return nil
}
