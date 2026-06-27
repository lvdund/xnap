package ies

import (
	"fmt"
	"io"

	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

// ErrorIndication ::= SEQUENCE {
//	protocolIEs  ProtocolIE-Container {{ErrorIndication-IEs}},
//	...
// }

type ErrorIndication struct {
	OldNGRANnodeUEXnAPID        *NGRANnodeUEXnAPID
	NewNGRANnodeUEXnAPID        *NGRANnodeUEXnAPID
	Cause                       *Cause
	CriticalityDiagnostics      *CriticalityDiagnostics
	InterfaceInstanceIndication *InterfaceInstanceIndication
}

func (msg *ErrorIndication) toIEs() []XnAPMessageIE {
	ies := make([]XnAPMessageIE, 0)
	if msg.OldNGRANnodeUEXnAPID != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_OldNGRANnodeUEXnAPID)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.OldNGRANnodeUEXnAPID,
		})
	}
	if msg.NewNGRANnodeUEXnAPID != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_NewNGRANnodeUEXnAPID)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.NewNGRANnodeUEXnAPID,
		})
	}
	if msg.Cause != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_Cause)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.Cause,
		})
	}
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

func (msg *ErrorIndication) Encode(w io.Writer) error {
	inner := per.NewEncoder(per.APER)
	if err := encodeBody(msg, inner); err != nil {
		return err
	}

	outer := per.NewEncoder(per.APER)

	if err := choiceXnMsg(outer, common.InitiatingMessage); err != nil {
		return err
	}
	if err := procedureCodeXnMsg(outer, common.ProcedureCode_ErrorIndication); err != nil {
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

func (msg *ErrorIndication) Decode(wire []byte) error {
	d := per.NewDecoder(wire, per.APER)

	seq := d.NewSequenceDecoder(per.SequenceConstraints{Extensible: true})
	if err := seq.DecodeExtensionBit(); err != nil {
		return fmt.Errorf("decode ErrorIndication extension bit: %w", err)
	}

	seqOf := d.NewSequenceOfDecoder(per.SizeConstraints{
		Extensible: false,
		Min:        common.Ptr(int64(0)),
		Max:        common.Ptr(int64(MaxProtocolIEs)),
	})
	n, err := seqOf.DecodeLength()
	if err != nil {
		return fmt.Errorf("decode ErrorIndication protocolIEs length: %w", err)
	}

	seen := make(map[int64]bool, n)
	for i := range n {
		id, err := d.DecodeInteger(protocolIEIDConstraints)
		if err != nil {
			return fmt.Errorf("decode ErrorIndication IE[%d] id: %w", i, err)
		}
		seen[id] = true

		var criticality Criticality
		if err := criticality.Decode(d); err != nil {
			return fmt.Errorf("decode ErrorIndication IE[%d] criticality: %w", i, err)
		}

		payload, err := d.DecodeOpenType()
		if err != nil {
			return fmt.Errorf("decode ErrorIndication IE[%d] value: %w", i, err)
		}
		ies := per.NewDecoder(payload, per.APER)

		switch id {
		case common.ProtocolIEID_OldNGRANnodeUEXnAPID:
			v := &NGRANnodeUEXnAPID{}
			err = v.Decode(ies)
			msg.OldNGRANnodeUEXnAPID = v
		case common.ProtocolIEID_NewNGRANnodeUEXnAPID:
			v := &NGRANnodeUEXnAPID{}
			err = v.Decode(ies)
			msg.NewNGRANnodeUEXnAPID = v
		case common.ProtocolIEID_Cause:
			v := &Cause{}
			err = v.Decode(ies)
			msg.Cause = v
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
				return fmt.Errorf("unsupported ErrorIndication IE id=%d with reject criticality", id)
			}
		}
		if err != nil {
			return fmt.Errorf("decode ErrorIndication IE id=%d: %w", id, err)
		}
	}

	if _, err := seq.DecodeExtensionAdditions(); err != nil {
		return fmt.Errorf("decode ErrorIndication extensions: %w", err)
	}

	return nil
}
