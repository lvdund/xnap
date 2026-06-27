package ies

import (
	"fmt"
	"io"

	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

// IABTransportMigrationModificationRequest ::= SEQUENCE {
//	protocolIEs  ProtocolIE-Container {{IABTransportMigrationModificationRequest-IEs}},
//	...
// }

type IABTransportMigrationModificationRequest struct {
	F1TerminatingIABDonorUEXnAPID    NGRANnodeUEXnAPID
	NonF1TerminatingIABDonorUEXnAPID NGRANnodeUEXnAPID
	TrafficRequiredToBeModifiedList  *TrafficRequiredToBeModifiedList
	TrafficToBeReleaseInformation    *TrafficToBeReleaseInformation
	IABTNLAddressToBeAdded           *IABTNLAddressResponse
	IABTNLAddressToBeReleasedList    *IABTNLAddressToBeReleasedList
	IABAuthorizationStatus           *IABAuthorizationStatus
	MobileIABAuthorizationStatus     *MobileIABAuthorizationStatus
}

func (msg *IABTransportMigrationModificationRequest) toIEs() []XnAPMessageIE {
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
	if msg.TrafficRequiredToBeModifiedList != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_TrafficRequiredToBeModifiedList)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.TrafficRequiredToBeModifiedList,
		})
	}
	if msg.TrafficToBeReleaseInformation != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_TrafficToBeReleaseInformation)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.TrafficToBeReleaseInformation,
		})
	}
	if msg.IABTNLAddressToBeAdded != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_IABTNLAddressToBeAdded)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.IABTNLAddressToBeAdded,
		})
	}
	if msg.IABTNLAddressToBeReleasedList != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_IABTNLAddressToBeReleasedList)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.IABTNLAddressToBeReleasedList,
		})
	}
	if msg.IABAuthorizationStatus != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_IABAuthorizationStatus)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.IABAuthorizationStatus,
		})
	}
	if msg.MobileIABAuthorizationStatus != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_MobileIABAuthorizationStatus)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.MobileIABAuthorizationStatus,
		})
	}
	return ies
}

func (msg *IABTransportMigrationModificationRequest) Encode(w io.Writer) error {
	inner := per.NewEncoder(per.APER)
	if err := encodeBody(msg, inner); err != nil {
		return err
	}

	outer := per.NewEncoder(per.APER)

	if err := choiceXnMsg(outer, common.InitiatingMessage); err != nil {
		return err
	}
	if err := procedureCodeXnMsg(outer, common.ProcedureCode_IABTransportMigrationModification); err != nil {
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

func (msg *IABTransportMigrationModificationRequest) Decode(wire []byte) error {
	d := per.NewDecoder(wire, per.APER)

	seq := d.NewSequenceDecoder(per.SequenceConstraints{Extensible: true})
	if err := seq.DecodeExtensionBit(); err != nil {
		return fmt.Errorf("decode IABTransportMigrationModificationRequest extension bit: %w", err)
	}

	seqOf := d.NewSequenceOfDecoder(per.SizeConstraints{
		Extensible: false,
		Min:        common.Ptr(int64(0)),
		Max:        common.Ptr(int64(MaxProtocolIEs)),
	})
	n, err := seqOf.DecodeLength()
	if err != nil {
		return fmt.Errorf("decode IABTransportMigrationModificationRequest protocolIEs length: %w", err)
	}

	seen := make(map[int64]bool, n)
	for i := range n {
		id, err := d.DecodeInteger(protocolIEIDConstraints)
		if err != nil {
			return fmt.Errorf("decode IABTransportMigrationModificationRequest IE[%d] id: %w", i, err)
		}
		seen[id] = true

		var criticality Criticality
		if err := criticality.Decode(d); err != nil {
			return fmt.Errorf("decode IABTransportMigrationModificationRequest IE[%d] criticality: %w", i, err)
		}

		payload, err := d.DecodeOpenType()
		if err != nil {
			return fmt.Errorf("decode IABTransportMigrationModificationRequest IE[%d] value: %w", i, err)
		}
		ies := per.NewDecoder(payload, per.APER)

		switch id {
		case common.ProtocolIEID_F1TerminatingIABDonorUEXnAPID:
			err = msg.F1TerminatingIABDonorUEXnAPID.Decode(ies)
		case common.ProtocolIEID_NonF1TerminatingIABDonorUEXnAPID:
			err = msg.NonF1TerminatingIABDonorUEXnAPID.Decode(ies)
		case common.ProtocolIEID_TrafficRequiredToBeModifiedList:
			v := &TrafficRequiredToBeModifiedList{}
			err = v.Decode(ies)
			msg.TrafficRequiredToBeModifiedList = v
		case common.ProtocolIEID_TrafficToBeReleaseInformation:
			v := &TrafficToBeReleaseInformation{}
			err = v.Decode(ies)
			msg.TrafficToBeReleaseInformation = v
		case common.ProtocolIEID_IABTNLAddressToBeAdded:
			v := &IABTNLAddressResponse{}
			err = v.Decode(ies)
			msg.IABTNLAddressToBeAdded = v
		case common.ProtocolIEID_IABTNLAddressToBeReleasedList:
			v := &IABTNLAddressToBeReleasedList{}
			err = v.Decode(ies)
			msg.IABTNLAddressToBeReleasedList = v
		case common.ProtocolIEID_IABAuthorizationStatus:
			v := &IABAuthorizationStatus{}
			err = v.Decode(ies)
			msg.IABAuthorizationStatus = v
		case common.ProtocolIEID_MobileIABAuthorizationStatus:
			v := &MobileIABAuthorizationStatus{}
			err = v.Decode(ies)
			msg.MobileIABAuthorizationStatus = v
		default:
			if criticality.Value == CriticalityReject {
				return fmt.Errorf("unsupported IABTransportMigrationModificationRequest IE id=%d with reject criticality", id)
			}
		}
		if err != nil {
			return fmt.Errorf("decode IABTransportMigrationModificationRequest IE id=%d: %w", id, err)
		}
	}

	if _, err := seq.DecodeExtensionAdditions(); err != nil {
		return fmt.Errorf("decode IABTransportMigrationModificationRequest extensions: %w", err)
	}

	for _, id := range []int64{
		common.ProtocolIEID_F1TerminatingIABDonorUEXnAPID,
		common.ProtocolIEID_NonF1TerminatingIABDonorUEXnAPID,
	} {
		if !seen[id] {
			return fmt.Errorf("missing mandatory IABTransportMigrationModificationRequest IE id=%d", id)
		}
	}

	return nil
}
