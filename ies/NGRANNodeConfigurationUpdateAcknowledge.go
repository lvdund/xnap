package ies

import (
	"fmt"
	"io"

	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

// NGRANNodeConfigurationUpdateAcknowledge ::= SEQUENCE {
//	protocolIEs  ProtocolIE-Container {{NGRANNodeConfigurationUpdateAcknowledge-IEs}},
//	...
// }

type NGRANNodeConfigurationUpdateAcknowledge struct {
	RespondingNodeTypeConfigUpdateAck RespondingNodeTypeConfigUpdateAck
	TNLASetupList                     *TNLASetupList
	TNLAFailedToSetupList             *TNLAFailedToSetupList
	CriticalityDiagnostics            *CriticalityDiagnostics
	InterfaceInstanceIndication       *InterfaceInstanceIndication
	TNLConfigurationInfo              *TNLConfigurationInfo
	LocalNGRANNodeIdentifier          *LocalNGRANNodeIdentifier
	NeighbourNGRANNodeList            *NeighbourNGRANNodeList
	LocalNGRANNodeIdentifierRemoval   *LocalNGRANNodeIdentifier
}

func (msg *NGRANNodeConfigurationUpdateAcknowledge) toIEs() []XnAPMessageIE {
	ies := make([]XnAPMessageIE, 0)
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_RespondingNodeTypeConfigUpdateAck)},
		Criticality: Criticality{Value: CriticalityIgnore},
		Value:       &msg.RespondingNodeTypeConfigUpdateAck,
	})
	if msg.TNLASetupList != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_TNLASetupList)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.TNLASetupList,
		})
	}
	if msg.TNLAFailedToSetupList != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_TNLAFailedToSetupList)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.TNLAFailedToSetupList,
		})
	}
	if msg.CriticalityDiagnostics != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_CriticalityDiagnostics)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.CriticalityDiagnostics,
		})
	}
	if msg.InterfaceInstanceIndication != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_InterfaceInstanceIndication)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.InterfaceInstanceIndication,
		})
	}
	if msg.TNLConfigurationInfo != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_TNLConfigurationInfo)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.TNLConfigurationInfo,
		})
	}
	if msg.LocalNGRANNodeIdentifier != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_LocalNGRANNodeIdentifier)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.LocalNGRANNodeIdentifier,
		})
	}
	if msg.NeighbourNGRANNodeList != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_NeighbourNGRANNodeList)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.NeighbourNGRANNodeList,
		})
	}
	if msg.LocalNGRANNodeIdentifierRemoval != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_LocalNGRANNodeIdentifierRemoval)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.LocalNGRANNodeIdentifierRemoval,
		})
	}
	return ies
}

func (msg *NGRANNodeConfigurationUpdateAcknowledge) Encode(w io.Writer) error {
	inner := per.NewEncoder(per.APER)
	if err := encodeBody(msg, inner); err != nil {
		return err
	}

	outer := per.NewEncoder(per.APER)

	if err := choiceXnMsg(outer, common.SuccessfulOutcome); err != nil {
		return err
	}
	if err := procedureCodeXnMsg(outer, common.ProcedureCode_NGRANnodeConfigurationUpdate); err != nil {
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

func (msg *NGRANNodeConfigurationUpdateAcknowledge) Decode(wire []byte) error {
	d := per.NewDecoder(wire, per.APER)

	seq := d.NewSequenceDecoder(per.SequenceConstraints{Extensible: true})
	if err := seq.DecodeExtensionBit(); err != nil {
		return fmt.Errorf("decode NGRANNodeConfigurationUpdateAcknowledge extension bit: %w", err)
	}

	seqOf := d.NewSequenceOfDecoder(per.SizeConstraints{
		Extensible: false,
		Min:        common.Ptr(int64(0)),
		Max:        common.Ptr(int64(MaxProtocolIEs)),
	})
	n, err := seqOf.DecodeLength()
	if err != nil {
		return fmt.Errorf("decode NGRANNodeConfigurationUpdateAcknowledge protocolIEs length: %w", err)
	}

	seen := make(map[int64]bool, n)
	for i := range n {
		id, err := d.DecodeInteger(protocolIEIDConstraints)
		if err != nil {
			return fmt.Errorf("decode NGRANNodeConfigurationUpdateAcknowledge IE[%d] id: %w", i, err)
		}
		seen[id] = true

		var criticality Criticality
		if err := criticality.Decode(d); err != nil {
			return fmt.Errorf("decode NGRANNodeConfigurationUpdateAcknowledge IE[%d] criticality: %w", i, err)
		}

		payload, err := d.DecodeOpenType()
		if err != nil {
			return fmt.Errorf("decode NGRANNodeConfigurationUpdateAcknowledge IE[%d] value: %w", i, err)
		}
		ies := per.NewDecoder(payload, per.APER)

		switch id {
		case common.ProtocolIEID_RespondingNodeTypeConfigUpdateAck:
			err = msg.RespondingNodeTypeConfigUpdateAck.Decode(ies)
		case common.ProtocolIEID_TNLASetupList:
			v := &TNLASetupList{}
			err = v.Decode(ies)
			msg.TNLASetupList = v
		case common.ProtocolIEID_TNLAFailedToSetupList:
			v := &TNLAFailedToSetupList{}
			err = v.Decode(ies)
			msg.TNLAFailedToSetupList = v
		case common.ProtocolIEID_CriticalityDiagnostics:
			v := &CriticalityDiagnostics{}
			err = v.Decode(ies)
			msg.CriticalityDiagnostics = v
		case common.ProtocolIEID_InterfaceInstanceIndication:
			v := &InterfaceInstanceIndication{}
			err = v.Decode(ies)
			msg.InterfaceInstanceIndication = v
		case common.ProtocolIEID_TNLConfigurationInfo:
			v := &TNLConfigurationInfo{}
			err = v.Decode(ies)
			msg.TNLConfigurationInfo = v
		case common.ProtocolIEID_LocalNGRANNodeIdentifier:
			v := &LocalNGRANNodeIdentifier{}
			err = v.Decode(ies)
			msg.LocalNGRANNodeIdentifier = v
		case common.ProtocolIEID_NeighbourNGRANNodeList:
			v := &NeighbourNGRANNodeList{}
			err = v.Decode(ies)
			msg.NeighbourNGRANNodeList = v
		case common.ProtocolIEID_LocalNGRANNodeIdentifierRemoval:
			v := &LocalNGRANNodeIdentifier{}
			err = v.Decode(ies)
			msg.LocalNGRANNodeIdentifierRemoval = v
		default:
			if criticality.Value == CriticalityReject {
				return fmt.Errorf("unsupported NGRANNodeConfigurationUpdateAcknowledge IE id=%d with reject criticality", id)
			}
		}
		if err != nil {
			return fmt.Errorf("decode NGRANNodeConfigurationUpdateAcknowledge IE id=%d: %w", id, err)
		}
	}

	if _, err := seq.DecodeExtensionAdditions(); err != nil {
		return fmt.Errorf("decode NGRANNodeConfigurationUpdateAcknowledge extensions: %w", err)
	}

	for _, id := range []int64{
		common.ProtocolIEID_RespondingNodeTypeConfigUpdateAck,
	} {
		if !seen[id] {
			return fmt.Errorf("missing mandatory NGRANNodeConfigurationUpdateAcknowledge IE id=%d", id)
		}
	}

	return nil
}
