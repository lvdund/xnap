package ies

import (
	"fmt"
	"io"

	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

// XnSetupResponse ::= SEQUENCE {
//	protocolIEs  ProtocolIE-Container {{XnSetupResponse-IEs}},
//	...
// }

type XnSetupResponse struct {
	GlobalNGRANNodeID                  GlobalNGRANNodeID
	TAISupportList                     TAISupportList
	ListOfServedCellsNR                *ServedCellsNR
	ListOfServedCellsEUTRA             *ServedCellsEUTRA
	CriticalityDiagnostics             *CriticalityDiagnostics
	AMFRegionInformation               *AMFRegionInformation
	InterfaceInstanceIndication        *InterfaceInstanceIndication
	TNLConfigurationInfo               *TNLConfigurationInfo
	PartialListIndicatorNR             *PartialListIndicator
	CellAndCapacityAssistanceInfoNR    *CellAndCapacityAssistanceInfoNR
	PartialListIndicatorEUTRA          *PartialListIndicator
	CellAndCapacityAssistanceInfoEUTRA *CellAndCapacityAssistanceInfoEUTRA
	LocalNGRANNodeIdentifier           *LocalNGRANNodeIdentifier
	NeighbourNGRANNodeList             *NeighbourNGRANNodeList
}

func (msg *XnSetupResponse) toIEs() []XnAPMessageIE {
	ies := make([]XnAPMessageIE, 0)
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_GlobalNGRANNodeID)},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.GlobalNGRANNodeID,
	})
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_TAISupportList)},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.TAISupportList,
	})
	if msg.ListOfServedCellsNR != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_ListOfServedCellsNR)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.ListOfServedCellsNR,
		})
	}
	if msg.ListOfServedCellsEUTRA != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_ListOfServedCellsEUTRA)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.ListOfServedCellsEUTRA,
		})
	}
	if msg.CriticalityDiagnostics != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_CriticalityDiagnostics)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.CriticalityDiagnostics,
		})
	}
	if msg.AMFRegionInformation != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_AMFRegionInformation)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.AMFRegionInformation,
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
	if msg.PartialListIndicatorNR != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_PartialListIndicatorNR)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.PartialListIndicatorNR,
		})
	}
	if msg.CellAndCapacityAssistanceInfoNR != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_CellAndCapacityAssistanceInfoNR)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.CellAndCapacityAssistanceInfoNR,
		})
	}
	if msg.PartialListIndicatorEUTRA != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_PartialListIndicatorEUTRA)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.PartialListIndicatorEUTRA,
		})
	}
	if msg.CellAndCapacityAssistanceInfoEUTRA != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_CellAndCapacityAssistanceInfoEUTRA)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.CellAndCapacityAssistanceInfoEUTRA,
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
	return ies
}

