package ies

import (
	"fmt"
	"io"

	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

// HandoverPreparationFailure ::= SEQUENCE {
//	protocolIEs  ProtocolIE-Container {{HandoverPreparationFailure-IEs}},
//	...
// }

type HandoverPreparationFailure struct {
	SourceNGRANnodeUEXnAPID     NGRANnodeUEXnAPID
	Cause                       Cause
	CriticalityDiagnostics      *CriticalityDiagnostics
	RequestedTargetCellGlobalID *TargetCGI
}

func (msg *HandoverPreparationFailure) toIEs() []XnAPMessageIE {
	ies := make([]XnAPMessageIE, 0)
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SourceNGRANnodeUEXnAPID)},
		Criticality: Criticality{Value: CriticalityIgnore},
		Value:       &msg.SourceNGRANnodeUEXnAPID,
	})
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_Cause)},
		Criticality: Criticality{Value: CriticalityIgnore},
		Value:       &msg.Cause,
	})
	if msg.CriticalityDiagnostics != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_CriticalityDiagnostics)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.CriticalityDiagnostics,
		})
	}
	if msg.RequestedTargetCellGlobalID != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_RequestedTargetCellGlobalID)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.RequestedTargetCellGlobalID,
		})
	}
	return ies
}

func (msg *HandoverPreparationFailure) Encode(w io.Writer) error {
	inner := per.NewEncoder(per.APER)
	if err := encodeBody(msg, inner); err != nil {
		return err
	}

	outer := per.NewEncoder(per.APER)

	if err := choiceXnMsg(outer, common.UnsuccessfulOutcome); err != nil {
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

func (msg *HandoverPreparationFailure) Decode(wire []byte) error {
	d := per.NewDecoder(wire, per.APER)

	seq := d.NewSequenceDecoder(per.SequenceConstraints{Extensible: true})
	if err := seq.DecodeExtensionBit(); err != nil {
		return fmt.Errorf("decode HandoverPreparationFailure extension bit: %w", err)
	}

	seqOf := d.NewSequenceOfDecoder(per.SizeConstraints{
		Extensible: false,
		Min:        common.Ptr(int64(0)),
		Max:        common.Ptr(int64(MaxProtocolIEs)),
	})
	n, err := seqOf.DecodeLength()
	if err != nil {
		return fmt.Errorf("decode HandoverPreparationFailure protocolIEs length: %w", err)
	}

	seen := make(map[int64]bool, n)
	for i := range n {
		id, err := d.DecodeInteger(protocolIEIDConstraints)
		if err != nil {
			return fmt.Errorf("decode HandoverPreparationFailure IE[%d] id: %w", i, err)
		}
		seen[id] = true

		var criticality Criticality
		if err := criticality.Decode(d); err != nil {
			return fmt.Errorf("decode HandoverPreparationFailure IE[%d] criticality: %w", i, err)
		}

		payload, err := d.DecodeOpenType()
		if err != nil {
			return fmt.Errorf("decode HandoverPreparationFailure IE[%d] value: %w", i, err)
		}
		ies := per.NewDecoder(payload, per.APER)

		switch id {
		case common.ProtocolIEID_SourceNGRANnodeUEXnAPID:
			err = msg.SourceNGRANnodeUEXnAPID.Decode(ies)
		case common.ProtocolIEID_Cause:
			err = msg.Cause.Decode(ies)
		case common.ProtocolIEID_CriticalityDiagnostics:
			v := &CriticalityDiagnostics{}
			err = v.Decode(ies)
			msg.CriticalityDiagnostics = v
		case common.ProtocolIEID_RequestedTargetCellGlobalID:
			v := &TargetCGI{}
			err = v.Decode(ies)
			msg.RequestedTargetCellGlobalID = v
		default:
			if criticality.Value == CriticalityReject {
				return fmt.Errorf("unsupported HandoverPreparationFailure IE id=%d with reject criticality", id)
			}
		}
		if err != nil {
			return fmt.Errorf("decode HandoverPreparationFailure IE id=%d: %w", id, err)
		}
	}

	if _, err := seq.DecodeExtensionAdditions(); err != nil {
		return fmt.Errorf("decode HandoverPreparationFailure extensions: %w", err)
	}

	for _, id := range []int64{
		common.ProtocolIEID_SourceNGRANnodeUEXnAPID,
		common.ProtocolIEID_Cause,
	} {
		if !seen[id] {
			return fmt.Errorf("missing mandatory HandoverPreparationFailure IE id=%d", id)
		}
	}

	return nil
}
