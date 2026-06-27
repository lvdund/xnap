package ies

import (
	"fmt"
	"io"

	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

// ActivityNotification ::= SEQUENCE {
//	protocolIEs  ProtocolIE-Container {{ActivityNotification-IEs}},
//	...
// }

type ActivityNotification struct {
	MNGRANnodeUEXnAPID                    NGRANnodeUEXnAPID
	SNGRANnodeUEXnAPID                    NGRANnodeUEXnAPID
	UserPlaneTrafficActivityReport        *UserPlaneTrafficActivityReport
	PDUSessionResourcesActivityNotifyList *PDUSessionResourcesActivityNotifyList
	RANPagingFailure                      *RANPagingFailure
}

func (msg *ActivityNotification) toIEs() []XnAPMessageIE {
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
	if msg.UserPlaneTrafficActivityReport != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_UserPlaneTrafficActivityReport)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.UserPlaneTrafficActivityReport,
		})
	}
	if msg.PDUSessionResourcesActivityNotifyList != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_PDUSessionResourcesActivityNotifyList)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.PDUSessionResourcesActivityNotifyList,
		})
	}
	if msg.RANPagingFailure != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_RANPagingFailure)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.RANPagingFailure,
		})
	}
	return ies
}

func (msg *ActivityNotification) Encode(w io.Writer) error {
	inner := per.NewEncoder(per.APER)
	if err := encodeBody(msg, inner); err != nil {
		return err
	}

	outer := per.NewEncoder(per.APER)

	if err := choiceXnMsg(outer, common.InitiatingMessage); err != nil {
		return err
	}
	if err := procedureCodeXnMsg(outer, common.ProcedureCode_ActivityNotification); err != nil {
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

func (msg *ActivityNotification) Decode(wire []byte) error {
	d := per.NewDecoder(wire, per.APER)

	seq := d.NewSequenceDecoder(per.SequenceConstraints{Extensible: true})
	if err := seq.DecodeExtensionBit(); err != nil {
		return fmt.Errorf("decode ActivityNotification extension bit: %w", err)
	}

	seqOf := d.NewSequenceOfDecoder(per.SizeConstraints{
		Extensible: false,
		Min:        common.Ptr(int64(0)),
		Max:        common.Ptr(int64(MaxProtocolIEs)),
	})
	n, err := seqOf.DecodeLength()
	if err != nil {
		return fmt.Errorf("decode ActivityNotification protocolIEs length: %w", err)
	}

	seen := make(map[int64]bool, n)
	for i := range n {
		id, err := d.DecodeInteger(protocolIEIDConstraints)
		if err != nil {
			return fmt.Errorf("decode ActivityNotification IE[%d] id: %w", i, err)
		}
		seen[id] = true

		var criticality Criticality
		if err := criticality.Decode(d); err != nil {
			return fmt.Errorf("decode ActivityNotification IE[%d] criticality: %w", i, err)
		}

		payload, err := d.DecodeOpenType()
		if err != nil {
			return fmt.Errorf("decode ActivityNotification IE[%d] value: %w", i, err)
		}
		ies := per.NewDecoder(payload, per.APER)

		switch id {
		case common.ProtocolIEID_MNGRANnodeUEXnAPID:
			err = msg.MNGRANnodeUEXnAPID.Decode(ies)
		case common.ProtocolIEID_SNGRANnodeUEXnAPID:
			err = msg.SNGRANnodeUEXnAPID.Decode(ies)
		case common.ProtocolIEID_UserPlaneTrafficActivityReport:
			v := &UserPlaneTrafficActivityReport{}
			err = v.Decode(ies)
			msg.UserPlaneTrafficActivityReport = v
		case common.ProtocolIEID_PDUSessionResourcesActivityNotifyList:
			v := &PDUSessionResourcesActivityNotifyList{}
			err = v.Decode(ies)
			msg.PDUSessionResourcesActivityNotifyList = v
		case common.ProtocolIEID_RANPagingFailure:
			v := &RANPagingFailure{}
			err = v.Decode(ies)
			msg.RANPagingFailure = v
		default:
			if criticality.Value == CriticalityReject {
				return fmt.Errorf("unsupported ActivityNotification IE id=%d with reject criticality", id)
			}
		}
		if err != nil {
			return fmt.Errorf("decode ActivityNotification IE id=%d: %w", id, err)
		}
	}

	if _, err := seq.DecodeExtensionAdditions(); err != nil {
		return fmt.Errorf("decode ActivityNotification extensions: %w", err)
	}

	for _, id := range []int64{
		common.ProtocolIEID_MNGRANnodeUEXnAPID,
		common.ProtocolIEID_SNGRANnodeUEXnAPID,
	} {
		if !seen[id] {
			return fmt.Errorf("missing mandatory ActivityNotification IE id=%d", id)
		}
	}

	return nil
}
