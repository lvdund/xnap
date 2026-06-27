package xn

import (
	"fmt"
	"io"

	"github.com/lvdund/xnap/common"
	"github.com/lvdund/xnap/ies"
)

type XnApPDU struct {
	Choice        uint8
	ProcedureCode ies.ProcedureCode
	Criticality   ies.Criticality
	Message       XnApMessageDecoder
}

// interface to message encoder all message need to implement this interface
type XnApMessageEncoder interface {
	Encode(io.Writer) error
}

// interface to message decoder all message need to implement this interface
type XnApMessageDecoder interface {
	Decode([]byte) error
}

// decodeMessage dispatches the decoded XnAP-PDU body to the concrete
// message identified by the (choice, procedureCode) pair. Each pair maps
// to exactly one message type (see 1_Elementary_Procedure_definitions.asn1).
func decodeMessage(choice uint8, procCode int64, body []byte) (XnApMessageDecoder, error) {
	switch choice {
	case uint8(common.InitiatingMessage):
		switch procCode {
		case common.ProcedureCode_AccessAndMobilityIndication:
			m := &ies.AccessAndMobilityIndication{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_ActivityNotification:
			m := &ies.ActivityNotification{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_CPCCancel:
			m := &ies.CPCCancel{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_CellActivation:
			m := &ies.CellActivationRequest{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_CellTrafficTrace:
			m := &ies.CellTrafficTrace{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_ConditionalHandoverCancel:
			m := &ies.ConditionalHandoverCancel{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_DataCollectionReporting:
			m := &ies.DataCollectionUpdate{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_DataCollectionReportingInitiation:
			m := &ies.DataCollectionRequest{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_DeactivateTrace:
			m := &ies.DeactivateTrace{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_EUTRANRCellResourceCoordination:
			m := &ies.EUTRANRCellResourceCoordinationRequest{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_EarlyStatusTransfer:
			m := &ies.EarlyStatusTransfer{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_ErrorIndication:
			m := &ies.ErrorIndication{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_F1CTrafficTransfer:
			m := &ies.F1CTrafficTransfer{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_FailureIndication:
			m := &ies.FailureIndication{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_HandoverCancel:
			m := &ies.HandoverCancel{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_HandoverPreparation:
			m := &ies.HandoverRequest{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_HandoverReport:
			m := &ies.HandoverReport{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_HandoverSuccess:
			m := &ies.HandoverSuccess{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_IABResourceCoordination:
			m := &ies.IABResourceCoordinationRequest{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_IABTransportMigrationManagement:
			m := &ies.IABTransportMigrationManagementRequest{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_IABTransportMigrationModification:
			m := &ies.IABTransportMigrationModificationRequest{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_MNGRANnodeinitiatedSNGRANnodeModificationPreparation:
			m := &ies.SNodeModificationRequest{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_MNGRANnodeinitiatedSNGRANnodeRelease:
			m := &ies.SNodeReleaseRequest{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_MobilitySettingsChange:
			m := &ies.MobilityChangeRequest{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_NGRANnodeConfigurationUpdate:
			m := &ies.NGRANNodeConfigurationUpdate{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_NotificationControl:
			m := &ies.NotificationControlIndication{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_PartialUEContextTransfer:
			m := &ies.PartialUEContextTransfer{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_RANMulticastGroupPaging:
			m := &ies.RANMulticastGroupPaging{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_RANPaging:
			m := &ies.RANPaging{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_RRCTransfer:
			m := &ies.RRCTransfer{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_RachIndication:
			m := &ies.RachIndication{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_Reset:
			m := &ies.ResetRequest{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_ResourceStatusReporting:
			m := &ies.ResourceStatusUpdate{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_ResourceStatusReportingInitiation:
			m := &ies.ResourceStatusRequest{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_RetrieveUEContext:
			m := &ies.RetrieveUEContextRequest{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_RetrieveUEContextConfirm:
			m := &ies.RetrieveUEContextConfirm{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_SNGRANnodeAdditionPreparation:
			m := &ies.SNodeAdditionRequest{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_SNGRANnodeChange:
			m := &ies.SNodeChangeRequired{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_SNGRANnodeCounterCheck:
			m := &ies.SNodeCounterCheckRequest{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_SNGRANnodeReconfigurationCompletion:
			m := &ies.SNodeReconfigurationComplete{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_SNGRANnodeinitiatedSNGRANnodeModificationPreparation:
			m := &ies.SNodeModificationRequired{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_SNGRANnodeinitiatedSNGRANnodeRelease:
			m := &ies.SNodeReleaseRequired{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_SNStatusTransfer:
			m := &ies.SNStatusTransfer{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_ScgFailureInformationReport:
			m := &ies.ScgFailureInformationReport{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_ScgFailureTransfer:
			m := &ies.ScgFailureTransfer{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_SecondaryRATDataUsageReport:
			m := &ies.SecondaryRATDataUsageReport{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_TraceStart:
			m := &ies.TraceStart{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_UEContextRelease:
			m := &ies.UEContextRelease{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_XnRemoval:
			m := &ies.XnRemovalRequest{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_XnSetup:
			m := &ies.XnSetupRequest{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_XnUAddressIndication:
			m := &ies.XnUAddressIndication{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		default:
			return nil, fmt.Errorf("unsupported XnAP initiating-message procedureCode=%d", procCode)
		}
	case uint8(common.SuccessfulOutcome):
		switch procCode {
		case common.ProcedureCode_CellActivation:
			m := &ies.CellActivationResponse{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_DataCollectionReportingInitiation:
			m := &ies.DataCollectionResponse{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_EUTRANRCellResourceCoordination:
			m := &ies.EUTRANRCellResourceCoordinationResponse{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_HandoverPreparation:
			m := &ies.HandoverRequestAcknowledge{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_IABResourceCoordination:
			m := &ies.IABResourceCoordinationResponse{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_IABTransportMigrationManagement:
			m := &ies.IABTransportMigrationManagementResponse{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_IABTransportMigrationModification:
			m := &ies.IABTransportMigrationModificationResponse{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_MNGRANnodeinitiatedSNGRANnodeModificationPreparation:
			m := &ies.SNodeModificationRequestAcknowledge{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_MNGRANnodeinitiatedSNGRANnodeRelease:
			m := &ies.SNodeReleaseRequestAcknowledge{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_MobilitySettingsChange:
			m := &ies.MobilityChangeAcknowledge{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_NGRANnodeConfigurationUpdate:
			m := &ies.NGRANNodeConfigurationUpdateAcknowledge{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_PartialUEContextTransfer:
			m := &ies.PartialUEContextTransferAcknowledge{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_Reset:
			m := &ies.ResetResponse{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_ResourceStatusReportingInitiation:
			m := &ies.ResourceStatusResponse{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_RetrieveUEContext:
			m := &ies.RetrieveUEContextResponse{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_SNGRANnodeAdditionPreparation:
			m := &ies.SNodeAdditionRequestAcknowledge{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_SNGRANnodeChange:
			m := &ies.SNodeChangeConfirm{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_SNGRANnodeinitiatedSNGRANnodeModificationPreparation:
			m := &ies.SNodeModificationConfirm{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_SNGRANnodeinitiatedSNGRANnodeRelease:
			m := &ies.SNodeReleaseConfirm{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_XnRemoval:
			m := &ies.XnRemovalResponse{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_XnSetup:
			m := &ies.XnSetupResponse{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		default:
			return nil, fmt.Errorf("unsupported XnAP successful-outcome procedureCode=%d", procCode)
		}
	case uint8(common.UnsuccessfulOutcome):
		switch procCode {
		case common.ProcedureCode_CellActivation:
			m := &ies.CellActivationFailure{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_DataCollectionReportingInitiation:
			m := &ies.DataCollectionFailure{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_HandoverPreparation:
			m := &ies.HandoverPreparationFailure{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_IABTransportMigrationManagement:
			m := &ies.IABTransportMigrationManagementReject{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_MNGRANnodeinitiatedSNGRANnodeModificationPreparation:
			m := &ies.SNodeModificationRequestReject{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_MNGRANnodeinitiatedSNGRANnodeRelease:
			m := &ies.SNodeReleaseReject{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_MobilitySettingsChange:
			m := &ies.MobilityChangeFailure{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_NGRANnodeConfigurationUpdate:
			m := &ies.NGRANNodeConfigurationUpdateFailure{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_PartialUEContextTransfer:
			m := &ies.PartialUEContextTransferFailure{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_ResourceStatusReportingInitiation:
			m := &ies.ResourceStatusFailure{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_RetrieveUEContext:
			m := &ies.RetrieveUEContextFailure{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_SNGRANnodeAdditionPreparation:
			m := &ies.SNodeAdditionRequestReject{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_SNGRANnodeChange:
			m := &ies.SNodeChangeRefuse{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_SNGRANnodeinitiatedSNGRANnodeModificationPreparation:
			m := &ies.SNodeModificationRefuse{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_XnRemoval:
			m := &ies.XnRemovalFailure{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		case common.ProcedureCode_XnSetup:
			m := &ies.XnSetupFailure{}
			if err := m.Decode(body); err != nil {
				return nil, err
			}
			return m, nil
		default:
			return nil, fmt.Errorf("unsupported XnAP unsuccessful-outcome procedureCode=%d", procCode)
		}
	default:
		return nil, fmt.Errorf("unsupported XnAP-PDU choice=%d", choice)
	}
}
