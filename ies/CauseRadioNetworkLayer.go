package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	CauseRadioNetworkLayerCellNotAvailable                                         int64 = 0
	CauseRadioNetworkLayerHandoverDesirableForRadioReasons                         int64 = 1
	CauseRadioNetworkLayerHandoverTargetNotAllowed                                 int64 = 2
	CauseRadioNetworkLayerInvalidAMFSetID                                          int64 = 3
	CauseRadioNetworkLayerNoRadioResourcesAvailableInTargetCell                    int64 = 4
	CauseRadioNetworkLayerPartialHandover                                          int64 = 5
	CauseRadioNetworkLayerReduceLoadInServingCell                                  int64 = 6
	CauseRadioNetworkLayerResourceOptimisationHandover                             int64 = 7
	CauseRadioNetworkLayerTimeCriticalHandover                                     int64 = 8
	CauseRadioNetworkLayerTXnRELOCoverallExpiry                                    int64 = 9
	CauseRadioNetworkLayerTXnRELOCprepExpiry                                       int64 = 10
	CauseRadioNetworkLayerUnknownGUAMIID                                           int64 = 11
	CauseRadioNetworkLayerUnknownLocalNGRANNodeUEXnAPID                            int64 = 12
	CauseRadioNetworkLayerInconsistentRemoteNGRANNodeUEXnAPID                      int64 = 13
	CauseRadioNetworkLayerEncryptionAndOrIntegrityProtectionAlgorithmsNotSupported int64 = 14
	CauseRadioNetworkLayerNotUsedCausesValue1                                      int64 = 15
	CauseRadioNetworkLayerMultiplePDUSessionIDInstances                            int64 = 16
	CauseRadioNetworkLayerUnknownPDUSessionID                                      int64 = 17
	CauseRadioNetworkLayerUnknownQoSFlowID                                         int64 = 18
	CauseRadioNetworkLayerMultipleQoSFlowIDInstances                               int64 = 19
	CauseRadioNetworkLayerSwitchOffOngoing                                         int64 = 20
	CauseRadioNetworkLayerNotSupported5QIValue                                     int64 = 21
	CauseRadioNetworkLayerTXnDCoverallExpiry                                       int64 = 22
	CauseRadioNetworkLayerTXnDCprepExpiry                                          int64 = 23
	CauseRadioNetworkLayerActionDesirableForRadioReasons                           int64 = 24
	CauseRadioNetworkLayerReduceLoad                                               int64 = 25
	CauseRadioNetworkLayerResourceOptimisation                                     int64 = 26
	CauseRadioNetworkLayerTimeCriticalAction                                       int64 = 27
	CauseRadioNetworkLayerTargetNotAllowed                                         int64 = 28
	CauseRadioNetworkLayerNoRadioResourcesAvailable                                int64 = 29
	CauseRadioNetworkLayerInvalidQoSCombination                                    int64 = 30
	CauseRadioNetworkLayerEncryptionAlgorithmsNotSupported                         int64 = 31
	CauseRadioNetworkLayerProcedureCancelled                                       int64 = 32
	CauseRadioNetworkLayerRRMPurpose                                               int64 = 33
	CauseRadioNetworkLayerImproveUserBitRate                                       int64 = 34
	CauseRadioNetworkLayerUserInactivity                                           int64 = 35
	CauseRadioNetworkLayerRadioConnectionWithUELost                                int64 = 36
	CauseRadioNetworkLayerFailureInTheRadioInterfaceProcedure                      int64 = 37
	CauseRadioNetworkLayerBearerOptionNotSupported                                 int64 = 38
	CauseRadioNetworkLayerUpIntegrityProtectionNotPossible                         int64 = 39
	CauseRadioNetworkLayerUpConfidentialityProtectionNotPossible                   int64 = 40
	CauseRadioNetworkLayerResourcesNotAvailableForTheSliceS                        int64 = 41
	CauseRadioNetworkLayerUeMaxIPDataRateReason                                    int64 = 42
	CauseRadioNetworkLayerCPIntegrityProtectionFailure                             int64 = 43
	CauseRadioNetworkLayerUPIntegrityProtectionFailure                             int64 = 44
	CauseRadioNetworkLayerSliceNotSupportedByNGRAN                                 int64 = 45
	CauseRadioNetworkLayerMNMobility                                               int64 = 46
	CauseRadioNetworkLayerSNMobility                                               int64 = 47
	CauseRadioNetworkLayerCountReachesMaxValue                                     int64 = 48
	CauseRadioNetworkLayerUnknownOldNGRANNodeUEXnAPID                              int64 = 49
	CauseRadioNetworkLayerPDCPOverload                                             int64 = 50
	CauseRadioNetworkLayerDrbIdNotAvailable                                        int64 = 51
	CauseRadioNetworkLayerUnspecified                                              int64 = 52
	CauseRadioNetworkLayerUeContextIdNotKnown                                      int64 = 53
	CauseRadioNetworkLayerNonRelocationOfContext                                   int64 = 54
	CauseRadioNetworkLayerChoCpcResourcesTobechanged                               int64 = 55
	CauseRadioNetworkLayerRSNNotAvailableForTheUP                                  int64 = 56
	CauseRadioNetworkLayerNpnAccessDenied                                          int64 = 57
	CauseRadioNetworkLayerReportCharacteristicsEmpty                               int64 = 58
	CauseRadioNetworkLayerExistingMeasurementID                                    int64 = 59
	CauseRadioNetworkLayerMeasurementTemporarilyNotAvailable                       int64 = 60
	CauseRadioNetworkLayerMeasurementNotSupportedForTheObject                      int64 = 61
	CauseRadioNetworkLayerUePowerSaving                                            int64 = 62
	CauseRadioNetworkLayerNotExistingNGRANNode2MeasurementID                       int64 = 63
	CauseRadioNetworkLayerInsufficientUeCapabilities                               int64 = 64
	CauseRadioNetworkLayerNormalRelease                                            int64 = 65
	CauseRadioNetworkLayerValueOutOfAllowedRange                                   int64 = 66
	CauseRadioNetworkLayerScgActivationDeactivationFailure                         int64 = 67
	CauseRadioNetworkLayerScgDeactivationFailureDueToDataTransmission              int64 = 68
	CauseRadioNetworkLayerSsbNotAvailable                                          int64 = 69
	CauseRadioNetworkLayerLTMTriggered                                             int64 = 70
	CauseRadioNetworkLayerNoBackhaulResource                                       int64 = 71
	CauseRadioNetworkLayerMIABNodeNotAuthorized                                    int64 = 72
	CauseRadioNetworkLayerIABNotAuthorized                                         int64 = 73
)

var causeRadioNetworkLayerConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52},
	ExtValues:  []int64{53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73},
}

type CauseRadioNetworkLayer struct {
	Value int64
}

func (ie *CauseRadioNetworkLayer) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, causeRadioNetworkLayerConstraints)
}

func (ie *CauseRadioNetworkLayer) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(causeRadioNetworkLayerConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
