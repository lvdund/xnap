package ies

import (
	"fmt"
	"io"

	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

// HandoverSuccess ::= SEQUENCE {
//	protocolIEs  ProtocolIE-Container {{HandoverSuccess-IEs}},
//	...
// }

type HandoverSuccess struct {
	SourceNGRANnodeUEXnAPID     NGRANnodeUEXnAPID
	TargetNGRANnodeUEXnAPID     NGRANnodeUEXnAPID
	RequestedTargetCellGlobalID TargetCGI
	AccessedPSCellID            *NRCGI
}

func (msg *HandoverSuccess) toIEs() []XnAPMessageIE {
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
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_RequestedTargetCellGlobalID)},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.RequestedTargetCellGlobalID,
	})
	if msg.AccessedPSCellID != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_AccessedPSCellID)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.AccessedPSCellID,
		})
	}
	return ies
}

func (msg *HandoverSuccess) Encode(w io.Writer) error {
	inner := per.NewEncoder(per.APER)
	if err := encodeBody(msg, inner); err != nil {
		return err
	}

	outer := per.NewEncoder(per.APER)

	if err := choiceXnMsg(outer, common.InitiatingMessage); err != nil {
		return err
	}
	if err := procedureCodeXnMsg(outer, common.ProcedureCode_HandoverSuccess); err != nil {
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

func (msg *HandoverSuccess) Decode(wire []byte) error {
	d := per.NewDecoder(wire, per.APER)

	seq := d.NewSequenceDecoder(per.SequenceConstraints{Extensible: true})
	if err := seq.DecodeExtensionBit(); err != nil {
		return fmt.Errorf("decode HandoverSuccess extension bit: %w", err)
	}

	seqOf := d.NewSequenceOfDecoder(per.SizeConstraints{
		Extensible: false,
		Min:        common.Ptr(int64(0)),
		Max:        common.Ptr(int64(MaxProtocolIEs)),
	})
	n, err := seqOf.DecodeLength()
	if err != nil {
		return fmt.Errorf("decode HandoverSuccess protocolIEs length: %w", err)
	}

	seen := make(map[int64]bool, n)
	for i := range n {
		id, err := d.DecodeInteger(protocolIEIDConstraints)
		if err != nil {
			return fmt.Errorf("decode HandoverSuccess IE[%d] id: %w", i, err)
		}
		seen[id] = true

		var criticality Criticality
		if err := criticality.Decode(d); err != nil {
			return fmt.Errorf("decode HandoverSuccess IE[%d] criticality: %w", i, err)
		}

		payload, err := d.DecodeOpenType()
		if err != nil {
			return fmt.Errorf("decode HandoverSuccess IE[%d] value: %w", i, err)
		}
		ies := per.NewDecoder(payload, per.APER)

		switch id {
		case common.ProtocolIEID_SourceNGRANnodeUEXnAPID:
			err = msg.SourceNGRANnodeUEXnAPID.Decode(ies)
		case common.ProtocolIEID_TargetNGRANnodeUEXnAPID:
			err = msg.TargetNGRANnodeUEXnAPID.Decode(ies)
		case common.ProtocolIEID_RequestedTargetCellGlobalID:
			err = msg.RequestedTargetCellGlobalID.Decode(ies)
		case common.ProtocolIEID_AccessedPSCellID:
			v := &NRCGI{}
			err = v.Decode(ies)
			msg.AccessedPSCellID = v
		default:
			if criticality.Value == CriticalityReject {
				return fmt.Errorf("unsupported HandoverSuccess IE id=%d with reject criticality", id)
			}
		}
		if err != nil {
			return fmt.Errorf("decode HandoverSuccess IE id=%d: %w", id, err)
		}
	}

	if _, err := seq.DecodeExtensionAdditions(); err != nil {
		return fmt.Errorf("decode HandoverSuccess extensions: %w", err)
	}

	for _, id := range []int64{
		common.ProtocolIEID_SourceNGRANnodeUEXnAPID,
		common.ProtocolIEID_TargetNGRANnodeUEXnAPID,
		common.ProtocolIEID_RequestedTargetCellGlobalID,
	} {
		if !seen[id] {
			return fmt.Errorf("missing mandatory HandoverSuccess IE id=%d", id)
		}
	}

	return nil
}
