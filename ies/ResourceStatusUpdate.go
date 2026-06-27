package ies

import (
	"fmt"
	"io"

	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

// ResourceStatusUpdate ::= SEQUENCE {
//	protocolIEs  ProtocolIE-Container {{ResourceStatusUpdate-IEs}},
//	...
// }

type ResourceStatusUpdate struct {
	NGRANNode1MeasurementID MeasurementID
	NGRANNode2MeasurementID MeasurementID
	CellMeasurementResult   CellMeasurementResult
}

func (msg *ResourceStatusUpdate) toIEs() []XnAPMessageIE {
	ies := make([]XnAPMessageIE, 0)
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_NGRANNode1MeasurementID)},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.NGRANNode1MeasurementID,
	})
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_NGRANNode2MeasurementID)},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.NGRANNode2MeasurementID,
	})
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_CellMeasurementResult)},
		Criticality: Criticality{Value: CriticalityIgnore},
		Value:       &msg.CellMeasurementResult,
	})
	return ies
}

func (msg *ResourceStatusUpdate) Encode(w io.Writer) error {
	inner := per.NewEncoder(per.APER)
	if err := encodeBody(msg, inner); err != nil {
		return err
	}

	outer := per.NewEncoder(per.APER)

	if err := choiceXnMsg(outer, common.InitiatingMessage); err != nil {
		return err
	}
	if err := procedureCodeXnMsg(outer, common.ProcedureCode_ResourceStatusReporting); err != nil {
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

func (msg *ResourceStatusUpdate) Decode(wire []byte) error {
	d := per.NewDecoder(wire, per.APER)

	seq := d.NewSequenceDecoder(per.SequenceConstraints{Extensible: true})
	if err := seq.DecodeExtensionBit(); err != nil {
		return fmt.Errorf("decode ResourceStatusUpdate extension bit: %w", err)
	}

	seqOf := d.NewSequenceOfDecoder(per.SizeConstraints{
		Extensible: false,
		Min:        common.Ptr(int64(0)),
		Max:        common.Ptr(int64(MaxProtocolIEs)),
	})
	n, err := seqOf.DecodeLength()
	if err != nil {
		return fmt.Errorf("decode ResourceStatusUpdate protocolIEs length: %w", err)
	}

	seen := make(map[int64]bool, n)
	for i := range n {
		id, err := d.DecodeInteger(protocolIEIDConstraints)
		if err != nil {
			return fmt.Errorf("decode ResourceStatusUpdate IE[%d] id: %w", i, err)
		}
		seen[id] = true

		var criticality Criticality
		if err := criticality.Decode(d); err != nil {
			return fmt.Errorf("decode ResourceStatusUpdate IE[%d] criticality: %w", i, err)
		}

		payload, err := d.DecodeOpenType()
		if err != nil {
			return fmt.Errorf("decode ResourceStatusUpdate IE[%d] value: %w", i, err)
		}
		ies := per.NewDecoder(payload, per.APER)

		switch id {
		case common.ProtocolIEID_NGRANNode1MeasurementID:
			err = msg.NGRANNode1MeasurementID.Decode(ies)
		case common.ProtocolIEID_NGRANNode2MeasurementID:
			err = msg.NGRANNode2MeasurementID.Decode(ies)
		case common.ProtocolIEID_CellMeasurementResult:
			err = msg.CellMeasurementResult.Decode(ies)
		default:
			if criticality.Value == CriticalityReject {
				return fmt.Errorf("unsupported ResourceStatusUpdate IE id=%d with reject criticality", id)
			}
		}
		if err != nil {
			return fmt.Errorf("decode ResourceStatusUpdate IE id=%d: %w", id, err)
		}
	}

	if _, err := seq.DecodeExtensionAdditions(); err != nil {
		return fmt.Errorf("decode ResourceStatusUpdate extensions: %w", err)
	}

	for _, id := range []int64{
		common.ProtocolIEID_NGRANNode1MeasurementID,
		common.ProtocolIEID_NGRANNode2MeasurementID,
		common.ProtocolIEID_CellMeasurementResult,
	} {
		if !seen[id] {
			return fmt.Errorf("missing mandatory ResourceStatusUpdate IE id=%d", id)
		}
	}

	return nil
}
