package ies

import (
	"fmt"
	"io"

	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

// HandoverRequestAcknowledge ::= SEQUENCE {
//	protocolIEs  ProtocolIE-Container {{HandoverRequestAcknowledge-IEs}},
//	...
// }

type HandoverRequestAcknowledge struct {
	SourceNGRANnodeUEXnAPID               NGRANnodeUEXnAPID
	TargetNGRANnodeUEXnAPID               NGRANnodeUEXnAPID
	PDUSessionResourcesAdmittedList       PDUSessionResourcesAdmittedList
	PDUSessionResourcesNotAdmittedList    *PDUSessionResourcesNotAdmittedList
	Target2SourceNGRANnodeTranspContainer []byte
	UEContextKeptIndicator                *UEContextKeptIndicator
	CriticalityDiagnostics                *CriticalityDiagnostics
	DRBsTransferredToMN                   *DRBList
	DAPSResponseInfoList                  *DAPSResponseInfoList
	CHOinformationAck                     *CHOinformationAck
	MBSSessionInformationResponseList     *MBSSessionInformationResponseList
	RRCConfigIndication                   *RRCConfigIndication
	PDUSetbasedHandlingIndicator          *PDUSetbasedHandlingIndicator
}

func (msg *HandoverRequestAcknowledge) toIEs() []XnAPMessageIE {
	ies := make([]XnAPMessageIE, 0)
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SourceNGRANnodeUEXnAPID)},
		Criticality: Criticality{Value: CriticalityIgnore},
		Value:       &msg.SourceNGRANnodeUEXnAPID,
	})
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_TargetNGRANnodeUEXnAPID)},
		Criticality: Criticality{Value: CriticalityIgnore},
		Value:       &msg.TargetNGRANnodeUEXnAPID,
	})
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_PDUSessionResourcesAdmittedList)},
		Criticality: Criticality{Value: CriticalityIgnore},
		Value:       &msg.PDUSessionResourcesAdmittedList,
	})
	if msg.PDUSessionResourcesNotAdmittedList != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_PDUSessionResourcesNotAdmittedList)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.PDUSessionResourcesNotAdmittedList,
		})
	}
	ies = append(ies, XnAPMessageIE{
		ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_Target2SourceNGRANnodeTranspContainer)},
		Criticality: Criticality{Value: CriticalityIgnore},
		Value:       &octetStringIE{Value: msg.Target2SourceNGRANnodeTranspContainer},
	})
	if msg.UEContextKeptIndicator != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_UEContextKeptIndicator)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.UEContextKeptIndicator,
		})
	}
	if msg.CriticalityDiagnostics != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_CriticalityDiagnostics)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.CriticalityDiagnostics,
		})
	}
	if msg.DRBsTransferredToMN != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_DRBsTransferredToMN)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.DRBsTransferredToMN,
		})
	}
	if msg.DAPSResponseInfoList != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_DAPSResponseInfoList)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.DAPSResponseInfoList,
		})
	}
	if msg.CHOinformationAck != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_CHOinformationAck)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.CHOinformationAck,
		})
	}
	if msg.MBSSessionInformationResponseList != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_MBSSessionInformationResponseList)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.MBSSessionInformationResponseList,
		})
	}
	if msg.RRCConfigIndication != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_RRCConfigIndication)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.RRCConfigIndication,
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

