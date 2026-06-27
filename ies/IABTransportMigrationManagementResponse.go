package ies

import (
	"fmt"
	"io"

	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

// IABTransportMigrationManagementResponse ::= SEQUENCE {
//	protocolIEs  ProtocolIE-Container {{IABTransportMigrationManagementResponse-IEs}},
//	...
// }

type IABTransportMigrationManagementResponse struct {
	F1TerminatingIABDonorUEXnAPID    NGRANnodeUEXnAPID
	NonF1TerminatingIABDonorUEXnAPID NGRANnodeUEXnAPID
	TrafficAddedList                 *TrafficAddedList
	TrafficModifiedList              *TrafficModifiedList
	TrafficNotAddedList              *TrafficNotAddedList
	TrafficNotModifiedList           *TrafficNotModifiedList
	IABTNLAddressResponse            *IABTNLAddressResponse
	TrafficReleasedList              *TrafficReleasedList
}

func (msg *IABTransportMigrationManagementResponse) toIEs() []XnAPMessageIE {
	ies := make([]XnAPMessageIE, 0)
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_F1TerminatingIABDonorUEXnAPID)},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.F1TerminatingIABDonorUEXnAPID,
	})
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_NonF1TerminatingIABDonorUEXnAPID)},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.NonF1TerminatingIABDonorUEXnAPID,
	})
	if msg.TrafficAddedList != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_TrafficAddedList)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.TrafficAddedList,
		})
	}
	if msg.TrafficModifiedList != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_TrafficModifiedList)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.TrafficModifiedList,
		})
	}
	if msg.TrafficNotAddedList != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_TrafficNotAddedList)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.TrafficNotAddedList,
		})
	}
	if msg.TrafficNotModifiedList != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_TrafficNotModifiedList)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.TrafficNotModifiedList,
		})
	}
	if msg.IABTNLAddressResponse != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_IABTNLAddressResponse)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.IABTNLAddressResponse,
		})
	}
	if msg.TrafficReleasedList != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_TrafficReleasedList)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.TrafficReleasedList,
		})
	}
	return ies
}

func (msg *IABTransportMigrationManagementResponse) Encode(w io.Writer) error {
	inner := per.NewEncoder(per.APER)
	if err := encodeBody(msg, inner); err != nil {
		return err
	}

	outer := per.NewEncoder(per.APER)

	if err := choiceXnMsg(outer, common.SuccessfulOutcome); err != nil {
		return err
	}
	if err := procedureCodeXnMsg(outer, common.ProcedureCode_IABTransportMigrationManagement); err != nil {
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

func (msg *IABTransportMigrationManagementResponse) Decode(wire []byte) error {
	d := per.NewDecoder(wire, per.APER)

	seq := d.NewSequenceDecoder(per.SequenceConstraints{Extensible: true})
	if err := seq.DecodeExtensionBit(); err != nil {
		return fmt.Errorf("decode IABTransportMigrationManagementResponse extension bit: %w", err)
	}

	seqOf := d.NewSequenceOfDecoder(per.SizeConstraints{
		Extensible: false,
		Min:        common.Ptr(int64(0)),
		Max:        common.Ptr(int64(MaxProtocolIEs)),
	})
	n, err := seqOf.DecodeLength()
	if err != nil {
		return fmt.Errorf("decode IABTransportMigrationManagementResponse protocolIEs length: %w", err)
	}

	seen := make(map[int64]bool, n)
	for i := range n {
		id, err := d.DecodeInteger(protocolIEIDConstraints)
		if err != nil {
			return fmt.Errorf("decode IABTransportMigrationManagementResponse IE[%d] id: %w", i, err)
		}
		seen[id] = true

		var criticality Criticality
		if err := criticality.Decode(d); err != nil {
			return fmt.Errorf("decode IABTransportMigrationManagementResponse IE[%d] criticality: %w", i, err)
		}

		payload, err := d.DecodeOpenType()
		if err != nil {
			return fmt.Errorf("decode IABTransportMigrationManagementResponse IE[%d] value: %w", i, err)
		}
		ies := per.NewDecoder(payload, per.APER)

		switch id {
		case common.ProtocolIEID_F1TerminatingIABDonorUEXnAPID:
			err = msg.F1TerminatingIABDonorUEXnAPID.Decode(ies)
		case common.ProtocolIEID_NonF1TerminatingIABDonorUEXnAPID:
			err = msg.NonF1TerminatingIABDonorUEXnAPID.Decode(ies)
		case common.ProtocolIEID_TrafficAddedList:
			v := &TrafficAddedList{}
			err = v.Decode(ies)
			msg.TrafficAddedList = v
		case common.ProtocolIEID_TrafficModifiedList:
			v := &TrafficModifiedList{}
			err = v.Decode(ies)
			msg.TrafficModifiedList = v
		case common.ProtocolIEID_TrafficNotAddedList:
			v := &TrafficNotAddedList{}
			err = v.Decode(ies)
			msg.TrafficNotAddedList = v
		case common.ProtocolIEID_TrafficNotModifiedList:
			v := &TrafficNotModifiedList{}
			err = v.Decode(ies)
			msg.TrafficNotModifiedList = v
		case common.ProtocolIEID_IABTNLAddressResponse:
			v := &IABTNLAddressResponse{}
			err = v.Decode(ies)
			msg.IABTNLAddressResponse = v
		case common.ProtocolIEID_TrafficReleasedList:
			v := &TrafficReleasedList{}
			err = v.Decode(ies)
			msg.TrafficReleasedList = v
		default:
			if criticality.Value == CriticalityReject {
				return fmt.Errorf("unsupported IABTransportMigrationManagementResponse IE id=%d with reject criticality", id)
			}
		}
		if err != nil {
			return fmt.Errorf("decode IABTransportMigrationManagementResponse IE id=%d: %w", id, err)
		}
	}

	if _, err := seq.DecodeExtensionAdditions(); err != nil {
		return fmt.Errorf("decode IABTransportMigrationManagementResponse extensions: %w", err)
	}

	for _, id := range []int64{
		common.ProtocolIEID_F1TerminatingIABDonorUEXnAPID,
		common.ProtocolIEID_NonF1TerminatingIABDonorUEXnAPID,
	} {
		if !seen[id] {
			return fmt.Errorf("missing mandatory IABTransportMigrationManagementResponse IE id=%d", id)
		}
	}

	return nil
}
