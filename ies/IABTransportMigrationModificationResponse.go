package ies

import (
	"fmt"
	"io"

	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

// IABTransportMigrationModificationResponse ::= SEQUENCE {
//	protocolIEs  ProtocolIE-Container {{IABTransportMigrationModificationResponse-IEs}},
//	...
// }

type IABTransportMigrationModificationResponse struct {
	F1TerminatingIABDonorUEXnAPID    NGRANnodeUEXnAPID
	NonF1TerminatingIABDonorUEXnAPID NGRANnodeUEXnAPID
	TrafficRequiredModifiedList      *TrafficRequiredModifiedList
	TrafficReleasedList              *TrafficReleasedList
}

func (msg *IABTransportMigrationModificationResponse) toIEs() []XnAPMessageIE {
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
	if msg.TrafficRequiredModifiedList != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_TrafficRequiredModifiedList)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.TrafficRequiredModifiedList,
		})
	}
	if msg.TrafficReleasedList != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_TrafficReleasedList)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.TrafficReleasedList,
		})
	}
	return ies
}

func (msg *IABTransportMigrationModificationResponse) Encode(w io.Writer) error {
	inner := per.NewEncoder(per.APER)
	if err := encodeBody(msg, inner); err != nil {
		return err
	}

	outer := per.NewEncoder(per.APER)

	if err := choiceXnMsg(outer, common.SuccessfulOutcome); err != nil {
		return err
	}
	if err := procedureCodeXnMsg(outer, common.ProcedureCode_IABTransportMigrationModification); err != nil {
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

func (msg *IABTransportMigrationModificationResponse) Decode(wire []byte) error {
	d := per.NewDecoder(wire, per.APER)

	seq := d.NewSequenceDecoder(per.SequenceConstraints{Extensible: true})
	if err := seq.DecodeExtensionBit(); err != nil {
		return fmt.Errorf("decode IABTransportMigrationModificationResponse extension bit: %w", err)
	}

	seqOf := d.NewSequenceOfDecoder(per.SizeConstraints{
		Extensible: false,
		Min:        common.Ptr(int64(0)),
		Max:        common.Ptr(int64(MaxProtocolIEs)),
	})
	n, err := seqOf.DecodeLength()
	if err != nil {
		return fmt.Errorf("decode IABTransportMigrationModificationResponse protocolIEs length: %w", err)
	}

	seen := make(map[int64]bool, n)
	for i := range n {
		id, err := d.DecodeInteger(protocolIEIDConstraints)
		if err != nil {
			return fmt.Errorf("decode IABTransportMigrationModificationResponse IE[%d] id: %w", i, err)
		}
		seen[id] = true

		var criticality Criticality
		if err := criticality.Decode(d); err != nil {
			return fmt.Errorf("decode IABTransportMigrationModificationResponse IE[%d] criticality: %w", i, err)
		}

		payload, err := d.DecodeOpenType()
		if err != nil {
			return fmt.Errorf("decode IABTransportMigrationModificationResponse IE[%d] value: %w", i, err)
		}
		ies := per.NewDecoder(payload, per.APER)

		switch id {
		case common.ProtocolIEID_F1TerminatingIABDonorUEXnAPID:
			err = msg.F1TerminatingIABDonorUEXnAPID.Decode(ies)
		case common.ProtocolIEID_NonF1TerminatingIABDonorUEXnAPID:
			err = msg.NonF1TerminatingIABDonorUEXnAPID.Decode(ies)
		case common.ProtocolIEID_TrafficRequiredModifiedList:
			v := &TrafficRequiredModifiedList{}
			err = v.Decode(ies)
			msg.TrafficRequiredModifiedList = v
		case common.ProtocolIEID_TrafficReleasedList:
			v := &TrafficReleasedList{}
			err = v.Decode(ies)
			msg.TrafficReleasedList = v
		default:
			if criticality.Value == CriticalityReject {
				return fmt.Errorf("unsupported IABTransportMigrationModificationResponse IE id=%d with reject criticality", id)
			}
		}
		if err != nil {
			return fmt.Errorf("decode IABTransportMigrationModificationResponse IE id=%d: %w", id, err)
		}
	}

	if _, err := seq.DecodeExtensionAdditions(); err != nil {
		return fmt.Errorf("decode IABTransportMigrationModificationResponse extensions: %w", err)
	}

	for _, id := range []int64{
		common.ProtocolIEID_F1TerminatingIABDonorUEXnAPID,
		common.ProtocolIEID_NonF1TerminatingIABDonorUEXnAPID,
	} {
		if !seen[id] {
			return fmt.Errorf("missing mandatory IABTransportMigrationModificationResponse IE id=%d", id)
		}
	}

	return nil
}
