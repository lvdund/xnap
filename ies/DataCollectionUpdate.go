package ies

import (
	"fmt"
	"io"

	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

// DataCollectionUpdate ::= SEQUENCE {
//	protocolIEs  ProtocolIE-Container {{DataCollectionUpdate-IEs}},
//	...
// }

type DataCollectionUpdate struct {
	NGRANNode1MeasurementID                    MeasurementID
	NGRANNode2MeasurementID                    MeasurementID
	CellMeasurementResultForDataCollectionList *CellMeasurementResultForDataCollectionList
	UEAssociatedInfoResultList                 *UEAssociatedInfoResultList
	NodeAssociatedInfoResult                   *NodeAssociatedInfoResult
}

func (msg *DataCollectionUpdate) toIEs() []XnAPMessageIE {
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
	if msg.CellMeasurementResultForDataCollectionList != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_CellMeasurementResultForDataCollectionList)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.CellMeasurementResultForDataCollectionList,
		})
	}
	if msg.UEAssociatedInfoResultList != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_UEAssociatedInfoResultList)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.UEAssociatedInfoResultList,
		})
	}
	if msg.NodeAssociatedInfoResult != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_NodeAssociatedInfoResult)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.NodeAssociatedInfoResult,
		})
	}
	return ies
}

func (msg *DataCollectionUpdate) Encode(w io.Writer) error {
	inner := per.NewEncoder(per.APER)
	if err := encodeBody(msg, inner); err != nil {
		return err
	}

	outer := per.NewEncoder(per.APER)

	if err := choiceXnMsg(outer, common.InitiatingMessage); err != nil {
		return err
	}
	if err := procedureCodeXnMsg(outer, common.ProcedureCode_DataCollectionReporting); err != nil {
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

func (msg *DataCollectionUpdate) Decode(wire []byte) error {
	d := per.NewDecoder(wire, per.APER)

	seq := d.NewSequenceDecoder(per.SequenceConstraints{Extensible: true})
	if err := seq.DecodeExtensionBit(); err != nil {
		return fmt.Errorf("decode DataCollectionUpdate extension bit: %w", err)
	}

	seqOf := d.NewSequenceOfDecoder(per.SizeConstraints{
		Extensible: false,
		Min:        common.Ptr(int64(0)),
		Max:        common.Ptr(int64(MaxProtocolIEs)),
	})
	n, err := seqOf.DecodeLength()
	if err != nil {
		return fmt.Errorf("decode DataCollectionUpdate protocolIEs length: %w", err)
	}

	seen := make(map[int64]bool, n)
	for i := range n {
		id, err := d.DecodeInteger(protocolIEIDConstraints)
		if err != nil {
			return fmt.Errorf("decode DataCollectionUpdate IE[%d] id: %w", i, err)
		}
		seen[id] = true

		var criticality Criticality
		if err := criticality.Decode(d); err != nil {
			return fmt.Errorf("decode DataCollectionUpdate IE[%d] criticality: %w", i, err)
		}

		payload, err := d.DecodeOpenType()
		if err != nil {
			return fmt.Errorf("decode DataCollectionUpdate IE[%d] value: %w", i, err)
		}
		ies := per.NewDecoder(payload, per.APER)

		switch id {
		case common.ProtocolIEID_NGRANNode1MeasurementID:
			err = msg.NGRANNode1MeasurementID.Decode(ies)
		case common.ProtocolIEID_NGRANNode2MeasurementID:
			err = msg.NGRANNode2MeasurementID.Decode(ies)
		case common.ProtocolIEID_CellMeasurementResultForDataCollectionList:
			v := &CellMeasurementResultForDataCollectionList{}
			err = v.Decode(ies)
			msg.CellMeasurementResultForDataCollectionList = v
		case common.ProtocolIEID_UEAssociatedInfoResultList:
			v := &UEAssociatedInfoResultList{}
			err = v.Decode(ies)
			msg.UEAssociatedInfoResultList = v
		case common.ProtocolIEID_NodeAssociatedInfoResult:
			v := &NodeAssociatedInfoResult{}
			err = v.Decode(ies)
			msg.NodeAssociatedInfoResult = v
		default:
			if criticality.Value == CriticalityReject {
				return fmt.Errorf("unsupported DataCollectionUpdate IE id=%d with reject criticality", id)
			}
		}
		if err != nil {
			return fmt.Errorf("decode DataCollectionUpdate IE id=%d: %w", id, err)
		}
	}

	if _, err := seq.DecodeExtensionAdditions(); err != nil {
		return fmt.Errorf("decode DataCollectionUpdate extensions: %w", err)
	}

	for _, id := range []int64{
		common.ProtocolIEID_NGRANNode1MeasurementID,
		common.ProtocolIEID_NGRANNode2MeasurementID,
	} {
		if !seen[id] {
			return fmt.Errorf("missing mandatory DataCollectionUpdate IE id=%d", id)
		}
	}

	return nil
}
