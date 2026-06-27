package ies

import (
	"fmt"
	"io"

	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

// RetrieveUEContextResponse ::= SEQUENCE {
//	protocolIEs  ProtocolIE-Container {{RetrieveUEContextResponse-IEs}},
//	...
// }

type RetrieveUEContextResponse struct {
	NewNGRANnodeUEXnAPID                     NGRANnodeUEXnAPID
	OldNGRANnodeUEXnAPID                     NGRANnodeUEXnAPID
	GUAMI                                    GUAMI
	UEContextInfoRetrUECtxtResp              UEContextInfoRetrUECtxtResp
	TraceActivation                          *TraceActivation
	MaskedIMEISV                             *MaskedIMEISV
	LocationReportingInformation             *LocationReportingInformation
	CriticalityDiagnostics                   *CriticalityDiagnostics
	NRV2XServicesAuthorized                  *NRV2XServicesAuthorized
	LTEV2XServicesAuthorized                 *LTEV2XServicesAuthorized
	PC5QoSParameters                         *PC5QoSParameters
	UEHistoryInformation                     *UEHistoryInformation
	UEHistoryInformationFromTheUE            *UEHistoryInformationFromTheUE
	MDTPLMNList                              *MDTPLMNList
	IABNodeIndication                        *IABNodeIndication
	UEContextRefAtSNHORequest                *UEContextRefAtSNHORequest
	TimeSynchronizationAssistanceInformation *TimeSynchronizationAssistanceInformation
	QMCConfigInfo                            *QMCConfigInfo
	FiveGProSeAuthorized                     *FiveGProSeAuthorized
	FiveGProSePC5QoSParameters               *FiveGProSePC5QoSParameters
	AerialUESubscriptionInformation          *AerialUESubscriptionInformation
	NRA2XServicesAuthorized                  *NRA2XServicesAuthorized
	LTEA2XServicesAuthorized                 *LTEA2XServicesAuthorized
	A2XPC5QoSParameters                      *A2XPC5QoSParameters
	MobileIABAuthorizationStatus             *MobileIABAuthorizationStatus
	SLPositioningRangingServicesInfo         *SLPositioningRangingServicesInfo
}

func (msg *RetrieveUEContextResponse) toIEs() []XnAPMessageIE {
	ies := make([]XnAPMessageIE, 0)
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_NewNGRANnodeUEXnAPID)},
		Criticality: Criticality{Value: CriticalityIgnore},
		Value:       &msg.NewNGRANnodeUEXnAPID,
	})
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_OldNGRANnodeUEXnAPID)},
		Criticality: Criticality{Value: CriticalityIgnore},
		Value:       &msg.OldNGRANnodeUEXnAPID,
	})
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_GUAMI)},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.GUAMI,
	})
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_UEContextInfoRetrUECtxtResp)},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.UEContextInfoRetrUECtxtResp,
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
	if msg.LocationReportingInformation != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_LocationReportingInformation)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.LocationReportingInformation,
		})
	}
	if msg.CriticalityDiagnostics != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_CriticalityDiagnostics)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.CriticalityDiagnostics,
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
	if msg.MDTPLMNList != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_MDTPLMNList)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.MDTPLMNList,
		})
	}
	if msg.IABNodeIndication != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_IABNodeIndication)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.IABNodeIndication,
		})
	}
	if msg.UEContextRefAtSNHORequest != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_UEContextRefAtSNHORequest)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.UEContextRefAtSNHORequest,
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

