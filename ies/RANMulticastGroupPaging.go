package ies

import (
	"fmt"
	"io"

	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

// RANMulticastGroupPaging ::= SEQUENCE {
//	protocolIEs  ProtocolIE-Container {{RANMulticastGroupPaging-IEs}},
//	...
// }

type RANMulticastGroupPaging struct {
	MBSSessionID                      MBSSessionID
	UEIdentityIndexListMBSGroupPaging UEIdentityIndexListMBSGroupPaging
	MulticastRANPagingArea            RANPagingArea
}

func (msg *RANMulticastGroupPaging) toIEs() []XnAPMessageIE {
	ies := make([]XnAPMessageIE, 0)
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_MBSSessionID)},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.MBSSessionID,
	})
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_UEIdentityIndexListMBSGroupPaging)},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.UEIdentityIndexListMBSGroupPaging,
	})
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_MulticastRANPagingArea)},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.MulticastRANPagingArea,
	})
	return ies
}

func (msg *RANMulticastGroupPaging) Encode(w io.Writer) error {
	inner := per.NewEncoder(per.APER)
	if err := encodeBody(msg, inner); err != nil {
		return err
	}

	outer := per.NewEncoder(per.APER)

	if err := choiceXnMsg(outer, common.InitiatingMessage); err != nil {
		return err
	}
	if err := procedureCodeXnMsg(outer, common.ProcedureCode_RANMulticastGroupPaging); err != nil {
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

func (msg *RANMulticastGroupPaging) Decode(wire []byte) error {
	d := per.NewDecoder(wire, per.APER)

	seq := d.NewSequenceDecoder(per.SequenceConstraints{Extensible: true})
	if err := seq.DecodeExtensionBit(); err != nil {
		return fmt.Errorf("decode RANMulticastGroupPaging extension bit: %w", err)
	}

	seqOf := d.NewSequenceOfDecoder(per.SizeConstraints{
		Extensible: false,
		Min:        common.Ptr(int64(0)),
		Max:        common.Ptr(int64(MaxProtocolIEs)),
	})
	n, err := seqOf.DecodeLength()
	if err != nil {
		return fmt.Errorf("decode RANMulticastGroupPaging protocolIEs length: %w", err)
	}

	seen := make(map[int64]bool, n)
	for i := range n {
		id, err := d.DecodeInteger(protocolIEIDConstraints)
		if err != nil {
			return fmt.Errorf("decode RANMulticastGroupPaging IE[%d] id: %w", i, err)
		}
		seen[id] = true

		var criticality Criticality
		if err := criticality.Decode(d); err != nil {
			return fmt.Errorf("decode RANMulticastGroupPaging IE[%d] criticality: %w", i, err)
		}

		payload, err := d.DecodeOpenType()
		if err != nil {
			return fmt.Errorf("decode RANMulticastGroupPaging IE[%d] value: %w", i, err)
		}
		ies := per.NewDecoder(payload, per.APER)

		switch id {
		case common.ProtocolIEID_MBSSessionID:
			err = msg.MBSSessionID.Decode(ies)
		case common.ProtocolIEID_UEIdentityIndexListMBSGroupPaging:
			err = msg.UEIdentityIndexListMBSGroupPaging.Decode(ies)
		case common.ProtocolIEID_MulticastRANPagingArea:
			err = msg.MulticastRANPagingArea.Decode(ies)
		default:
			if criticality.Value == CriticalityReject {
				return fmt.Errorf("unsupported RANMulticastGroupPaging IE id=%d with reject criticality", id)
			}
		}
		if err != nil {
			return fmt.Errorf("decode RANMulticastGroupPaging IE id=%d: %w", id, err)
		}
	}

	if _, err := seq.DecodeExtensionAdditions(); err != nil {
		return fmt.Errorf("decode RANMulticastGroupPaging extensions: %w", err)
	}

	for _, id := range []int64{
		common.ProtocolIEID_MBSSessionID,
		common.ProtocolIEID_UEIdentityIndexListMBSGroupPaging,
		common.ProtocolIEID_MulticastRANPagingArea,
	} {
		if !seen[id] {
			return fmt.Errorf("missing mandatory RANMulticastGroupPaging IE id=%d", id)
		}
	}

	return nil
}
