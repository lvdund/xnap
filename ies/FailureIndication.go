package ies

import (
	"fmt"
	"io"

	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

// FailureIndication ::= SEQUENCE {
//	protocolIEs  ProtocolIE-Container {{FailureIndication-IEs}},
//	...
// }

type FailureIndication struct {
	InitiatingConditionFailureIndication InitiatingConditionFailureIndication
}

func (msg *FailureIndication) toIEs() []XnAPMessageIE {
	ies := make([]XnAPMessageIE, 0)
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_InitiatingConditionFailureIndication)},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.InitiatingConditionFailureIndication,
	})
	return ies
}

func (msg *FailureIndication) Encode(w io.Writer) error {
	inner := per.NewEncoder(per.APER)
	if err := encodeBody(msg, inner); err != nil {
		return err
	}

	outer := per.NewEncoder(per.APER)

	if err := choiceXnMsg(outer, common.InitiatingMessage); err != nil {
		return err
	}
	if err := procedureCodeXnMsg(outer, common.ProcedureCode_FailureIndication); err != nil {
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

func (msg *FailureIndication) Decode(wire []byte) error {
	d := per.NewDecoder(wire, per.APER)

	seq := d.NewSequenceDecoder(per.SequenceConstraints{Extensible: true})
	if err := seq.DecodeExtensionBit(); err != nil {
		return fmt.Errorf("decode FailureIndication extension bit: %w", err)
	}

	seqOf := d.NewSequenceOfDecoder(per.SizeConstraints{
		Extensible: false,
		Min:        common.Ptr(int64(0)),
		Max:        common.Ptr(int64(MaxProtocolIEs)),
	})
	n, err := seqOf.DecodeLength()
	if err != nil {
		return fmt.Errorf("decode FailureIndication protocolIEs length: %w", err)
	}

	seen := make(map[int64]bool, n)
	for i := range n {
		id, err := d.DecodeInteger(protocolIEIDConstraints)
		if err != nil {
			return fmt.Errorf("decode FailureIndication IE[%d] id: %w", i, err)
		}
		seen[id] = true

		var criticality Criticality
		if err := criticality.Decode(d); err != nil {
			return fmt.Errorf("decode FailureIndication IE[%d] criticality: %w", i, err)
		}

		payload, err := d.DecodeOpenType()
		if err != nil {
			return fmt.Errorf("decode FailureIndication IE[%d] value: %w", i, err)
		}
		ies := per.NewDecoder(payload, per.APER)

		switch id {
		case common.ProtocolIEID_InitiatingConditionFailureIndication:
			err = msg.InitiatingConditionFailureIndication.Decode(ies)
		default:
			if criticality.Value == CriticalityReject {
				return fmt.Errorf("unsupported FailureIndication IE id=%d with reject criticality", id)
			}
		}
		if err != nil {
			return fmt.Errorf("decode FailureIndication IE id=%d: %w", id, err)
		}
	}

	if _, err := seq.DecodeExtensionAdditions(); err != nil {
		return fmt.Errorf("decode FailureIndication extensions: %w", err)
	}

	for _, id := range []int64{
		common.ProtocolIEID_InitiatingConditionFailureIndication,
	} {
		if !seen[id] {
			return fmt.Errorf("missing mandatory FailureIndication IE id=%d", id)
		}
	}

	return nil
}