func (msg *XnSetupResponse) Encode(w io.Writer) error {
	inner := per.NewEncoder(per.APER)
	if err := encodeBody(msg, inner); err != nil {
		return err
	}

	outer := per.NewEncoder(per.APER)

	if err := choiceXnMsg(outer, common.SuccessfulOutcome); err != nil {
		return err
	}
	if err := procedureCodeXnMsg(outer, common.ProcedureCode_XnSetup); err != nil {
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

func (msg *XnSetupResponse) Decode(wire []byte) error {
	d := per.NewDecoder(wire, per.APER)

	seq := d.NewSequenceDecoder(per.SequenceConstraints{Extensible: true})
	if err := seq.DecodeExtensionBit(); err != nil {
		return fmt.Errorf("decode XnSetupResponse extension bit: %w", err)
	}

	seqOf := d.NewSequenceOfDecoder(per.SizeConstraints{
		Extensible: false,
		Min:        common.Ptr(int64(0)),
		Max:        common.Ptr(int64(MaxProtocolIEs)),
	})
	n, err := seqOf.DecodeLength()
	if err != nil {
		return fmt.Errorf("decode XnSetupResponse protocolIEs length: %w", err)
	}

	seen := make(map[int64]bool, n)
	for i := range n {
		id, err := d.DecodeInteger(protocolIEIDConstraints)
		if err != nil {
			return fmt.Errorf("decode XnSetupResponse IE[%d] id: %w", i, err)
		}
		seen[id] = true

		var criticality Criticality
		if err := criticality.Decode(d); err != nil {
			return fmt.Errorf("decode XnSetupResponse IE[%d] criticality: %w", i, err)
		}

		payload, err := d.DecodeOpenType()
		if err != nil {
			return fmt.Errorf("decode XnSetupResponse IE[%d] value: %w", i, err)
		}
		ies := per.NewDecoder(payload, per.APER)

		switch id {
		case common.ProtocolIEID_GlobalNGRANNodeID:
			err = msg.GlobalNGRANNodeID.Decode(ies)
		case common.ProtocolIEID_TAISupportList:
			err = msg.TAISupportList.Decode(ies)
		case common.ProtocolIEID_ListOfServedCellsNR:
			v := &ServedCellsNR{}
			err = v.Decode(ies)
			msg.ListOfServedCellsNR = v
		case common.ProtocolIEID_ListOfServedCellsEUTRA:
			v := &ServedCellsEUTRA{}
			err = v.Decode(ies)
			msg.ListOfServedCellsEUTRA = v
		case common.ProtocolIEID_CriticalityDiagnostics:
			v := &CriticalityDiagnostics{}
			err = v.Decode(ies)
			msg.CriticalityDiagnostics = v
		case common.ProtocolIEID_AMFRegionInformation:
			v := &AMFRegionInformation{}
			err = v.Decode(ies)
			msg.AMFRegionInformation = v
		case common.ProtocolIEID_InterfaceInstanceIndication:
			v := &InterfaceInstanceIndication{}
			err = v.Decode(ies)
			msg.InterfaceInstanceIndication = v
		case common.ProtocolIEID_TNLConfigurationInfo:
			v := &TNLConfigurationInfo{}
			err = v.Decode(ies)
			msg.TNLConfigurationInfo = v
		case common.ProtocolIEID_PartialListIndicatorNR:
			v := &PartialListIndicator{}
			err = v.Decode(ies)
			msg.PartialListIndicatorNR = v
		case common.ProtocolIEID_CellAndCapacityAssistanceInfoNR:
			v := &CellAndCapacityAssistanceInfoNR{}
			err = v.Decode(ies)
			msg.CellAndCapacityAssistanceInfoNR = v
		case common.ProtocolIEID_PartialListIndicatorEUTRA:
			v := &PartialListIndicator{}
			err = v.Decode(ies)
			msg.PartialListIndicatorEUTRA = v
		case common.ProtocolIEID_CellAndCapacityAssistanceInfoEUTRA:
			v := &CellAndCapacityAssistanceInfoEUTRA{}
			err = v.Decode(ies)
			msg.CellAndCapacityAssistanceInfoEUTRA = v
		case common.ProtocolIEID_LocalNGRANNodeIdentifier:
			v := &LocalNGRANNodeIdentifier{}
			err = v.Decode(ies)
			msg.LocalNGRANNodeIdentifier = v
		case common.ProtocolIEID_NeighbourNGRANNodeList:
			v := &NeighbourNGRANNodeList{}
			err = v.Decode(ies)
			msg.NeighbourNGRANNodeList = v
		default:
			if criticality.Value == CriticalityReject {
				return fmt.Errorf("unsupported XnSetupResponse IE id=%d with reject criticality", id)
			}
		}
		if err != nil {
			return fmt.Errorf("decode XnSetupResponse IE id=%d: %w", id, err)
		}
	}

	if _, err := seq.DecodeExtensionAdditions(); err != nil {
		return fmt.Errorf("decode XnSetupResponse extensions: %w", err)
	}

	for _, id := range []int64{
		common.ProtocolIEID_GlobalNGRANNodeID,
		common.ProtocolIEID_TAISupportList,
	} {
		if !seen[id] {
			return fmt.Errorf("missing mandatory XnSetupResponse IE id=%d", id)
		}
	}

	return nil
}
