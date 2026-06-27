package ies

import (
	"fmt"
	"io"

	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

// MobilityChangeFailure ::= SEQUENCE {
//	protocolIEs  ProtocolIE-Container {{MobilityChangeFailure-IEs}},
//	...
// }

type MobilityChangeFailure struct {
	NGRANnode1CellID                      GlobalNGRANCellID
	NGRANnode2CellID                      GlobalNGRANCellID
	Cause                                 Cause
	MobilityParametersModificationRange   *MobilityParametersModificationRange
	CriticalityDiagnostics                *CriticalityDiagnostics
	NGRANnode2SSBOffsetsModificationRange *NGRANnode2SSBOffsetsModificationRange
}

func (msg *MobilityChangeFailure) toIEs() []XnAPMessageIE {
	ies := make([]XnAPMessageIE, 0)
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_NGRANnode1CellID)},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.NGRANnode1CellID,
	})
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_NGRANnode2CellID)},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.NGRANnode2CellID,
	})
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_Cause)},
		Criticality: Criticality{Value: CriticalityIgnore},
		Value:       &msg.Cause,
	})
	if msg.MobilityParametersModificationRange != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_MobilityParametersModificationRange)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.MobilityParametersModificationRange,
		})
	}
	if msg.CriticalityDiagnostics != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_CriticalityDiagnostics)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.CriticalityDiagnostics,
		})
	}
	if msg.NGRANnode2SSBOffsetsModificationRange != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_NGRANnode2SSBOffsetsModificationRange)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.NGRANnode2SSBOffsetsModificationRange,
		})
	}
	return ies
}

func (msg *MobilityChangeFailure) Encode(w io.Writer) error {
	inner := per.NewEncoder(per.APER)
	if err := encodeBody(msg, inner); err != nil {
		return err
	}

	outer := per.NewEncoder(per.APER)

	if err := choiceXnMsg(outer, common.UnsuccessfulOutcome); err != nil {
		return err
	}
	if err := procedureCodeXnMsg(outer, common.ProcedureCode_MobilitySettingsChange); err != nil {
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

func (msg *MobilityChangeFailure) Decode(wire []byte) error {
	d := per.NewDecoder(wire, per.APER)

	seq := d.NewSequenceDecoder(per.SequenceConstraints{Extensible: true})
	if err := seq.DecodeExtensionBit(); err != nil {
		return fmt.Errorf("decode MobilityChangeFailure extension bit: %w", err)
	}

	seqOf := d.NewSequenceOfDecoder(per.SizeConstraints{
		Extensible: false,
		Min:        common.Ptr(int64(0)),
		Max:        common.Ptr(int64(MaxProtocolIEs)),
	})
	n, err := seqOf.DecodeLength()
	if err != nil {
		return fmt.Errorf("decode MobilityChangeFailure protocolIEs length: %w", err)
	}

	seen := make(map[int64]bool, n)
	for i := range n {
		id, err := d.DecodeInteger(protocolIEIDConstraints)
		if err != nil {
			return fmt.Errorf("decode MobilityChangeFailure IE[%d] id: %w", i, err)
		}
		seen[id] = true

		var criticality Criticality
		if err := criticality.Decode(d); err != nil {
			return fmt.Errorf("decode MobilityChangeFailure IE[%d] criticality: %w", i, err)
		}

		payload, err := d.DecodeOpenType()
		if err != nil {
			return fmt.Errorf("decode MobilityChangeFailure IE[%d] value: %w", i, err)
		}
		ies := per.NewDecoder(payload, per.APER)

		switch id {
		case common.ProtocolIEID_NGRANnode1CellID:
			err = msg.NGRANnode1CellID.Decode(ies)
		case common.ProtocolIEID_NGRANnode2CellID:
			err = msg.NGRANnode2CellID.Decode(ies)
		case common.ProtocolIEID_Cause:
			err = msg.Cause.Decode(ies)
		case common.ProtocolIEID_MobilityParametersModificationRange:
			v := &MobilityParametersModificationRange{}
			err = v.Decode(ies)
			msg.MobilityParametersModificationRange = v
		case common.ProtocolIEID_CriticalityDiagnostics:
			v := &CriticalityDiagnostics{}
			err = v.Decode(ies)
			msg.CriticalityDiagnostics = v
		case common.ProtocolIEID_NGRANnode2SSBOffsetsModificationRange:
			v := &NGRANnode2SSBOffsetsModificationRange{}
			err = v.Decode(ies)
			msg.NGRANnode2SSBOffsetsModificationRange = v
		default:
			if criticality.Value == CriticalityReject {
				return fmt.Errorf("unsupported MobilityChangeFailure IE id=%d with reject criticality", id)
			}
		}
		if err != nil {
			return fmt.Errorf("decode MobilityChangeFailure IE id=%d: %w", id, err)
		}
	}

	if _, err := seq.DecodeExtensionAdditions(); err != nil {
		return fmt.Errorf("decode MobilityChangeFailure extensions: %w", err)
	}

	for _, id := range []int64{
		common.ProtocolIEID_NGRANnode1CellID,
		common.ProtocolIEID_NGRANnode2CellID,
		common.ProtocolIEID_Cause,
	} {
		if !seen[id] {
			return fmt.Errorf("missing mandatory MobilityChangeFailure IE id=%d", id)
		}
	}

	return nil
}
