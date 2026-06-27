package ies

import (
	"fmt"
	"io"

	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

// ScgFailureInformationReport ::= SEQUENCE {
//	protocolIEs  ProtocolIE-Container {{ScgFailureInformationReport-IEs}},
//	...
// }

type ScgFailureInformationReport struct {
	MNGRANnodeUEXnAPID        NGRANnodeUEXnAPID
	SNGRANnodeUEXnAPID        NGRANnodeUEXnAPID
	SourcePSCellCGI           *GlobalNGRANCellID
	FailedPSCellCGI           *GlobalNGRANCellID
	SCGFailureReportContainer SCGFailureReportContainer
	SNMobilityInformation     *SNMobilityInformation
	CPACConfiguration         *CPACConfiguration
}

func (msg *ScgFailureInformationReport) toIEs() []XnAPMessageIE {
	ies := make([]XnAPMessageIE, 0)
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_MNGRANnodeUEXnAPID)},
		Criticality: Criticality{Value: CriticalityIgnore},
		Value:       &msg.MNGRANnodeUEXnAPID,
	})
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SNGRANnodeUEXnAPID)},
		Criticality: Criticality{Value: CriticalityIgnore},
		Value:       &msg.SNGRANnodeUEXnAPID,
	})
	if msg.SourcePSCellCGI != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SourcePSCellCGI)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.SourcePSCellCGI,
		})
	}
	if msg.FailedPSCellCGI != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_FailedPSCellCGI)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.FailedPSCellCGI,
		})
	}
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SCGFailureReportContainer)},
		Criticality: Criticality{Value: CriticalityIgnore},
		Value:       &msg.SCGFailureReportContainer,
	})
	if msg.SNMobilityInformation != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SNMobilityInformation)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.SNMobilityInformation,
		})
	}
	if msg.CPACConfiguration != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_CPACConfiguration)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.CPACConfiguration,
		})
	}
	return ies
}

func (msg *ScgFailureInformationReport) Encode(w io.Writer) error {
	inner := per.NewEncoder(per.APER)
	if err := encodeBody(msg, inner); err != nil {
		return err
	}

	outer := per.NewEncoder(per.APER)

	if err := choiceXnMsg(outer, common.InitiatingMessage); err != nil {
		return err
	}
	if err := procedureCodeXnMsg(outer, common.ProcedureCode_ScgFailureInformationReport); err != nil {
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

func (msg *ScgFailureInformationReport) Decode(wire []byte) error {
	d := per.NewDecoder(wire, per.APER)

	seq := d.NewSequenceDecoder(per.SequenceConstraints{Extensible: true})
	if err := seq.DecodeExtensionBit(); err != nil {
		return fmt.Errorf("decode ScgFailureInformationReport extension bit: %w", err)
	}

	seqOf := d.NewSequenceOfDecoder(per.SizeConstraints{
		Extensible: false,
		Min:        common.Ptr(int64(0)),
		Max:        common.Ptr(int64(MaxProtocolIEs)),
	})
	n, err := seqOf.DecodeLength()
	if err != nil {
		return fmt.Errorf("decode ScgFailureInformationReport protocolIEs length: %w", err)
	}

	seen := make(map[int64]bool, n)
	for i := range n {
		id, err := d.DecodeInteger(protocolIEIDConstraints)
		if err != nil {
			return fmt.Errorf("decode ScgFailureInformationReport IE[%d] id: %w", i, err)
		}
		seen[id] = true

		var criticality Criticality
		if err := criticality.Decode(d); err != nil {
			return fmt.Errorf("decode ScgFailureInformationReport IE[%d] criticality: %w", i, err)
		}

		payload, err := d.DecodeOpenType()
		if err != nil {
			return fmt.Errorf("decode ScgFailureInformationReport IE[%d] value: %w", i, err)
		}
		ies := per.NewDecoder(payload, per.APER)

		switch id {
		case common.ProtocolIEID_MNGRANnodeUEXnAPID:
			err = msg.MNGRANnodeUEXnAPID.Decode(ies)
		case common.ProtocolIEID_SNGRANnodeUEXnAPID:
			err = msg.SNGRANnodeUEXnAPID.Decode(ies)
		case common.ProtocolIEID_SourcePSCellCGI:
			v := &GlobalNGRANCellID{}
			err = v.Decode(ies)
			msg.SourcePSCellCGI = v
		case common.ProtocolIEID_FailedPSCellCGI:
			v := &GlobalNGRANCellID{}
			err = v.Decode(ies)
			msg.FailedPSCellCGI = v
		case common.ProtocolIEID_SCGFailureReportContainer:
			err = msg.SCGFailureReportContainer.Decode(ies)
		case common.ProtocolIEID_SNMobilityInformation:
			v := &SNMobilityInformation{}
			err = v.Decode(ies)
			msg.SNMobilityInformation = v
		case common.ProtocolIEID_CPACConfiguration:
			v := &CPACConfiguration{}
			err = v.Decode(ies)
			msg.CPACConfiguration = v
		default:
			if criticality.Value == CriticalityReject {
				return fmt.Errorf("unsupported ScgFailureInformationReport IE id=%d with reject criticality", id)
			}
		}
		if err != nil {
			return fmt.Errorf("decode ScgFailureInformationReport IE id=%d: %w", id, err)
		}
	}

	if _, err := seq.DecodeExtensionAdditions(); err != nil {
		return fmt.Errorf("decode ScgFailureInformationReport extensions: %w", err)
	}

	for _, id := range []int64{
		common.ProtocolIEID_MNGRANnodeUEXnAPID,
		common.ProtocolIEID_SNGRANnodeUEXnAPID,
		common.ProtocolIEID_SCGFailureReportContainer,
	} {
		if !seen[id] {
			return fmt.Errorf("missing mandatory ScgFailureInformationReport IE id=%d", id)
		}
	}

	return nil
}
