package ies

import (
	"fmt"
	"io"

	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

// PartialUEContextTransfer ::= SEQUENCE {
//	protocolIEs  ProtocolIE-Container {{PartialUEContextTransfer-IEs}},
//	...
// }

type PartialUEContextTransfer struct {
	NewNGRANnodeUEXnAPID    NGRANnodeUEXnAPID
	OldNGRANnodeUEXnAPID    NGRANnodeUEXnAPID
	SDTPartialUEContextInfo SDTPartialUEContextInfo
	PosPartialUEContextInfo *PosPartialUEContextInfo
}

func (msg *PartialUEContextTransfer) toIEs() []XnAPMessageIE {
	ies := make([]XnAPMessageIE, 0)
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_NewNGRANnodeUEXnAPID)},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.NewNGRANnodeUEXnAPID,
	})
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_OldNGRANnodeUEXnAPID)},
		Criticality: Criticality{Value: CriticalityIgnore},
		Value:       &msg.OldNGRANnodeUEXnAPID,
	})
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SDTPartialUEContextInfo)},
		Criticality: Criticality{Value: CriticalityIgnore},
		Value:       &msg.SDTPartialUEContextInfo,
	})
	if msg.PosPartialUEContextInfo != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_PosPartialUEContextInfo)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.PosPartialUEContextInfo,
		})
	}
	return ies
}

func (msg *PartialUEContextTransfer) Encode(w io.Writer) error {
	inner := per.NewEncoder(per.APER)
	if err := encodeBody(msg, inner); err != nil {
		return err
	}

	outer := per.NewEncoder(per.APER)

	if err := choiceXnMsg(outer, common.InitiatingMessage); err != nil {
		return err
	}
	if err := procedureCodeXnMsg(outer, common.ProcedureCode_PartialUEContextTransfer); err != nil {
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

func (msg *PartialUEContextTransfer) Decode(wire []byte) error {
	d := per.NewDecoder(wire, per.APER)

	seq := d.NewSequenceDecoder(per.SequenceConstraints{Extensible: true})
	if err := seq.DecodeExtensionBit(); err != nil {
		return fmt.Errorf("decode PartialUEContextTransfer extension bit: %w", err)
	}

	seqOf := d.NewSequenceOfDecoder(per.SizeConstraints{
		Extensible: false,
		Min:        common.Ptr(int64(0)),
		Max:        common.Ptr(int64(MaxProtocolIEs)),
	})
	n, err := seqOf.DecodeLength()
	if err != nil {
		return fmt.Errorf("decode PartialUEContextTransfer protocolIEs length: %w", err)
	}

	seen := make(map[int64]bool, n)
	for i := range n {
		id, err := d.DecodeInteger(protocolIEIDConstraints)
		if err != nil {
			return fmt.Errorf("decode PartialUEContextTransfer IE[%d] id: %w", i, err)
		}
		seen[id] = true

		var criticality Criticality
		if err := criticality.Decode(d); err != nil {
			return fmt.Errorf("decode PartialUEContextTransfer IE[%d] criticality: %w", i, err)
		}

		payload, err := d.DecodeOpenType()
		if err != nil {
			return fmt.Errorf("decode PartialUEContextTransfer IE[%d] value: %w", i, err)
		}
		ies := per.NewDecoder(payload, per.APER)

		switch id {
		case common.ProtocolIEID_NewNGRANnodeUEXnAPID:
			err = msg.NewNGRANnodeUEXnAPID.Decode(ies)
		case common.ProtocolIEID_OldNGRANnodeUEXnAPID:
			err = msg.OldNGRANnodeUEXnAPID.Decode(ies)
		case common.ProtocolIEID_SDTPartialUEContextInfo:
			err = msg.SDTPartialUEContextInfo.Decode(ies)
		case common.ProtocolIEID_PosPartialUEContextInfo:
			v := &PosPartialUEContextInfo{}
			err = v.Decode(ies)
			msg.PosPartialUEContextInfo = v
		default:
			if criticality.Value == CriticalityReject {
				return fmt.Errorf("unsupported PartialUEContextTransfer IE id=%d with reject criticality", id)
			}
		}
		if err != nil {
			return fmt.Errorf("decode PartialUEContextTransfer IE id=%d: %w", id, err)
		}
	}

	if _, err := seq.DecodeExtensionAdditions(); err != nil {
		return fmt.Errorf("decode PartialUEContextTransfer extensions: %w", err)
	}

	for _, id := range []int64{
		common.ProtocolIEID_NewNGRANnodeUEXnAPID,
		common.ProtocolIEID_OldNGRANnodeUEXnAPID,
		common.ProtocolIEID_SDTPartialUEContextInfo,
	} {
		if !seen[id] {
			return fmt.Errorf("missing mandatory PartialUEContextTransfer IE id=%d", id)
		}
	}

	return nil
}
