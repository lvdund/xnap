package ies

import (
	"fmt"
	"io"

	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

// SNodeChangeRequired ::= SEQUENCE {
//	protocolIEs  ProtocolIE-Container {{SNodeChangeRequired-IEs}},
//	...
// }

type SNodeChangeRequired struct {
	MNGRANnodeUEXnAPID             NGRANnodeUEXnAPID
	SNGRANnodeUEXnAPID             NGRANnodeUEXnAPID
	TargetSNGRANnodeID             GlobalNGRANNodeID
	Cause                          Cause
	PDUSessionSNChangeRequiredList *PDUSessionSNChangeRequiredList
	SNToMNContainer                []byte
	SCGUEHistoryInformation        *SCGUEHistoryInformation
	SNMobilityInformation          *SNMobilityInformation
	SourcePSCellID                 *GlobalNGRANCellID
	CPCInformationRequired         *CPCInformationRequired
	SourceSNToTargetSNQMCInfo      *QMCConfigInfo
}

func (msg *SNodeChangeRequired) toIEs() []XnAPMessageIE {
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
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_TargetSNGRANnodeID)},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &msg.TargetSNGRANnodeID,
	})
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_Cause)},
		Criticality: Criticality{Value: CriticalityIgnore},
		Value:       &msg.Cause,
	})
	if msg.PDUSessionSNChangeRequiredList != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_PDUSessionSNChangeRequiredList)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.PDUSessionSNChangeRequiredList,
		})
	}
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SNToMNContainer)},
		Criticality: Criticality{Value: CriticalityReject},
		Value:       &octetStringIE{Value: msg.SNToMNContainer},
	})
	if msg.SCGUEHistoryInformation != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SCGUEHistoryInformation)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.SCGUEHistoryInformation,
		})
	}
	if msg.SNMobilityInformation != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SNMobilityInformation)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.SNMobilityInformation,
		})
	}
	if msg.SourcePSCellID != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SourcePSCellID)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.SourcePSCellID,
		})
	}
	if msg.CPCInformationRequired != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_CPCInformationRequired)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.CPCInformationRequired,
		})
	}
	if msg.SourceSNToTargetSNQMCInfo != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SourceSNToTargetSNQMCInfo)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.SourceSNToTargetSNQMCInfo,
		})
	}
	return ies
}

func (msg *SNodeChangeRequired) Encode(w io.Writer) error {
	inner := per.NewEncoder(per.APER)
	if err := encodeBody(msg, inner); err != nil {
		return err
	}

	outer := per.NewEncoder(per.APER)

	if err := choiceXnMsg(outer, common.InitiatingMessage); err != nil {
		return err
	}
	if err := procedureCodeXnMsg(outer, common.ProcedureCode_SNGRANnodeChange); err != nil {
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

func (msg *SNodeChangeRequired) Decode(wire []byte) error {
	d := per.NewDecoder(wire, per.APER)

	seq := d.NewSequenceDecoder(per.SequenceConstraints{Extensible: true})
	if err := seq.DecodeExtensionBit(); err != nil {
		return fmt.Errorf("decode SNodeChangeRequired extension bit: %w", err)
	}

	seqOf := d.NewSequenceOfDecoder(per.SizeConstraints{
		Extensible: false,
		Min:        common.Ptr(int64(0)),
		Max:        common.Ptr(int64(MaxProtocolIEs)),
	})
	n, err := seqOf.DecodeLength()
	if err != nil {
		return fmt.Errorf("decode SNodeChangeRequired protocolIEs length: %w", err)
	}

	seen := make(map[int64]bool, n)
	for i := range n {
		id, err := d.DecodeInteger(protocolIEIDConstraints)
		if err != nil {
			return fmt.Errorf("decode SNodeChangeRequired IE[%d] id: %w", i, err)
		}
		seen[id] = true

		var criticality Criticality
		if err := criticality.Decode(d); err != nil {
			return fmt.Errorf("decode SNodeChangeRequired IE[%d] criticality: %w", i, err)
		}

		payload, err := d.DecodeOpenType()
		if err != nil {
			return fmt.Errorf("decode SNodeChangeRequired IE[%d] value: %w", i, err)
		}
		ies := per.NewDecoder(payload, per.APER)

		switch id {
		case common.ProtocolIEID_MNGRANnodeUEXnAPID:
			err = msg.MNGRANnodeUEXnAPID.Decode(ies)
		case common.ProtocolIEID_SNGRANnodeUEXnAPID:
			err = msg.SNGRANnodeUEXnAPID.Decode(ies)
		case common.ProtocolIEID_TargetSNGRANnodeID:
			err = msg.TargetSNGRANnodeID.Decode(ies)
		case common.ProtocolIEID_Cause:
			err = msg.Cause.Decode(ies)
		case common.ProtocolIEID_PDUSessionSNChangeRequiredList:
			v := &PDUSessionSNChangeRequiredList{}
			err = v.Decode(ies)
			msg.PDUSessionSNChangeRequiredList = v
		case common.ProtocolIEID_SNToMNContainer:
			v := octetStringIE{}
			err = v.Decode(ies)
			msg.SNToMNContainer = v.Value
		case common.ProtocolIEID_SCGUEHistoryInformation:
			v := &SCGUEHistoryInformation{}
			err = v.Decode(ies)
			msg.SCGUEHistoryInformation = v
		case common.ProtocolIEID_SNMobilityInformation:
			v := &SNMobilityInformation{}
			err = v.Decode(ies)
			msg.SNMobilityInformation = v
		case common.ProtocolIEID_SourcePSCellID:
			v := &GlobalNGRANCellID{}
			err = v.Decode(ies)
			msg.SourcePSCellID = v
		case common.ProtocolIEID_CPCInformationRequired:
			v := &CPCInformationRequired{}
			err = v.Decode(ies)
			msg.CPCInformationRequired = v
		case common.ProtocolIEID_SourceSNToTargetSNQMCInfo:
			v := &QMCConfigInfo{}
			err = v.Decode(ies)
			msg.SourceSNToTargetSNQMCInfo = v
		default:
			if criticality.Value == CriticalityReject {
				return fmt.Errorf("unsupported SNodeChangeRequired IE id=%d with reject criticality", id)
			}
		}
		if err != nil {
			return fmt.Errorf("decode SNodeChangeRequired IE id=%d: %w", id, err)
		}
	}

	if _, err := seq.DecodeExtensionAdditions(); err != nil {
		return fmt.Errorf("decode SNodeChangeRequired extensions: %w", err)
	}

	for _, id := range []int64{
		common.ProtocolIEID_MNGRANnodeUEXnAPID,
		common.ProtocolIEID_SNGRANnodeUEXnAPID,
		common.ProtocolIEID_TargetSNGRANnodeID,
		common.ProtocolIEID_Cause,
		common.ProtocolIEID_SNToMNContainer,
	} {
		if !seen[id] {
			return fmt.Errorf("missing mandatory SNodeChangeRequired IE id=%d", id)
		}
	}

	return nil
}
