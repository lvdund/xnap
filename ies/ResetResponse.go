package ies

import (
	"fmt"
	"io"

	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

// ResetResponse ::= SEQUENCE {
//	protocolIEs  ProtocolIE-Container {{ResetResponse-IEs}},
//	...
// }

type ResetResponse struct {
	ResetResponseTypeInfo       ResetResponseTypeInfo
	CriticalityDiagnostics      *CriticalityDiagnostics
	InterfaceInstanceIndication *InterfaceInstanceIndication
}

func (msg *ResetResponse) toIEs() []XnAPMessageIE {
	ies := make([]XnAPMessageIE, 0)
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_ResetResponseTypeInfo)},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.ResetResponseTypeInfo,
	})
	if msg.CriticalityDiagnostics != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_CriticalityDiagnostics)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.CriticalityDiagnostics,
		})
	}
	if msg.InterfaceInstanceIndication != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_InterfaceInstanceIndication)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.InterfaceInstanceIndication,
		})
	}
	return ies
}

func (msg *ResetResponse) Encode(w io.Writer) error {
	inner := per.NewEncoder(per.APER)
	if err := encodeBody(msg, inner); err != nil {
		return err
	}

	outer := per.NewEncoder(per.APER)

	if err := choiceXnMsg(outer, common.SuccessfulOutcome); err != nil {
		return err
	}
	if err := procedureCodeXnMsg(outer, common.ProcedureCode_Reset); err != nil {
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

func (msg *ResetResponse) Decode(wire []byte) error {
	d := per.NewDecoder(wire, per.APER)

	seq := d.NewSequenceDecoder(per.SequenceConstraints{Extensible: true})
	if err := seq.DecodeExtensionBit(); err != nil {
		return fmt.Errorf("decode ResetResponse extension bit: %w", err)
	}

	seqOf := d.NewSequenceOfDecoder(per.SizeConstraints{
		Extensible: false,
		Min:        common.Ptr(int64(0)),
		Max:        common.Ptr(int64(MaxProtocolIEs)),
	})
	n, err := seqOf.DecodeLength()
	if err != nil {
		return fmt.Errorf("decode ResetResponse protocolIEs length: %w", err)
	}

	seen := make(map[int64]bool, n)
	for i := range n {
		id, err := d.DecodeInteger(protocolIEIDConstraints)
		if err != nil {
			return fmt.Errorf("decode ResetResponse IE[%d] id: %w", i, err)
		}
		seen[id] = true

		var criticality Criticality
		if err := criticality.Decode(d); err != nil {
			return fmt.Errorf("decode ResetResponse IE[%d] criticality: %w", i, err)
		}

		payload, err := d.DecodeOpenType()
		if err != nil {
			return fmt.Errorf("decode ResetResponse IE[%d] value: %w", i, err)
		}
		ies := per.NewDecoder(payload, per.APER)

		switch id {
		case common.ProtocolIEID_ResetResponseTypeInfo:
			err = msg.ResetResponseTypeInfo.Decode(ies)
		case common.ProtocolIEID_CriticalityDiagnostics:
			v := &CriticalityDiagnostics{}
			err = v.Decode(ies)
			msg.CriticalityDiagnostics = v
		case common.ProtocolIEID_InterfaceInstanceIndication:
			v := &InterfaceInstanceIndication{}
			err = v.Decode(ies)
			msg.InterfaceInstanceIndication = v
		default:
			if criticality.Value == CriticalityReject {
				return fmt.Errorf("unsupported ResetResponse IE id=%d with reject criticality", id)
			}
		}
		if err != nil {
			return fmt.Errorf("decode ResetResponse IE id=%d: %w", id, err)
		}
	}

	if _, err := seq.DecodeExtensionAdditions(); err != nil {
		return fmt.Errorf("decode ResetResponse extensions: %w", err)
	}

	for _, id := range []int64{
		common.ProtocolIEID_ResetResponseTypeInfo,
	} {
		if !seen[id] {
			return fmt.Errorf("missing mandatory ResetResponse IE id=%d", id)
		}
	}

	return nil
}
