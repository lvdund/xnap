package ies

import (
	"fmt"
	"io"

	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

// SNodeReleaseRequired ::= SEQUENCE {
//	protocolIEs  ProtocolIE-Container {{SNodeReleaseRequired-IEs}},
//	...
// }

type SNodeReleaseRequired struct {
	MNGRANnodeUEXnAPID                 NGRANnodeUEXnAPID
	SNGRANnodeUEXnAPID                 NGRANnodeUEXnAPID
	PDUSessionToBeReleasedListRelRqd   *PDUSessionToBeReleasedListRelRqd
	Cause                              Cause
	SNToMNContainer                    []byte
	SCGUEHistoryInformation            *SCGUEHistoryInformation
	PDUSessionsListToBeReleasedUPError *PDUSessionsListToBeReleasedUPError
}

func (msg *SNodeReleaseRequired) toIEs() []XnAPMessageIE {
	ies := make([]XnAPMessageIE, 0)
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_MNGRANnodeUEXnAPID)},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.MNGRANnodeUEXnAPID,
	})
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SNGRANnodeUEXnAPID)},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.SNGRANnodeUEXnAPID,
	})
	if msg.PDUSessionToBeReleasedListRelRqd != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_PDUSessionToBeReleasedListRelRqd)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.PDUSessionToBeReleasedListRelRqd,
		})
	}
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_Cause)},
		Criticality: Criticality{Value: CriticalityIgnore},
		Value:       &msg.Cause,
	})
	if len(msg.SNToMNContainer) > 0 {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SNToMNContainer)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       &octetStringIE{Value: msg.SNToMNContainer},
		})
	}
	if msg.SCGUEHistoryInformation != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SCGUEHistoryInformation)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.SCGUEHistoryInformation,
		})
	}
	if msg.PDUSessionsListToBeReleasedUPError != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_PDUSessionsListToBeReleasedUPError)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.PDUSessionsListToBeReleasedUPError,
		})
	}
	return ies
}

func (msg *SNodeReleaseRequired) Encode(w io.Writer) error {
	inner := per.NewEncoder(per.APER)
	if err := encodeBody(msg, inner); err != nil {
		return err
	}

	outer := per.NewEncoder(per.APER)

	if err := choiceXnMsg(outer, common.InitiatingMessage); err != nil {
		return err
	}
	if err := procedureCodeXnMsg(outer, common.ProcedureCode_SNGRANnodeinitiatedSNGRANnodeRelease); err != nil {
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

func (msg *SNodeReleaseRequired) Decode(wire []byte) error {
	d := per.NewDecoder(wire, per.APER)

	seq := d.NewSequenceDecoder(per.SequenceConstraints{Extensible: true})
	if err := seq.DecodeExtensionBit(); err != nil {
		return fmt.Errorf("decode SNodeReleaseRequired extension bit: %w", err)
	}

	seqOf := d.NewSequenceOfDecoder(per.SizeConstraints{
		Extensible: false,
		Min:        common.Ptr(int64(0)),
		Max:        common.Ptr(int64(MaxProtocolIEs)),
	})
	n, err := seqOf.DecodeLength()
	if err != nil {
		return fmt.Errorf("decode SNodeReleaseRequired protocolIEs length: %w", err)
	}

	seen := make(map[int64]bool, n)
	for i := range n {
		id, err := d.DecodeInteger(protocolIEIDConstraints)
		if err != nil {
			return fmt.Errorf("decode SNodeReleaseRequired IE[%d] id: %w", i, err)
		}
		seen[id] = true

		var criticality Criticality
		if err := criticality.Decode(d); err != nil {
			return fmt.Errorf("decode SNodeReleaseRequired IE[%d] criticality: %w", i, err)
		}

		payload, err := d.DecodeOpenType()
		if err != nil {
			return fmt.Errorf("decode SNodeReleaseRequired IE[%d] value: %w", i, err)
		}
		ies := per.NewDecoder(payload, per.APER)

		switch id {
		case common.ProtocolIEID_MNGRANnodeUEXnAPID:
			err = msg.MNGRANnodeUEXnAPID.Decode(ies)
		case common.ProtocolIEID_SNGRANnodeUEXnAPID:
			err = msg.SNGRANnodeUEXnAPID.Decode(ies)
		case common.ProtocolIEID_PDUSessionToBeReleasedListRelRqd:
			v := &PDUSessionToBeReleasedListRelRqd{}
			err = v.Decode(ies)
			msg.PDUSessionToBeReleasedListRelRqd = v
		case common.ProtocolIEID_Cause:
			err = msg.Cause.Decode(ies)
		case common.ProtocolIEID_SNToMNContainer:
			v := octetStringIE{}
			err = v.Decode(ies)
			msg.SNToMNContainer = v.Value
		case common.ProtocolIEID_SCGUEHistoryInformation:
			v := &SCGUEHistoryInformation{}
			err = v.Decode(ies)
			msg.SCGUEHistoryInformation = v
		case common.ProtocolIEID_PDUSessionsListToBeReleasedUPError:
			v := &PDUSessionsListToBeReleasedUPError{}
			err = v.Decode(ies)
			msg.PDUSessionsListToBeReleasedUPError = v
		default:
			if criticality.Value == CriticalityReject {
				return fmt.Errorf("unsupported SNodeReleaseRequired IE id=%d with reject criticality", id)
			}
		}
		if err != nil {
			return fmt.Errorf("decode SNodeReleaseRequired IE id=%d: %w", id, err)
		}
	}

	if _, err := seq.DecodeExtensionAdditions(); err != nil {
		return fmt.Errorf("decode SNodeReleaseRequired extensions: %w", err)
	}

	for _, id := range []int64{
		common.ProtocolIEID_MNGRANnodeUEXnAPID,
		common.ProtocolIEID_SNGRANnodeUEXnAPID,
		common.ProtocolIEID_Cause,
	} {
		if !seen[id] {
			return fmt.Errorf("missing mandatory SNodeReleaseRequired IE id=%d", id)
		}
	}

	return nil
}
