package ies

import (
	"fmt"
	"io"

	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

// ConditionalHandoverCancel ::= SEQUENCE {
//	protocolIEs  ProtocolIE-Container {{ConditionalHandoverCancel-IEs}},
//	...
// }

type ConditionalHandoverCancel struct {
	SourceNGRANnodeUEXnAPID         NGRANnodeUEXnAPID
	TargetNGRANnodeUEXnAPID         NGRANnodeUEXnAPID
	Cause                           Cause
	TargetCellsToCancel             *TargetCellList
	ConditionalReconfigToCancelList *ConditionalReconfigList
}

func (msg *ConditionalHandoverCancel) toIEs() []XnAPMessageIE {
	ies := make([]XnAPMessageIE, 0)
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SourceNGRANnodeUEXnAPID)},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.SourceNGRANnodeUEXnAPID,
	})
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_TargetNGRANnodeUEXnAPID)},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.TargetNGRANnodeUEXnAPID,
	})
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_Cause)},
		Criticality: Criticality{Value: CriticalityIgnore},
		Value:       &msg.Cause,
	})
	if msg.TargetCellsToCancel != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_TargetCellsToCancel)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.TargetCellsToCancel,
		})
	}
	if msg.ConditionalReconfigToCancelList != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_ConditionalReconfigToCancelList)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.ConditionalReconfigToCancelList,
		})
	}
	return ies
}

func (msg *ConditionalHandoverCancel) Encode(w io.Writer) error {
	inner := per.NewEncoder(per.APER)
	if err := encodeBody(msg, inner); err != nil {
		return err
	}

	outer := per.NewEncoder(per.APER)

	if err := choiceXnMsg(outer, common.InitiatingMessage); err != nil {
		return err
	}
	if err := procedureCodeXnMsg(outer, common.ProcedureCode_ConditionalHandoverCancel); err != nil {
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

func (msg *ConditionalHandoverCancel) Decode(wire []byte) error {
	d := per.NewDecoder(wire, per.APER)

	seq := d.NewSequenceDecoder(per.SequenceConstraints{Extensible: true})
	if err := seq.DecodeExtensionBit(); err != nil {
		return fmt.Errorf("decode ConditionalHandoverCancel extension bit: %w", err)
	}

	seqOf := d.NewSequenceOfDecoder(per.SizeConstraints{
		Extensible: false,
		Min:        common.Ptr(int64(0)),
		Max:        common.Ptr(int64(MaxProtocolIEs)),
	})
	n, err := seqOf.DecodeLength()
	if err != nil {
		return fmt.Errorf("decode ConditionalHandoverCancel protocolIEs length: %w", err)
	}

	seen := make(map[int64]bool, n)
	for i := range n {
		id, err := d.DecodeInteger(protocolIEIDConstraints)
		if err != nil {
			return fmt.Errorf("decode ConditionalHandoverCancel IE[%d] id: %w", i, err)
		}
		seen[id] = true

		var criticality Criticality
		if err := criticality.Decode(d); err != nil {
			return fmt.Errorf("decode ConditionalHandoverCancel IE[%d] criticality: %w", i, err)
		}

		payload, err := d.DecodeOpenType()
		if err != nil {
			return fmt.Errorf("decode ConditionalHandoverCancel IE[%d] value: %w", i, err)
		}
		ies := per.NewDecoder(payload, per.APER)

		switch id {
		case common.ProtocolIEID_SourceNGRANnodeUEXnAPID:
			err = msg.SourceNGRANnodeUEXnAPID.Decode(ies)
		case common.ProtocolIEID_TargetNGRANnodeUEXnAPID:
			err = msg.TargetNGRANnodeUEXnAPID.Decode(ies)
		case common.ProtocolIEID_Cause:
			err = msg.Cause.Decode(ies)
		case common.ProtocolIEID_TargetCellsToCancel:
			v := &TargetCellList{}
			err = v.Decode(ies)
			msg.TargetCellsToCancel = v
		case common.ProtocolIEID_ConditionalReconfigToCancelList:
			v := &ConditionalReconfigList{}
			err = v.Decode(ies)
			msg.ConditionalReconfigToCancelList = v
		default:
			if criticality.Value == CriticalityReject {
				return fmt.Errorf("unsupported ConditionalHandoverCancel IE id=%d with reject criticality", id)
			}
		}
		if err != nil {
			return fmt.Errorf("decode ConditionalHandoverCancel IE id=%d: %w", id, err)
		}
	}

	if _, err := seq.DecodeExtensionAdditions(); err != nil {
		return fmt.Errorf("decode ConditionalHandoverCancel extensions: %w", err)
	}

	for _, id := range []int64{
		common.ProtocolIEID_SourceNGRANnodeUEXnAPID,
		common.ProtocolIEID_TargetNGRANnodeUEXnAPID,
		common.ProtocolIEID_Cause,
	} {
		if !seen[id] {
			return fmt.Errorf("missing mandatory ConditionalHandoverCancel IE id=%d", id)
		}
	}

	return nil
}
