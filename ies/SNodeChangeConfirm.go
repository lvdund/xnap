package ies

import (
	"fmt"
	"io"

	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

// SNodeChangeConfirm ::= SEQUENCE {
//	protocolIEs  ProtocolIE-Container {{SNodeChangeConfirm-IEs}},
//	...
// }

type SNodeChangeConfirm struct {
	MNGRANnodeUEXnAPID            NGRANnodeUEXnAPID
	SNGRANnodeUEXnAPID            NGRANnodeUEXnAPID
	PDUSessionSNChangeConfirmList *PDUSessionSNChangeConfirmList
	CriticalityDiagnostics        *CriticalityDiagnostics
	CPCInformationConfirm         *CPCInformationConfirm
	MNToSNContainer               []byte
}

func (msg *SNodeChangeConfirm) toIEs() []XnAPMessageIE {
	ies := make([]XnAPMessageIE, 0)
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_MNGRANnodeUEXnAPID)},
		Criticality: Criticality{Value: CriticalityIgnore},
		Value:       &msg.MNGRANnodeUEXnAPID,
	})
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SNGRANnodeUEXnAPID)},
		Criticality: Criticality{Value: CriticalityIgnore},
		Value:       &msg.SNGRANnodeUEXnAPID,
	})
	if msg.PDUSessionSNChangeConfirmList != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_PDUSessionSNChangeConfirmList)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.PDUSessionSNChangeConfirmList,
		})
	}
	if msg.CriticalityDiagnostics != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_CriticalityDiagnostics)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.CriticalityDiagnostics,
		})
	}
	if msg.CPCInformationConfirm != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_CPCInformationConfirm)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.CPCInformationConfirm,
		})
	}
	if len(msg.MNToSNContainer) > 0 {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_MNToSNContainer)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       &octetStringIE{Value: msg.MNToSNContainer},
		})
	}
	return ies
}

func (msg *SNodeChangeConfirm) Encode(w io.Writer) error {
	inner := per.NewEncoder(per.APER)
	if err := encodeBody(msg, inner); err != nil {
		return err
	}

	outer := per.NewEncoder(per.APER)

	if err := choiceXnMsg(outer, common.SuccessfulOutcome); err != nil {
		return err
	}
	if err := procedureCodeXnMsg(outer, common.ProcedureCode_SNGRANnodeChange); err != nil {
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

func (msg *SNodeChangeConfirm) Decode(wire []byte) error {
	d := per.NewDecoder(wire, per.APER)

	seq := d.NewSequenceDecoder(per.SequenceConstraints{Extensible: true})
	if err := seq.DecodeExtensionBit(); err != nil {
		return fmt.Errorf("decode SNodeChangeConfirm extension bit: %w", err)
	}

	seqOf := d.NewSequenceOfDecoder(per.SizeConstraints{
		Extensible: false,
		Min:        common.Ptr(int64(0)),
		Max:        common.Ptr(int64(MaxProtocolIEs)),
	})
	n, err := seqOf.DecodeLength()
	if err != nil {
		return fmt.Errorf("decode SNodeChangeConfirm protocolIEs length: %w", err)
	}

	seen := make(map[int64]bool, n)
	for i := range n {
		id, err := d.DecodeInteger(protocolIEIDConstraints)
		if err != nil {
			return fmt.Errorf("decode SNodeChangeConfirm IE[%d] id: %w", i, err)
		}
		seen[id] = true

		var criticality Criticality
		if err := criticality.Decode(d); err != nil {
			return fmt.Errorf("decode SNodeChangeConfirm IE[%d] criticality: %w", i, err)
		}

		payload, err := d.DecodeOpenType()
		if err != nil {
			return fmt.Errorf("decode SNodeChangeConfirm IE[%d] value: %w", i, err)
		}
		ies := per.NewDecoder(payload, per.APER)

		switch id {
		case common.ProtocolIEID_MNGRANnodeUEXnAPID:
			err = msg.MNGRANnodeUEXnAPID.Decode(ies)
		case common.ProtocolIEID_SNGRANnodeUEXnAPID:
			err = msg.SNGRANnodeUEXnAPID.Decode(ies)
		case common.ProtocolIEID_PDUSessionSNChangeConfirmList:
			v := &PDUSessionSNChangeConfirmList{}
			err = v.Decode(ies)
			msg.PDUSessionSNChangeConfirmList = v
		case common.ProtocolIEID_CriticalityDiagnostics:
			v := &CriticalityDiagnostics{}
			err = v.Decode(ies)
			msg.CriticalityDiagnostics = v
		case common.ProtocolIEID_CPCInformationConfirm:
			v := &CPCInformationConfirm{}
			err = v.Decode(ies)
			msg.CPCInformationConfirm = v
		case common.ProtocolIEID_MNToSNContainer:
			v := octetStringIE{}
			err = v.Decode(ies)
			msg.MNToSNContainer = v.Value
		default:
			if criticality.Value == CriticalityReject {
				return fmt.Errorf("unsupported SNodeChangeConfirm IE id=%d with reject criticality", id)
			}
		}
		if err != nil {
			return fmt.Errorf("decode SNodeChangeConfirm IE id=%d: %w", id, err)
		}
	}

	if _, err := seq.DecodeExtensionAdditions(); err != nil {
		return fmt.Errorf("decode SNodeChangeConfirm extensions: %w", err)
	}

	for _, id := range []int64{
		common.ProtocolIEID_MNGRANnodeUEXnAPID,
		common.ProtocolIEID_SNGRANnodeUEXnAPID,
	} {
		if !seen[id] {
			return fmt.Errorf("missing mandatory SNodeChangeConfirm IE id=%d", id)
		}
	}

	return nil
}
