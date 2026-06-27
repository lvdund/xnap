package ies

import (
	"fmt"
	"io"

	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

// XnUAddressIndication ::= SEQUENCE {
//	protocolIEs  ProtocolIE-Container {{XnUAddressIndication-IEs}},
//	...
// }

type XnUAddressIndication struct {
	NewNGRANnodeUEXnAPID              NGRANnodeUEXnAPID
	OldNGRANnodeUEXnAPID              NGRANnodeUEXnAPID
	XnUAddressInfoperPDUSessionList   XnUAddressInfoperPDUSessionList
	CHOMRDCIndicator                  *CHOMRDCIndicator
	CHOMRDCEarlyDataForwarding        *CHOMRDCEarlyDataForwarding
	CPCDataForwardingIndicator        *CPCDataForwardingIndicator
	MBSDataForwardingIndicator        *MBSDataForwardingIndicator
	MBSSessionInformationResponseList *MBSSessionInformationResponseList
	PDUSetbasedHandlingIndicator      *PDUSetbasedHandlingIndicator
}

func (msg *XnUAddressIndication) toIEs() []XnAPMessageIE {
	ies := make([]XnAPMessageIE, 0)
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_NewNGRANnodeUEXnAPID)},
		Criticality: Criticality{Value: CriticalityIgnore},
		Value:       &msg.NewNGRANnodeUEXnAPID,
	})
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_OldNGRANnodeUEXnAPID)},
		Criticality: Criticality{Value: CriticalityIgnore},
		Value:       &msg.OldNGRANnodeUEXnAPID,
	})
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_XnUAddressInfoperPDUSessionList)},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.XnUAddressInfoperPDUSessionList,
	})
	if msg.CHOMRDCIndicator != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_CHOMRDCIndicator)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.CHOMRDCIndicator,
		})
	}
	if msg.CHOMRDCEarlyDataForwarding != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_CHOMRDCEarlyDataForwarding)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.CHOMRDCEarlyDataForwarding,
		})
	}
	if msg.CPCDataForwardingIndicator != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_CPCDataForwardingIndicator)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.CPCDataForwardingIndicator,
		})
	}
	if msg.MBSDataForwardingIndicator != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_MBSDataForwardingIndicator)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.MBSDataForwardingIndicator,
		})
	}
	if msg.MBSSessionInformationResponseList != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_MBSSessionInformationResponseList)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.MBSSessionInformationResponseList,
		})
	}
	if msg.PDUSetbasedHandlingIndicator != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_PDUSetbasedHandlingIndicator)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.PDUSetbasedHandlingIndicator,
		})
	}
	return ies
}

func (msg *XnUAddressIndication) Encode(w io.Writer) error {
	inner := per.NewEncoder(per.APER)
	if err := encodeBody(msg, inner); err != nil {
		return err
	}

	outer := per.NewEncoder(per.APER)

	if err := choiceXnMsg(outer, common.InitiatingMessage); err != nil {
		return err
	}
	if err := procedureCodeXnMsg(outer, common.ProcedureCode_XnUAddressIndication); err != nil {
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

func (msg *XnUAddressIndication) Decode(wire []byte) error {
	d := per.NewDecoder(wire, per.APER)

	seq := d.NewSequenceDecoder(per.SequenceConstraints{Extensible: true})
	if err := seq.DecodeExtensionBit(); err != nil {
		return fmt.Errorf("decode XnUAddressIndication extension bit: %w", err)
	}

	seqOf := d.NewSequenceOfDecoder(per.SizeConstraints{
		Extensible: false,
		Min:        common.Ptr(int64(0)),
		Max:        common.Ptr(int64(MaxProtocolIEs)),
	})
	n, err := seqOf.DecodeLength()
	if err != nil {
		return fmt.Errorf("decode XnUAddressIndication protocolIEs length: %w", err)
	}

	seen := make(map[int64]bool, n)
	for i := range n {
		id, err := d.DecodeInteger(protocolIEIDConstraints)
		if err != nil {
			return fmt.Errorf("decode XnUAddressIndication IE[%d] id: %w", i, err)
		}
		seen[id] = true

		var criticality Criticality
		if err := criticality.Decode(d); err != nil {
			return fmt.Errorf("decode XnUAddressIndication IE[%d] criticality: %w", i, err)
		}

		payload, err := d.DecodeOpenType()
		if err != nil {
			return fmt.Errorf("decode XnUAddressIndication IE[%d] value: %w", i, err)
		}
		ies := per.NewDecoder(payload, per.APER)

		switch id {
		case common.ProtocolIEID_NewNGRANnodeUEXnAPID:
			err = msg.NewNGRANnodeUEXnAPID.Decode(ies)
		case common.ProtocolIEID_OldNGRANnodeUEXnAPID:
			err = msg.OldNGRANnodeUEXnAPID.Decode(ies)
		case common.ProtocolIEID_XnUAddressInfoperPDUSessionList:
			err = msg.XnUAddressInfoperPDUSessionList.Decode(ies)
		case common.ProtocolIEID_CHOMRDCIndicator:
			v := &CHOMRDCIndicator{}
			err = v.Decode(ies)
			msg.CHOMRDCIndicator = v
		case common.ProtocolIEID_CHOMRDCEarlyDataForwarding:
			v := &CHOMRDCEarlyDataForwarding{}
			err = v.Decode(ies)
			msg.CHOMRDCEarlyDataForwarding = v
		case common.ProtocolIEID_CPCDataForwardingIndicator:
			v := &CPCDataForwardingIndicator{}
			err = v.Decode(ies)
			msg.CPCDataForwardingIndicator = v
		case common.ProtocolIEID_MBSDataForwardingIndicator:
			v := &MBSDataForwardingIndicator{}
			err = v.Decode(ies)
			msg.MBSDataForwardingIndicator = v
		case common.ProtocolIEID_MBSSessionInformationResponseList:
			v := &MBSSessionInformationResponseList{}
			err = v.Decode(ies)
			msg.MBSSessionInformationResponseList = v
		case common.ProtocolIEID_PDUSetbasedHandlingIndicator:
			v := &PDUSetbasedHandlingIndicator{}
			err = v.Decode(ies)
			msg.PDUSetbasedHandlingIndicator = v
		default:
			if criticality.Value == CriticalityReject {
				return fmt.Errorf("unsupported XnUAddressIndication IE id=%d with reject criticality", id)
			}
		}
		if err != nil {
			return fmt.Errorf("decode XnUAddressIndication IE id=%d: %w", id, err)
		}
	}

	if _, err := seq.DecodeExtensionAdditions(); err != nil {
		return fmt.Errorf("decode XnUAddressIndication extensions: %w", err)
	}

	for _, id := range []int64{
		common.ProtocolIEID_NewNGRANnodeUEXnAPID,
		common.ProtocolIEID_OldNGRANnodeUEXnAPID,
		common.ProtocolIEID_XnUAddressInfoperPDUSessionList,
	} {
		if !seen[id] {
			return fmt.Errorf("missing mandatory XnUAddressIndication IE id=%d", id)
		}
	}

	return nil
}
