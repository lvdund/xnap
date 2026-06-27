package ies

import (
	"fmt"
	"io"

	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

// DataCollectionRequest ::= SEQUENCE {
//	protocolIEs  ProtocolIE-Container {{DataCollectionRequest-IEs}},
//	...
// }

type DataCollectionRequest struct {
	NGRANNode1MeasurementID                MeasurementID
	NGRANNode2MeasurementID                *MeasurementID
	RegistrationRequestForDataCollection   RegistrationRequestForDataCollection
	ReportCharacteristicsForDataCollection *ReportCharacteristicsForDataCollection
	CellToReportForDataCollectionList      *CellToReportForDataCollectionList
	ReportingPeriodicityForDataCollection  *ReportingPeriodicityForDataCollection
	RequestedPredictionTime                *RequestedPredictionTime
	UETrajectoryCollectionConfiguration    *UETrajectoryCollectionConfiguration
	UEPerformanceCollectionConfiguration   *UEPerformanceCollectionConfiguration
}

func (msg *DataCollectionRequest) toIEs() []XnAPMessageIE {
	ies := make([]XnAPMessageIE, 0)
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_NGRANNode1MeasurementID)},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.NGRANNode1MeasurementID,
	})
	if msg.NGRANNode2MeasurementID != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_NGRANNode2MeasurementID)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.NGRANNode2MeasurementID,
		})
	}
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_RegistrationRequestForDataCollection)},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.RegistrationRequestForDataCollection,
	})
	if msg.ReportCharacteristicsForDataCollection != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_ReportCharacteristicsForDataCollection)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.ReportCharacteristicsForDataCollection,
		})
	}
	if msg.CellToReportForDataCollectionList != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_CellToReportForDataCollectionList)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.CellToReportForDataCollectionList,
		})
	}
	if msg.ReportingPeriodicityForDataCollection != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_ReportingPeriodicityForDataCollection)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.ReportingPeriodicityForDataCollection,
		})
	}
	if msg.RequestedPredictionTime != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_RequestedPredictionTime)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.RequestedPredictionTime,
		})
	}
	if msg.UETrajectoryCollectionConfiguration != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_UETrajectoryCollectionConfiguration)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.UETrajectoryCollectionConfiguration,
		})
	}
	if msg.UEPerformanceCollectionConfiguration != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_UEPerformanceCollectionConfiguration)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.UEPerformanceCollectionConfiguration,
		})
	}
	return ies
}

func (msg *DataCollectionRequest) Encode(w io.Writer) error {
	inner := per.NewEncoder(per.APER)
	if err := encodeBody(msg, inner); err != nil {
		return err
	}

	outer := per.NewEncoder(per.APER)

	if err := choiceXnMsg(outer, common.InitiatingMessage); err != nil {
		return err
	}
	if err := procedureCodeXnMsg(outer, common.ProcedureCode_DataCollectionReportingInitiation); err != nil {
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

func (msg *DataCollectionRequest) Decode(wire []byte) error {
	d := per.NewDecoder(wire, per.APER)

	seq := d.NewSequenceDecoder(per.SequenceConstraints{Extensible: true})
	if err := seq.DecodeExtensionBit(); err != nil {
		return fmt.Errorf("decode DataCollectionRequest extension bit: %w", err)
	}

	seqOf := d.NewSequenceOfDecoder(per.SizeConstraints{
		Extensible: false,
		Min:        common.Ptr(int64(0)),
		Max:        common.Ptr(int64(MaxProtocolIEs)),
	})
	n, err := seqOf.DecodeLength()
	if err != nil {
		return fmt.Errorf("decode DataCollectionRequest protocolIEs length: %w", err)
	}

	seen := make(map[int64]bool, n)
	for i := range n {
		id, err := d.DecodeInteger(protocolIEIDConstraints)
		if err != nil {
			return fmt.Errorf("decode DataCollectionRequest IE[%d] id: %w", i, err)
		}
		seen[id] = true

		var criticality Criticality
		if err := criticality.Decode(d); err != nil {
			return fmt.Errorf("decode DataCollectionRequest IE[%d] criticality: %w", i, err)
		}

		payload, err := d.DecodeOpenType()
		if err != nil {
			return fmt.Errorf("decode DataCollectionRequest IE[%d] value: %w", i, err)
		}
		ies := per.NewDecoder(payload, per.APER)

		switch id {
		case common.ProtocolIEID_NGRANNode1MeasurementID:
			err = msg.NGRANNode1MeasurementID.Decode(ies)
		case common.ProtocolIEID_NGRANNode2MeasurementID:
			v := &MeasurementID{}
			err = v.Decode(ies)
			msg.NGRANNode2MeasurementID = v
		case common.ProtocolIEID_RegistrationRequestForDataCollection:
			err = msg.RegistrationRequestForDataCollection.Decode(ies)
		case common.ProtocolIEID_ReportCharacteristicsForDataCollection:
			v := &ReportCharacteristicsForDataCollection{}
			err = v.Decode(ies)
			msg.ReportCharacteristicsForDataCollection = v
		case common.ProtocolIEID_CellToReportForDataCollectionList:
			v := &CellToReportForDataCollectionList{}
			err = v.Decode(ies)
			msg.CellToReportForDataCollectionList = v
		case common.ProtocolIEID_ReportingPeriodicityForDataCollection:
			v := &ReportingPeriodicityForDataCollection{}
			err = v.Decode(ies)
			msg.ReportingPeriodicityForDataCollection = v
		case common.ProtocolIEID_RequestedPredictionTime:
			v := &RequestedPredictionTime{}
			err = v.Decode(ies)
			msg.RequestedPredictionTime = v
		case common.ProtocolIEID_UETrajectoryCollectionConfiguration:
			v := &UETrajectoryCollectionConfiguration{}
			err = v.Decode(ies)
			msg.UETrajectoryCollectionConfiguration = v
		case common.ProtocolIEID_UEPerformanceCollectionConfiguration:
			v := &UEPerformanceCollectionConfiguration{}
			err = v.Decode(ies)
			msg.UEPerformanceCollectionConfiguration = v
		default:
			if criticality.Value == CriticalityReject {
				return fmt.Errorf("unsupported DataCollectionRequest IE id=%d with reject criticality", id)
			}
		}
		if err != nil {
			return fmt.Errorf("decode DataCollectionRequest IE id=%d: %w", id, err)
		}
	}

	if _, err := seq.DecodeExtensionAdditions(); err != nil {
		return fmt.Errorf("decode DataCollectionRequest extensions: %w", err)
	}

	for _, id := range []int64{
		common.ProtocolIEID_NGRANNode1MeasurementID,
		common.ProtocolIEID_RegistrationRequestForDataCollection,
	} {
		if !seen[id] {
			return fmt.Errorf("missing mandatory DataCollectionRequest IE id=%d", id)
		}
	}

	return nil
}