func (msg *HandoverRequestAcknowledge) Encode(w io.Writer) error {
	inner := per.NewEncoder(per.APER)
	if err := encodeBody(msg, inner); err != nil {
		return err
	}

	outer := per.NewEncoder(per.APER)

	if err := choiceXnMsg(outer, common.SuccessfulOutcome); err != nil {
		return err
	}
	if err := procedureCodeXnMsg(outer, common.ProcedureCode_HandoverPreparation); err != nil {
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

func (msg *HandoverRequestAcknowledge) Decode(wire []byte) error {
	d := per.NewDecoder(wire, per.APER)

	seq := d.NewSequenceDecoder(per.SequenceConstraints{Extensible: true})
	if err := seq.DecodeExtensionBit(); err != nil {
		return fmt.Errorf("decode HandoverRequestAcknowledge extension bit: %w", err)
	}

	seqOf := d.NewSequenceOfDecoder(per.SizeConstraints{
		Extensible: false,
		Min:        common.Ptr(int64(0)),
		Max:        common.Ptr(int64(MaxProtocolIEs)),
	})
	n, err := seqOf.DecodeLength()
	if err != nil {
		return fmt.Errorf("decode HandoverRequestAcknowledge protocolIEs length: %w", err)
	}

	seen := make(map[int64]bool, n)
	for i := range n {
		id, err := d.DecodeInteger(protocolIEIDConstraints)
		if err != nil {
			return fmt.Errorf("decode HandoverRequestAcknowledge IE[%d] id: %w", i, err)
		}
		seen[id] = true

		var criticality Criticality
		if err := criticality.Decode(d); err != nil {
			return fmt.Errorf("decode HandoverRequestAcknowledge IE[%d] criticality: %w", i, err)
		}

		payload, err := d.DecodeOpenType()
		if err != nil {
			return fmt.Errorf("decode HandoverRequestAcknowledge IE[%d] value: %w", i, err)
		}
		ies := per.NewDecoder(payload, per.APER)

		switch id {
		case common.ProtocolIEID_SourceNGRANnodeUEXnAPID:
			err = msg.SourceNGRANnodeUEXnAPID.Decode(ies)
		case common.ProtocolIEID_TargetNGRANnodeUEXnAPID:
			err = msg.TargetNGRANnodeUEXnAPID.Decode(ies)
		case common.ProtocolIEID_PDUSessionResourcesAdmittedList:
			err = msg.PDUSessionResourcesAdmittedList.Decode(ies)
		case common.ProtocolIEID_PDUSessionResourcesNotAdmittedList:
			v := &PDUSessionResourcesNotAdmittedList{}
			err = v.Decode(ies)
			msg.PDUSessionResourcesNotAdmittedList = v
		case common.ProtocolIEID_Target2SourceNGRANnodeTranspContainer:
			v := octetStringIE{}
			err = v.Decode(ies)
			msg.Target2SourceNGRANnodeTranspContainer = v.Value
		case common.ProtocolIEID_UEContextKeptIndicator:
			v := &UEContextKeptIndicator{}
			err = v.Decode(ies)
			msg.UEContextKeptIndicator = v
		case common.ProtocolIEID_CriticalityDiagnostics:
			v := &CriticalityDiagnostics{}
			err = v.Decode(ies)
			msg.CriticalityDiagnostics = v
		case common.ProtocolIEID_DRBsTransferredToMN:
			v := &DRBList{}
			err = v.Decode(ies)
			msg.DRBsTransferredToMN = v
		case common.ProtocolIEID_DAPSResponseInfoList:
			v := &DAPSResponseInfoList{}
			err = v.Decode(ies)
			msg.DAPSResponseInfoList = v
		case common.ProtocolIEID_CHOinformationAck:
			v := &CHOinformationAck{}
			err = v.Decode(ies)
			msg.CHOinformationAck = v
		case common.ProtocolIEID_MBSSessionInformationResponseList:
			v := &MBSSessionInformationResponseList{}
			err = v.Decode(ies)
			msg.MBSSessionInformationResponseList = v
		case common.ProtocolIEID_RRCConfigIndication:
			v := &RRCConfigIndication{}
			err = v.Decode(ies)
			msg.RRCConfigIndication = v
		case common.ProtocolIEID_PDUSetbasedHandlingIndicator:
			v := &PDUSetbasedHandlingIndicator{}
			err = v.Decode(ies)
			msg.PDUSetbasedHandlingIndicator = v
		default:
			if criticality.Value == CriticalityReject {
				return fmt.Errorf("unsupported HandoverRequestAcknowledge IE id=%d with reject criticality", id)
			}
		}
		if err != nil {
			return fmt.Errorf("decode HandoverRequestAcknowledge IE id=%d: %w", id, err)
		}
	}

	if _, err := seq.DecodeExtensionAdditions(); err != nil {
		return fmt.Errorf("decode HandoverRequestAcknowledge extensions: %w", err)
	}

	for _, id := range []int64{
		common.ProtocolIEID_SourceNGRANnodeUEXnAPID,
		common.ProtocolIEID_TargetNGRANnodeUEXnAPID,
		common.ProtocolIEID_PDUSessionResourcesAdmittedList,
		common.ProtocolIEID_Target2SourceNGRANnodeTranspContainer,
	} {
		if !seen[id] {
			return fmt.Errorf("missing mandatory HandoverRequestAcknowledge IE id=%d", id)
		}
	}

	return nil
}
