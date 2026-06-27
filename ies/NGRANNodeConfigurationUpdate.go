package ies

import (
	"fmt"
	"io"

	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

// NGRANNodeConfigurationUpdate ::= SEQUENCE {
//	protocolIEs  ProtocolIE-Container {{NGRANNodeConfigurationUpdate-IEs}},
//	...
// }

type NGRANNodeConfigurationUpdate struct {
	TAISupportList                          *TAISupportList
	ConfigurationUpdateInitiatingNodeChoice ConfigurationUpdateInitiatingNodeChoice
	TNLAToAddList                           *TNLAToAddList
	TNLAToRemoveList                        *TNLAToRemoveList
	TNLAToUpdateList                        *TNLAToUpdateList
	GlobalNGRANNodeID                       *GlobalNGRANNodeID
	AMFRegionInformationToAdd               *AMFRegionInformation
	AMFRegionInformationToDelete            *AMFRegionInformation
	InterfaceInstanceIndication             *InterfaceInstanceIndication
	TNLConfigurationInfo                    *TNLConfigurationInfo
	CoverageModificationList                *CoverageModificationList
	LocalNGRANNodeIdentifier                *LocalNGRANNodeIdentifier
	NeighbourNGRANNodeList                  *NeighbourNGRANNodeList
	LocalNGRANNodeIdentifierRemoval         *LocalNGRANNodeIdentifier
}

func (msg *NGRANNodeConfigurationUpdate) toIEs() []XnAPMessageIE {
	ies := make([]XnAPMessageIE, 0)
	if msg.TAISupportList != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_TAISupportList)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.TAISupportList,
		})
	}
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_ConfigurationUpdateInitiatingNodeChoice)},
		Criticality: Criticality{Value: CriticalityIgnore},
		Value:       &msg.ConfigurationUpdateInitiatingNodeChoice,
	})
	if msg.TNLAToAddList != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_TNLAToAddList)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.TNLAToAddList,
		})
	}
	if msg.TNLAToRemoveList != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_TNLAToRemoveList)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.TNLAToRemoveList,
		})
	}
	if msg.TNLAToUpdateList != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_TNLAToUpdateList)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.TNLAToUpdateList,
		})
	}
	if msg.GlobalNGRANNodeID != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_GlobalNGRANNodeID)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.GlobalNGRANNodeID,
		})
	}
	if msg.AMFRegionInformationToAdd != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_AMFRegionInformationToAdd)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.AMFRegionInformationToAdd,
		})
	}
	if msg.AMFRegionInformationToDelete != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_AMFRegionInformationToDelete)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.AMFRegionInformationToDelete,
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
	if msg.CoverageModificationList != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_CoverageModificationList)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.CoverageModificationList,
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

func (msg *NGRANNodeConfigurationUpdate) Encode(w io.Writer) error {
	inner := per.NewEncoder(per.APER)
	if err := encodeBody(msg, inner); err != nil {
		return err
	}

	outer := per.NewEncoder(per.APER)

	if err := choiceXnMsg(outer, common.InitiatingMessage); err != nil {
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

func (msg *NGRANNodeConfigurationUpdate) Decode(wire []byte) error {
	d := per.NewDecoder(wire, per.APER)

	seq := d.NewSequenceDecoder(per.SequenceConstraints{Extensible: true})
	if err := seq.DecodeExtensionBit(); err != nil {
		return fmt.Errorf("decode NGRANNodeConfigurationUpdate extension bit: %w", err)
	}

	seqOf := d.NewSequenceOfDecoder(per.SizeConstraints{
		Extensible: false,
		Min:        common.Ptr(int64(0)),
		Max:        common.Ptr(int64(MaxProtocolIEs)),
	})
	n, err := seqOf.DecodeLength()
	if err != nil {
		return fmt.Errorf("decode NGRANNodeConfigurationUpdate protocolIEs length: %w", err)
	}

	seen := make(map[int64]bool, n)
	for i := range n {
		id, err := d.DecodeInteger(protocolIEIDConstraints)
		if err != nil {
			return fmt.Errorf("decode NGRANNodeConfigurationUpdate IE[%d] id: %w", i, err)
		}
		seen[id] = true

		var criticality Criticality
		if err := criticality.Decode(d); err != nil {
			return fmt.Errorf("decode NGRANNodeConfigurationUpdate IE[%d] criticality: %w", i, err)
		}

		payload, err := d.DecodeOpenType()
		if err != nil {
			return fmt.Errorf("decode NGRANNodeConfigurationUpdate IE[%d] value: %w", i, err)
		}
		ies := per.NewDecoder(payload, per.APER)

		switch id {
		case common.ProtocolIEID_TAISupportList:
			v := &TAISupportList{}
			err = v.Decode(ies)
			msg.TAISupportList = v
		case common.ProtocolIEID_ConfigurationUpdateInitiatingNodeChoice:
			err = msg.ConfigurationUpdateInitiatingNodeChoice.Decode(ies)
		case common.ProtocolIEID_TNLAToAddList:
			v := &TNLAToAddList{}
			err = v.Decode(ies)
			msg.TNLAToAddList = v
		case common.ProtocolIEID_TNLAToRemoveList:
			v := &TNLAToRemoveList{}
			err = v.Decode(ies)
			msg.TNLAToRemoveList = v
		case common.ProtocolIEID_TNLAToUpdateList:
			v := &TNLAToUpdateList{}
			err = v.Decode(ies)
			msg.TNLAToUpdateList = v
		case common.ProtocolIEID_GlobalNGRANNodeID:
			v := &GlobalNGRANNodeID{}
			err = v.Decode(ies)
			msg.GlobalNGRANNodeID = v
		case common.ProtocolIEID_AMFRegionInformationToAdd:
			v := &AMFRegionInformation{}
			err = v.Decode(ies)
			msg.AMFRegionInformationToAdd = v
		case common.ProtocolIEID_AMFRegionInformationToDelete:
			v := &AMFRegionInformation{}
			err = v.Decode(ies)
			msg.AMFRegionInformationToDelete = v
		case common.ProtocolIEID_InterfaceInstanceIndication:
			v := &InterfaceInstanceIndication{}
			err = v.Decode(ies)
			msg.InterfaceInstanceIndication = v
		case common.ProtocolIEID_TNLConfigurationInfo:
			v := &TNLConfigurationInfo{}
			err = v.Decode(ies)
			msg.TNLConfigurationInfo = v
		case common.ProtocolIEID_CoverageModificationList:
			v := &CoverageModificationList{}
			err = v.Decode(ies)
			msg.CoverageModificationList = v
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
				return fmt.Errorf("unsupported NGRANNodeConfigurationUpdate IE id=%d with reject criticality", id)
			}
		}
		if err != nil {
			return fmt.Errorf("decode NGRANNodeConfigurationUpdate IE id=%d: %w", id, err)
		}
	}

	if _, err := seq.DecodeExtensionAdditions(); err != nil {
		return fmt.Errorf("decode NGRANNodeConfigurationUpdate extensions: %w", err)
	}

	for _, id := range []int64{
		common.ProtocolIEID_ConfigurationUpdateInitiatingNodeChoice,
	} {
		if !seen[id] {
			return fmt.Errorf("missing mandatory NGRANNodeConfigurationUpdate IE id=%d", id)
		}
	}

	return nil
}
