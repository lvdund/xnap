package ies

import (
	"fmt"
	"io"

	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

// HandoverReport ::= SEQUENCE {
//	protocolIEs  ProtocolIE-Container {{HandoverReport-IEs}},
//	...
// }

type HandoverReport struct {
	HandoverReportType     HandoverReportType
	HandoverCause          Cause
	SourceCellCGI          GlobalNGRANCellID
	TargetCellCGI          GlobalNGRANCellID
	ReEstablishmentCellCGI *GlobalCellID
	TargetCellinEUTRAN     *TargetCellinEUTRAN
	SourceCellCRNTI        *CRNTI
	MobilityInformation    *MobilityInformation
	UERLFReportContainer   *UERLFReportContainer
	CHOConfiguration       *CHOConfiguration
	TargetCellCRNTI        *CRNTI
	TimeSinceFailure       *TimeSinceFailure
}

func (msg *HandoverReport) toIEs() []XnAPMessageIE {
	ies := make([]XnAPMessageIE, 0)
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_HandoverReportType)},
		Criticality: Criticality{Value: CriticalityIgnore},
		Value:       &msg.HandoverReportType,
	})
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_HandoverCause)},
		Criticality: Criticality{Value: CriticalityIgnore},
		Value:       &msg.HandoverCause,
	})
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SourceCellCGI)},
		Criticality: Criticality{Value: CriticalityIgnore},
		Value:       &msg.SourceCellCGI,
	})
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_TargetCellCGI)},
		Criticality: Criticality{Value: CriticalityIgnore},
		Value:       &msg.TargetCellCGI,
	})
	if msg.ReEstablishmentCellCGI != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_ReEstablishmentCellCGI)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.ReEstablishmentCellCGI,
		})
	}
	if msg.TargetCellinEUTRAN != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_TargetCellinEUTRAN)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.TargetCellinEUTRAN,
		})
	}
	if msg.SourceCellCRNTI != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SourceCellCRNTI)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.SourceCellCRNTI,
		})
	}
	if msg.MobilityInformation != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_MobilityInformation)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.MobilityInformation,
		})
	}
	if msg.UERLFReportContainer != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_UERLFReportContainer)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.UERLFReportContainer,
		})
	}
	if msg.CHOConfiguration != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_CHOConfiguration)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.CHOConfiguration,
		})
	}
	if msg.TargetCellCRNTI != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_TargetCellCRNTI)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.TargetCellCRNTI,
		})
	}
	if msg.TimeSinceFailure != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_TimeSinceFailure)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.TimeSinceFailure,
		})
	}
	return ies
}

func (msg *HandoverReport) Encode(w io.Writer) error {
	inner := per.NewEncoder(per.APER)
	if err := encodeBody(msg, inner); err != nil {
		return err
	}

	outer := per.NewEncoder(per.APER)

	if err := choiceXnMsg(outer, common.InitiatingMessage); err != nil {
		return err
	}
	if err := procedureCodeXnMsg(outer, common.ProcedureCode_HandoverReport); err != nil {
		return err
	}
	if err := criticalityXnMsg(outer, CriticalityIgnore); err != nil {
		return err
	}

	if err := outer.EncodeOpenType(inner.Bytes()); err != nil {
		return err
	}

	_, err := w.Write(outer.Bytes())
	return err
}

func (msg *HandoverReport) Decode(wire []byte) error {
	d := per.NewDecoder(wire, per.APER)

	seq := d.NewSequenceDecoder(per.SequenceConstraints{Extensible: true})
	if err := seq.DecodeExtensionBit(); err != nil {
		return fmt.Errorf("decode HandoverReport extension bit: %w", err)
	}

	seqOf := d.NewSequenceOfDecoder(per.SizeConstraints{
		Extensible: false,
		Min:        common.Ptr(int64(0)),
		Max:        common.Ptr(int64(MaxProtocolIEs)),
	})
	n, err := seqOf.DecodeLength()
	if err != nil {
		return fmt.Errorf("decode HandoverReport protocolIEs length: %w", err)
	}

	seen := make(map[int64]bool, n)
	for i := range n {
		id, err := d.DecodeInteger(protocolIEIDConstraints)
		if err != nil {
			return fmt.Errorf("decode HandoverReport IE[%d] id: %w", i, err)
		}
		seen[id] = true

		var criticality Criticality
		if err := criticality.Decode(d); err != nil {
			return fmt.Errorf("decode HandoverReport IE[%d] criticality: %w", i, err)
		}

		payload, err := d.DecodeOpenType()
		if err != nil {
			return fmt.Errorf("decode HandoverReport IE[%d] value: %w", i, err)
		}
		ies := per.NewDecoder(payload, per.APER)

		switch id {
		case common.ProtocolIEID_HandoverReportType:
			err = msg.HandoverReportType.Decode(ies)
		case common.ProtocolIEID_HandoverCause:
			err = msg.HandoverCause.Decode(ies)
		case common.ProtocolIEID_SourceCellCGI:
			err = msg.SourceCellCGI.Decode(ies)
		case common.ProtocolIEID_TargetCellCGI:
			err = msg.TargetCellCGI.Decode(ies)
		case common.ProtocolIEID_ReEstablishmentCellCGI:
			v := &GlobalCellID{}
			err = v.Decode(ies)
			msg.ReEstablishmentCellCGI = v
		case common.ProtocolIEID_TargetCellinEUTRAN:
			v := &TargetCellinEUTRAN{}
			err = v.Decode(ies)
			msg.TargetCellinEUTRAN = v
		case common.ProtocolIEID_SourceCellCRNTI:
			v := &CRNTI{}
			err = v.Decode(ies)
			msg.SourceCellCRNTI = v
		case common.ProtocolIEID_MobilityInformation:
			v := &MobilityInformation{}
			err = v.Decode(ies)
			msg.MobilityInformation = v
		case common.ProtocolIEID_UERLFReportContainer:
			v := &UERLFReportContainer{}
			err = v.Decode(ies)
			msg.UERLFReportContainer = v
		case common.ProtocolIEID_CHOConfiguration:
			v := &CHOConfiguration{}
			err = v.Decode(ies)
			msg.CHOConfiguration = v
		case common.ProtocolIEID_TargetCellCRNTI:
			v := &CRNTI{}
			err = v.Decode(ies)
			msg.TargetCellCRNTI = v
		case common.ProtocolIEID_TimeSinceFailure:
			v := &TimeSinceFailure{}
			err = v.Decode(ies)
			msg.TimeSinceFailure = v
		default:
			if criticality.Value == CriticalityReject {
				return fmt.Errorf("unsupported HandoverReport IE id=%d with reject criticality", id)
			}
		}
		if err != nil {
			return fmt.Errorf("decode HandoverReport IE id=%d: %w", id, err)
		}
	}

	if _, err := seq.DecodeExtensionAdditions(); err != nil {
		return fmt.Errorf("decode HandoverReport extensions: %w", err)
	}

	for _, id := range []int64{
		common.ProtocolIEID_HandoverReportType,
		common.ProtocolIEID_HandoverCause,
		common.ProtocolIEID_SourceCellCGI,
		common.ProtocolIEID_TargetCellCGI,
	} {
		if !seen[id] {
			return fmt.Errorf("missing mandatory HandoverReport IE id=%d", id)
		}
	}

	return nil
}
