package ies

import (
	"fmt"
	"io"

	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

// NotificationControlIndication ::= SEQUENCE {
//	protocolIEs  ProtocolIE-Container {{NotificationControlIndication-IEs}},
//	...
// }

type NotificationControlIndication struct {
	MNGRANnodeUEXnAPID            NGRANnodeUEXnAPID
	SNGRANnodeUEXnAPID            NGRANnodeUEXnAPID
	PDUSessionResourcesNotifyList *PDUSessionResourcesNotifyList
}

func (msg *NotificationControlIndication) toIEs() []XnAPMessageIE {
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
	if msg.PDUSessionResourcesNotifyList != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_PDUSessionResourcesNotifyList)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.PDUSessionResourcesNotifyList,
		})
	}
	return ies
}

func (msg *NotificationControlIndication) Encode(w io.Writer) error {
	inner := per.NewEncoder(per.APER)
	if err := encodeBody(msg, inner); err != nil {
		return err
	}

	outer := per.NewEncoder(per.APER)

	if err := choiceXnMsg(outer, common.InitiatingMessage); err != nil {
		return err
	}
	if err := procedureCodeXnMsg(outer, common.ProcedureCode_NotificationControl); err != nil {
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

func (msg *NotificationControlIndication) Decode(wire []byte) error {
	d := per.NewDecoder(wire, per.APER)

	seq := d.NewSequenceDecoder(per.SequenceConstraints{Extensible: true})
	if err := seq.DecodeExtensionBit(); err != nil {
		return fmt.Errorf("decode NotificationControlIndication extension bit: %w", err)
	}

	seqOf := d.NewSequenceOfDecoder(per.SizeConstraints{
		Extensible: false,
		Min:        common.Ptr(int64(0)),
		Max:        common.Ptr(int64(MaxProtocolIEs)),
	})
	n, err := seqOf.DecodeLength()
	if err != nil {
		return fmt.Errorf("decode NotificationControlIndication protocolIEs length: %w", err)
	}

	seen := make(map[int64]bool, n)
	for i := range n {
		id, err := d.DecodeInteger(protocolIEIDConstraints)
		if err != nil {
			return fmt.Errorf("decode NotificationControlIndication IE[%d] id: %w", i, err)
		}
		seen[id] = true

		var criticality Criticality
		if err := criticality.Decode(d); err != nil {
			return fmt.Errorf("decode NotificationControlIndication IE[%d] criticality: %w", i, err)
		}

		payload, err := d.DecodeOpenType()
		if err != nil {
			return fmt.Errorf("decode NotificationControlIndication IE[%d] value: %w", i, err)
		}
		ies := per.NewDecoder(payload, per.APER)

		switch id {
		case common.ProtocolIEID_MNGRANnodeUEXnAPID:
			err = msg.MNGRANnodeUEXnAPID.Decode(ies)
		case common.ProtocolIEID_SNGRANnodeUEXnAPID:
			err = msg.SNGRANnodeUEXnAPID.Decode(ies)
		case common.ProtocolIEID_PDUSessionResourcesNotifyList:
			v := &PDUSessionResourcesNotifyList{}
			err = v.Decode(ies)
			msg.PDUSessionResourcesNotifyList = v
		default:
			if criticality.Value == CriticalityReject {
				return fmt.Errorf("unsupported NotificationControlIndication IE id=%d with reject criticality", id)
			}
		}
		if err != nil {
			return fmt.Errorf("decode NotificationControlIndication IE id=%d: %w", id, err)
		}
	}

	if _, err := seq.DecodeExtensionAdditions(); err != nil {
		return fmt.Errorf("decode NotificationControlIndication extensions: %w", err)
	}

	for _, id := range []int64{
		common.ProtocolIEID_MNGRANnodeUEXnAPID,
		common.ProtocolIEID_SNGRANnodeUEXnAPID,
	} {
		if !seen[id] {
			return fmt.Errorf("missing mandatory NotificationControlIndication IE id=%d", id)
		}
	}

	return nil
}
