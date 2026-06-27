package ies

import (
	"fmt"
	"io"

	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

// DataCollectionResponse ::= SEQUENCE {
//	protocolIEs  ProtocolIE-Container {{DataCollectionResponse-IEs}},
//	...
// }

type DataCollectionResponse struct {
	NGRANNode1MeasurementID             MeasurementID
	NGRANNode2MeasurementID             MeasurementID
	NodeMeasurementInitiationResultList *NodeMeasurementInitiationResultList
	CellMeasurementInitiationResultList *CellMeasurementInitiationResultList
	CriticalityDiagnostics              *CriticalityDiagnostics
}

func (msg *DataCollectionResponse) toIEs() []XnAPMessageIE {
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
	if msg.NodeMeasurementInitiationResultList != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_NodeMeasurementInitiationResultList)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.NodeMeasurementInitiationResultList,
		})
	}
	if msg.CellMeasurementInitiationResultList != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_CellMeasurementInitiationResultList)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.CellMeasurementInitiationResultList,
		})
	}
	if msg.CriticalityDiagnostics != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_CriticalityDiagnostics)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.CriticalityDiagnostics,
		})
	}
	return ies
}

func (msg *DataCollectionResponse) Encode(w io.Writer) error {
	inner := per.NewEncoder(per.APER)
	if err := encodeBody(msg, inner); err != nil {
		return err
	}

	outer := per.NewEncoder(per.APER)

	if err := choiceXnMsg(outer, common.SuccessfulOutcome); err != nil {
		return err
	}
	if err := procedureCodeXnMsg(outer, common.ProcedureCode_DataCollectionReportingInitiation); err != nil {
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

func (msg *DataCollectionResponse) Decode(wire []byte) error {
	d := per.NewDecoder(wire, per.APER)

	seq := d.NewSequenceDecoder(per.SequenceConstraints{Extensible: true})
	if err := seq.DecodeExtensionBit(); err != nil {
		return fmt.Errorf("decode DataCollectionResponse extension bit: %w", err)
	}

	seqOf := d.NewSequenceOfDecoder(per.SizeConstraints{
		Extensible: false,
		Min:        common.Ptr(int64(0)),
		Max:        common.Ptr(int64(MaxProtocolIEs)),
	})
	n, err := seqOf.DecodeLength()
	if err != nil {
		return fmt.Errorf("decode DataCollectionResponse protocolIEs length: %w", err)
	}

	seen := make(map[int64]bool, n)
	for i := range n {
		id, err := d.DecodeInteger(protocolIEIDConstraints)
		if err != nil {
			return fmt.Errorf("decode DataCollectionResponse IE[%d] id: %w", i, err)
		}
		seen[id] = true

		var criticality Criticality
		if err := criticality.Decode(d); err != nil {
			return fmt.Errorf("decode DataCollectionResponse IE[%d] criticality: %w", i, err)
		}

		payload, err := d.DecodeOpenType()
		if err != nil {
			return fmt.Errorf("decode DataCollectionResponse IE[%d] value: %w", i, err)
		}
		ies := per.NewDecoder(payload, per.APER)

		switch id {
		case common.ProtocolIEID_NGRANNode1MeasurementID:
			err = msg.NGRANNode1MeasurementID.Decode(ies)
		case common.ProtocolIEID_NGRANNode2MeasurementID:
			err = msg.NGRANNode2MeasurementID.Decode(ies)
		case common.ProtocolIEID_NodeMeasurementInitiationResultList:
			v := &NodeMeasurementInitiationResultList{}
			err = v.Decode(ies)
			msg.NodeMeasurementInitiationResultList = v
		case common.ProtocolIEID_CellMeasurementInitiationResultList:
			v := &CellMeasurementInitiationResultList{}
			err = v.Decode(ies)
			msg.CellMeasurementInitiationResultList = v
		case common.ProtocolIEID_CriticalityDiagnostics:
			v := &CriticalityDiagnostics{}
			err = v.Decode(ies)
			msg.CriticalityDiagnostics = v
		default:
			if criticality.Value == CriticalityReject {
				return fmt.Errorf("unsupported DataCollectionResponse IE id=%d with reject criticality", id)
			}
		}
		if err != nil {
			return fmt.Errorf("decode DataCollectionResponse IE id=%d: %w", id, err)
		}
	}

	if _, err := seq.DecodeExtensionAdditions(); err != nil {
		return fmt.Errorf("decode DataCollectionResponse extensions: %w", err)
	}

	for _, id := range []int64{
		common.ProtocolIEID_NGRANNode1MeasurementID,
		common.ProtocolIEID_NGRANNode2MeasurementID,
	} {
		if !seen[id] {
			return fmt.Errorf("missing mandatory DataCollectionResponse IE id=%d", id)
		}
	}

	return nil
}
