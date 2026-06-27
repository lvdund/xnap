package ies

import (
	"fmt"
	"io"

	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

// IABTransportMigrationManagementRequest ::= SEQUENCE {
//	protocolIEs  ProtocolIE-Container {{IABTransportMigrationManagementRequest-IEs}},
//	...
// }

type IABTransportMigrationManagementRequest struct {
	F1TerminatingIABDonorUEXnAPID    NGRANnodeUEXnAPID
	NonF1TerminatingIABDonorUEXnAPID NGRANnodeUEXnAPID
	TrafficToBeAddedList             *TrafficToBeAddedList
	TrafficToBeModifiedList          *TrafficToBeModifiedList
	TrafficToBeReleaseInformation    *TrafficToBeReleaseInformation
	IABTNLAddressRequest             *IABTNLAddressRequest
	IABTNLAddressException           *IABTNLAddressException
	MIABMTBAPAddress                 *BAPAddress
}

func (msg *IABTransportMigrationManagementRequest) toIEs() []XnAPMessageIE {
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
	if msg.TrafficToBeAddedList != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_TrafficToBeAddedList)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.TrafficToBeAddedList,
		})
	}
	if msg.TrafficToBeModifiedList != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_TrafficToBeModifiedList)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.TrafficToBeModifiedList,
		})
	}
	if msg.TrafficToBeReleaseInformation != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_TrafficToBeReleaseInformation)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.TrafficToBeReleaseInformation,
		})
	}
	if msg.IABTNLAddressRequest != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_IABTNLAddressRequest)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.IABTNLAddressRequest,
		})
	}
	if msg.IABTNLAddressException != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_IABTNLAddressException)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.IABTNLAddressException,
		})
	}
	if msg.MIABMTBAPAddress != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_MIABMTBAPAddress)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.MIABMTBAPAddress,
		})
	}
	return ies
}

func (msg *IABTransportMigrationManagementRequest) Encode(w io.Writer) error {
	inner := per.NewEncoder(per.APER)
	if err := encodeBody(msg, inner); err != nil {
		return err
	}

	outer := per.NewEncoder(per.APER)

	if err := choiceXnMsg(outer, common.InitiatingMessage); err != nil {
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

func (msg *IABTransportMigrationManagementRequest) Decode(wire []byte) error {
	d := per.NewDecoder(wire, per.APER)

	seq := d.NewSequenceDecoder(per.SequenceConstraints{Extensible: true})
	if err := seq.DecodeExtensionBit(); err != nil {
		return fmt.Errorf("decode IABTransportMigrationManagementRequest extension bit: %w", err)
	}

	seqOf := d.NewSequenceOfDecoder(per.SizeConstraints{
		Extensible: false,
		Min:        common.Ptr(int64(0)),
		Max:        common.Ptr(int64(MaxProtocolIEs)),
	})
	n, err := seqOf.DecodeLength()
	if err != nil {
		return fmt.Errorf("decode IABTransportMigrationManagementRequest protocolIEs length: %w", err)
	}

	seen := make(map[int64]bool, n)
	for i := range n {
		id, err := d.DecodeInteger(protocolIEIDConstraints)
		if err != nil {
			return fmt.Errorf("decode IABTransportMigrationManagementRequest IE[%d] id: %w", i, err)
		}
		seen[id] = true

		var criticality Criticality
		if err := criticality.Decode(d); err != nil {
			return fmt.Errorf("decode IABTransportMigrationManagementRequest IE[%d] criticality: %w", i, err)
		}

		payload, err := d.DecodeOpenType()
		if err != nil {
			return fmt.Errorf("decode IABTransportMigrationManagementRequest IE[%d] value: %w", i, err)
		}
		ies := per.NewDecoder(payload, per.APER)

		switch id {
		case common.ProtocolIEID_F1TerminatingIABDonorUEXnAPID:
			err = msg.F1TerminatingIABDonorUEXnAPID.Decode(ies)
		case common.ProtocolIEID_NonF1TerminatingIABDonorUEXnAPID:
			err = msg.NonF1TerminatingIABDonorUEXnAPID.Decode(ies)
		case common.ProtocolIEID_TrafficToBeAddedList:
			v := &TrafficToBeAddedList{}
			err = v.Decode(ies)
			msg.TrafficToBeAddedList = v
		case common.ProtocolIEID_TrafficToBeModifiedList:
			v := &TrafficToBeModifiedList{}
			err = v.Decode(ies)
			msg.TrafficToBeModifiedList = v
		case common.ProtocolIEID_TrafficToBeReleaseInformation:
			v := &TrafficToBeReleaseInformation{}
			err = v.Decode(ies)
			msg.TrafficToBeReleaseInformation = v
		case common.ProtocolIEID_IABTNLAddressRequest:
			v := &IABTNLAddressRequest{}
			err = v.Decode(ies)
			msg.IABTNLAddressRequest = v
		case common.ProtocolIEID_IABTNLAddressException:
			v := &IABTNLAddressException{}
			err = v.Decode(ies)
			msg.IABTNLAddressException = v
		case common.ProtocolIEID_MIABMTBAPAddress:
			v := &BAPAddress{}
			err = v.Decode(ies)
			msg.MIABMTBAPAddress = v
		default:
			if criticality.Value == CriticalityReject {
				return fmt.Errorf("unsupported IABTransportMigrationManagementRequest IE id=%d with reject criticality", id)
			}
		}
		if err != nil {
			return fmt.Errorf("decode IABTransportMigrationManagementRequest IE id=%d: %w", id, err)
		}
	}

	if _, err := seq.DecodeExtensionAdditions(); err != nil {
		return fmt.Errorf("decode IABTransportMigrationManagementRequest extensions: %w", err)
	}

	for _, id := range []int64{
		common.ProtocolIEID_F1TerminatingIABDonorUEXnAPID,
		common.ProtocolIEID_NonF1TerminatingIABDonorUEXnAPID,
	} {
		if !seen[id] {
			return fmt.Errorf("missing mandatory IABTransportMigrationManagementRequest IE id=%d", id)
		}
	}

	return nil
}
