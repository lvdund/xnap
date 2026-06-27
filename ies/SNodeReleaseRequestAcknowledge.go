package ies

import (
	"fmt"
	"io"

	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

// SNodeReleaseRequestAcknowledge ::= SEQUENCE {
//	protocolIEs  ProtocolIE-Container {{SNodeReleaseRequestAcknowledge-IEs}},
//	...
// }

type SNodeReleaseRequestAcknowledge struct {
	MNGRANnodeUEXnAPID              NGRANnodeUEXnAPID
	SNGRANnodeUEXnAPID              *NGRANnodeUEXnAPID
	PDUSessionToBeReleasedRelReqAck *PDUSessionToBeReleasedListRelReqAck
	CriticalityDiagnostics          *CriticalityDiagnostics
	SCGUEHistoryInformation         *SCGUEHistoryInformation
	SNMobilityInformation           *SNMobilityInformation
}

func (msg *SNodeReleaseRequestAcknowledge) toIEs() []XnAPMessageIE {
	ies := make([]XnAPMessageIE, 0)
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_MNGRANnodeUEXnAPID)},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.MNGRANnodeUEXnAPID,
	})
	if msg.SNGRANnodeUEXnAPID != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SNGRANnodeUEXnAPID)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.SNGRANnodeUEXnAPID,
		})
	}
	if msg.PDUSessionToBeReleasedRelReqAck != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_PDUSessionToBeReleasedRelReqAck)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.PDUSessionToBeReleasedRelReqAck,
		})
	}
	if msg.CriticalityDiagnostics != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_CriticalityDiagnostics)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.CriticalityDiagnostics,
		})
	}
	if msg.SCGUEHistoryInformation != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SCGUEHistoryInformation)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.SCGUEHistoryInformation,
		})
	}
	if msg.SNMobilityInformation != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SNMobilityInformation)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.SNMobilityInformation,
		})
	}
	return ies
}

func (msg *SNodeReleaseRequestAcknowledge) Encode(w io.Writer) error {
	inner := per.NewEncoder(per.APER)
	if err := encodeBody(msg, inner); err != nil {
		return err
	}

	outer := per.NewEncoder(per.APER)

	if err := choiceXnMsg(outer, common.SuccessfulOutcome); err != nil {
		return err
	}
	if err := procedureCodeXnMsg(outer, common.ProcedureCode_MNGRANnodeinitiatedSNGRANnodeRelease); err != nil {
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

func (msg *SNodeReleaseRequestAcknowledge) Decode(wire []byte) error {
	d := per.NewDecoder(wire, per.APER)

	seq := d.NewSequenceDecoder(per.SequenceConstraints{Extensible: true})
	if err := seq.DecodeExtensionBit(); err != nil {
		return fmt.Errorf("decode SNodeReleaseRequestAcknowledge extension bit: %w", err)
	}

	seqOf := d.NewSequenceOfDecoder(per.SizeConstraints{
		Extensible: false,
		Min:        common.Ptr(int64(0)),
		Max:        common.Ptr(int64(MaxProtocolIEs)),
	})
	n, err := seqOf.DecodeLength()
	if err != nil {
		return fmt.Errorf("decode SNodeReleaseRequestAcknowledge protocolIEs length: %w", err)
	}

	seen := make(map[int64]bool, n)
	for i := range n {
		id, err := d.DecodeInteger(protocolIEIDConstraints)
		if err != nil {
			return fmt.Errorf("decode SNodeReleaseRequestAcknowledge IE[%d] id: %w", i, err)
		}
		seen[id] = true

		var criticality Criticality
		if err := criticality.Decode(d); err != nil {
			return fmt.Errorf("decode SNodeReleaseRequestAcknowledge IE[%d] criticality: %w", i, err)
		}

		payload, err := d.DecodeOpenType()
		if err != nil {
			return fmt.Errorf("decode SNodeReleaseRequestAcknowledge IE[%d] value: %w", i, err)
		}
		ies := per.NewDecoder(payload, per.APER)

		switch id {
		case common.ProtocolIEID_MNGRANnodeUEXnAPID:
			err = msg.MNGRANnodeUEXnAPID.Decode(ies)
		case common.ProtocolIEID_SNGRANnodeUEXnAPID:
			v := &NGRANnodeUEXnAPID{}
			err = v.Decode(ies)
			msg.SNGRANnodeUEXnAPID = v
		case common.ProtocolIEID_PDUSessionToBeReleasedRelReqAck:
			v := &PDUSessionToBeReleasedListRelReqAck{}
			err = v.Decode(ies)
			msg.PDUSessionToBeReleasedRelReqAck = v
		case common.ProtocolIEID_CriticalityDiagnostics:
			v := &CriticalityDiagnostics{}
			err = v.Decode(ies)
			msg.CriticalityDiagnostics = v
		case common.ProtocolIEID_SCGUEHistoryInformation:
			v := &SCGUEHistoryInformation{}
			err = v.Decode(ies)
			msg.SCGUEHistoryInformation = v
		case common.ProtocolIEID_SNMobilityInformation:
			v := &SNMobilityInformation{}
			err = v.Decode(ies)
			msg.SNMobilityInformation = v
		default:
			if criticality.Value == CriticalityReject {
				return fmt.Errorf("unsupported SNodeReleaseRequestAcknowledge IE id=%d with reject criticality", id)
			}
		}
		if err != nil {
			return fmt.Errorf("decode SNodeReleaseRequestAcknowledge IE id=%d: %w", id, err)
		}
	}

	if _, err := seq.DecodeExtensionAdditions(); err != nil {
		return fmt.Errorf("decode SNodeReleaseRequestAcknowledge extensions: %w", err)
	}

	for _, id := range []int64{
		common.ProtocolIEID_MNGRANnodeUEXnAPID,
	} {
		if !seen[id] {
			return fmt.Errorf("missing mandatory SNodeReleaseRequestAcknowledge IE id=%d", id)
		}
	}

	return nil
}
