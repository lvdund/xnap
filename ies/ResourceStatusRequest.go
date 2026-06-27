package ies

import (
	"fmt"
	"io"

	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

// ResourceStatusRequest ::= SEQUENCE {
//	protocolIEs  ProtocolIE-Container {{ResourceStatusRequest-IEs}},
//	...
// }

type ResourceStatusRequest struct {
	NGRANNode1MeasurementID MeasurementID
	NGRANNode2MeasurementID *MeasurementID
	RegistrationRequest     RegistrationRequest
	ReportCharacteristics   *ReportCharacteristics
	CellToReport            *CellToReport
	ReportingPeriodicity    *ReportingPeriodicity
}

func (msg *ResourceStatusRequest) toIEs() []XnAPMessageIE {
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
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_RegistrationRequest)},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.RegistrationRequest,
	})
	if msg.ReportCharacteristics != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_ReportCharacteristics)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.ReportCharacteristics,
		})
	}
	if msg.CellToReport != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_CellToReport)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.CellToReport,
		})
	}
	if msg.ReportingPeriodicity != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_ReportingPeriodicity)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.ReportingPeriodicity,
		})
	}
	return ies
}

func (msg *ResourceStatusRequest) Encode(w io.Writer) error {
	inner := per.NewEncoder(per.APER)
	if err := encodeBody(msg, inner); err != nil {
		return err
	}

	outer := per.NewEncoder(per.APER)

	if err := choiceXnMsg(outer, common.InitiatingMessage); err != nil {
		return err
	}
	if err := procedureCodeXnMsg(outer, common.ProcedureCode_ResourceStatusReportingInitiation); err != nil {
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

func (msg *ResourceStatusRequest) Decode(wire []byte) error {
	d := per.NewDecoder(wire, per.APER)

	seq := d.NewSequenceDecoder(per.SequenceConstraints{Extensible: true})
	if err := seq.DecodeExtensionBit(); err != nil {
		return fmt.Errorf("decode ResourceStatusRequest extension bit: %w", err)
	}

	seqOf := d.NewSequenceOfDecoder(per.SizeConstraints{
		Extensible: false,
		Min:        common.Ptr(int64(0)),
		Max:        common.Ptr(int64(MaxProtocolIEs)),
	})
	n, err := seqOf.DecodeLength()
	if err != nil {
		return fmt.Errorf("decode ResourceStatusRequest protocolIEs length: %w", err)
	}

	seen := make(map[int64]bool, n)
	for i := range n {
		id, err := d.DecodeInteger(protocolIEIDConstraints)
		if err != nil {
			return fmt.Errorf("decode ResourceStatusRequest IE[%d] id: %w", i, err)
		}
		seen[id] = true

		var criticality Criticality
		if err := criticality.Decode(d); err != nil {
			return fmt.Errorf("decode ResourceStatusRequest IE[%d] criticality: %w", i, err)
		}

		payload, err := d.DecodeOpenType()
		if err != nil {
			return fmt.Errorf("decode ResourceStatusRequest IE[%d] value: %w", i, err)
		}
		ies := per.NewDecoder(payload, per.APER)

		switch id {
		case common.ProtocolIEID_NGRANNode1MeasurementID:
			err = msg.NGRANNode1MeasurementID.Decode(ies)
		case common.ProtocolIEID_NGRANNode2MeasurementID:
			v := &MeasurementID{}
			err = v.Decode(ies)
			msg.NGRANNode2MeasurementID = v
		case common.ProtocolIEID_RegistrationRequest:
			err = msg.RegistrationRequest.Decode(ies)
		case common.ProtocolIEID_ReportCharacteristics:
			v := &ReportCharacteristics{}
			err = v.Decode(ies)
			msg.ReportCharacteristics = v
		case common.ProtocolIEID_CellToReport:
			v := &CellToReport{}
			err = v.Decode(ies)
			msg.CellToReport = v
		case common.ProtocolIEID_ReportingPeriodicity:
			v := &ReportingPeriodicity{}
			err = v.Decode(ies)
			msg.ReportingPeriodicity = v
		default:
			if criticality.Value == CriticalityReject {
				return fmt.Errorf("unsupported ResourceStatusRequest IE id=%d with reject criticality", id)
			}
		}
		if err != nil {
			return fmt.Errorf("decode ResourceStatusRequest IE id=%d: %w", id, err)
		}
	}

	if _, err := seq.DecodeExtensionAdditions(); err != nil {
		return fmt.Errorf("decode ResourceStatusRequest extensions: %w", err)
	}

	for _, id := range []int64{
		common.ProtocolIEID_NGRANNode1MeasurementID,
		common.ProtocolIEID_RegistrationRequest,
	} {
		if !seen[id] {
			return fmt.Errorf("missing mandatory ResourceStatusRequest IE id=%d", id)
		}
	}

	return nil
}
