package ies

import (
	"fmt"
	"io"

	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

// IABResourceCoordinationRequest ::= SEQUENCE {
//	protocolIEs  ProtocolIE-Container {{IABResourceCoordinationRequest-IEs}},
//	...
// }

type IABResourceCoordinationRequest struct {
	F1TerminatingIABDonorUEXnAPID    NGRANnodeUEXnAPID
	NonF1TerminatingIABDonorUEXnAPID NGRANnodeUEXnAPID
	BoundaryNodeCellsList            *BoundaryNodeCellsList
	ParentNodeCellsList              *ParentNodeCellsList
}

func (msg *IABResourceCoordinationRequest) toIEs() []XnAPMessageIE {
	ies := make([]XnAPMessageIE, 0)
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_F1TerminatingIABDonorUEXnAPID)},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.F1TerminatingIABDonorUEXnAPID,
	})
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_NonF1TerminatingIABDonorUEXnAPID)},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.NonF1TerminatingIABDonorUEXnAPID,
	})
	if msg.BoundaryNodeCellsList != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_BoundaryNodeCellsList)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.BoundaryNodeCellsList,
		})
	}
	if msg.ParentNodeCellsList != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_ParentNodeCellsList)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.ParentNodeCellsList,
		})
	}
	return ies
}

func (msg *IABResourceCoordinationRequest) Encode(w io.Writer) error {
	inner := per.NewEncoder(per.APER)
	if err := encodeBody(msg, inner); err != nil {
		return err
	}

	outer := per.NewEncoder(per.APER)

	if err := choiceXnMsg(outer, common.InitiatingMessage); err != nil {
		return err
	}
	if err := procedureCodeXnMsg(outer, common.ProcedureCode_IABResourceCoordination); err != nil {
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

func (msg *IABResourceCoordinationRequest) Decode(wire []byte) error {
	d := per.NewDecoder(wire, per.APER)

	seq := d.NewSequenceDecoder(per.SequenceConstraints{Extensible: true})
	if err := seq.DecodeExtensionBit(); err != nil {
		return fmt.Errorf("decode IABResourceCoordinationRequest extension bit: %w", err)
	}

	seqOf := d.NewSequenceOfDecoder(per.SizeConstraints{
		Extensible: false,
		Min:        common.Ptr(int64(0)),
		Max:        common.Ptr(int64(MaxProtocolIEs)),
	})
	n, err := seqOf.DecodeLength()
	if err != nil {
		return fmt.Errorf("decode IABResourceCoordinationRequest protocolIEs length: %w", err)
	}

	seen := make(map[int64]bool, n)
	for i := range n {
		id, err := d.DecodeInteger(protocolIEIDConstraints)
		if err != nil {
			return fmt.Errorf("decode IABResourceCoordinationRequest IE[%d] id: %w", i, err)
		}
		seen[id] = true

		var criticality Criticality
		if err := criticality.Decode(d); err != nil {
			return fmt.Errorf("decode IABResourceCoordinationRequest IE[%d] criticality: %w", i, err)
		}

		payload, err := d.DecodeOpenType()
		if err != nil {
			return fmt.Errorf("decode IABResourceCoordinationRequest IE[%d] value: %w", i, err)
		}
		ies := per.NewDecoder(payload, per.APER)

		switch id {
		case common.ProtocolIEID_F1TerminatingIABDonorUEXnAPID:
			err = msg.F1TerminatingIABDonorUEXnAPID.Decode(ies)
		case common.ProtocolIEID_NonF1TerminatingIABDonorUEXnAPID:
			err = msg.NonF1TerminatingIABDonorUEXnAPID.Decode(ies)
		case common.ProtocolIEID_BoundaryNodeCellsList:
			v := &BoundaryNodeCellsList{}
			err = v.Decode(ies)
			msg.BoundaryNodeCellsList = v
		case common.ProtocolIEID_ParentNodeCellsList:
			v := &ParentNodeCellsList{}
			err = v.Decode(ies)
			msg.ParentNodeCellsList = v
		default:
			if criticality.Value == CriticalityReject {
				return fmt.Errorf("unsupported IABResourceCoordinationRequest IE id=%d with reject criticality", id)
			}
		}
		if err != nil {
			return fmt.Errorf("decode IABResourceCoordinationRequest IE id=%d: %w", id, err)
		}
	}

	if _, err := seq.DecodeExtensionAdditions(); err != nil {
		return fmt.Errorf("decode IABResourceCoordinationRequest extensions: %w", err)
	}

	for _, id := range []int64{
		common.ProtocolIEID_F1TerminatingIABDonorUEXnAPID,
		common.ProtocolIEID_NonF1TerminatingIABDonorUEXnAPID,
	} {
		if !seen[id] {
			return fmt.Errorf("missing mandatory IABResourceCoordinationRequest IE id=%d", id)
		}
	}

	return nil
}