func (msg *RetrieveUEContextResponse) Encode(w io.Writer) error {
	inner := per.NewEncoder(per.APER)
	if err := encodeBody(msg, inner); err != nil {
		return err
	}

	outer := per.NewEncoder(per.APER)

	if err := choiceXnMsg(outer, common.SuccessfulOutcome); err != nil {
		return err
	}
	if err := procedureCodeXnMsg(outer, common.ProcedureCode_RetrieveUEContext); err != nil {
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

func (msg *RetrieveUEContextResponse) Decode(wire []byte) error {
	d := per.NewDecoder(wire, per.APER)

	seq := d.NewSequenceDecoder(per.SequenceConstraints{Extensible: true})
	if err := seq.DecodeExtensionBit(); err != nil {
		return fmt.Errorf("decode RetrieveUEContextResponse extension bit: %w", err)
	}

	seqOf := d.NewSequenceOfDecoder(per.SizeConstraints{
		Extensible: false,
		Min:        common.Ptr(int64(0)),
		Max:        common.Ptr(int64(MaxProtocolIEs)),
	})
	n, err := seqOf.DecodeLength()
	if err != nil {
		return fmt.Errorf("decode RetrieveUEContextResponse protocolIEs length: %w", err)
	}

	seen := make(map[int64]bool, n)
	for i := range n {
		id, err := d.DecodeInteger(protocolIEIDConstraints)
		if err != nil {
			return fmt.Errorf("decode RetrieveUEContextResponse IE[%d] id: %w", i, err)
		}
		seen[id] = true

		var criticality Criticality
		if err := criticality.Decode(d); err != nil {
			return fmt.Errorf("decode RetrieveUEContextResponse IE[%d] criticality: %w", i, err)
		}

		payload, err := d.DecodeOpenType()
		if err != nil {
			return fmt.Errorf("decode RetrieveUEContextResponse IE[%d] value: %w", i, err)
		}
		ies := per.NewDecoder(payload, per.APER)

		switch id {
		case common.ProtocolIEID_NewNGRANnodeUEXnAPID:
			err = msg.NewNGRANnodeUEXnAPID.Decode(ies)
		case common.ProtocolIEID_OldNGRANnodeUEXnAPID:
			err = msg.OldNGRANnodeUEXnAPID.Decode(ies)
		case common.ProtocolIEID_GUAMI:
			err = msg.GUAMI.Decode(ies)
		case common.ProtocolIEID_UEContextInfoRetrUECtxtResp:
			err = msg.UEContextInfoRetrUECtxtResp.Decode(ies)
		case common.ProtocolIEID_TraceActivation:
			v := &TraceActivation{}
			err = v.Decode(ies)
			msg.TraceActivation = v
		case common.ProtocolIEID_MaskedIMEISV:
			v := &MaskedIMEISV{}
			err = v.Decode(ies)
			msg.MaskedIMEISV = v
		case common.ProtocolIEID_LocationReportingInformation:
			v := &LocationReportingInformation{}
			err = v.Decode(ies)
			msg.LocationReportingInformation = v
		case common.ProtocolIEID_CriticalityDiagnostics:
			v := &CriticalityDiagnostics{}
			err = v.Decode(ies)
			msg.CriticalityDiagnostics = v
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
		case common.ProtocolIEID_UEHistoryInformation:
			v := &UEHistoryInformation{}
			err = v.Decode(ies)
			msg.UEHistoryInformation = v
		case common.ProtocolIEID_UEHistoryInformationFromTheUE:
			v := &UEHistoryInformationFromTheUE{}
			err = v.Decode(ies)
			msg.UEHistoryInformationFromTheUE = v
		case common.ProtocolIEID_MDTPLMNList:
			v := &MDTPLMNList{}
			err = v.Decode(ies)
			msg.MDTPLMNList = v
		case common.ProtocolIEID_IABNodeIndication:
			v := &IABNodeIndication{}
			err = v.Decode(ies)
			msg.IABNodeIndication = v
		case common.ProtocolIEID_UEContextRefAtSNHORequest:
			v := &UEContextRefAtSNHORequest{}
			err = v.Decode(ies)
			msg.UEContextRefAtSNHORequest = v
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
				return fmt.Errorf("unsupported RetrieveUEContextResponse IE id=%d with reject criticality", id)
			}
		}
		if err != nil {
			return fmt.Errorf("decode RetrieveUEContextResponse IE id=%d: %w", id, err)
		}
	}

	if _, err := seq.DecodeExtensionAdditions(); err != nil {
		return fmt.Errorf("decode RetrieveUEContextResponse extensions: %w", err)
	}

	for _, id := range []int64{
		common.ProtocolIEID_NewNGRANnodeUEXnAPID,
		common.ProtocolIEID_OldNGRANnodeUEXnAPID,
		common.ProtocolIEID_GUAMI,
		common.ProtocolIEID_UEContextInfoRetrUECtxtResp,
	} {
		if !seen[id] {
			return fmt.Errorf("missing mandatory RetrieveUEContextResponse IE id=%d", id)
		}
	}

	return nil
}
