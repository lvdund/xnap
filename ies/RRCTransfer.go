package ies

import (
	"fmt"
	"io"

	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

// RRCTransfer ::= SEQUENCE {
//	protocolIEs  ProtocolIE-Container {{RRCTransfer-IEs}},
//	...
// }

type RRCTransfer struct {
	MNGRANnodeUEXnAPID               NGRANnodeUEXnAPID
	SNGRANnodeUEXnAPID               NGRANnodeUEXnAPID
	SplitSRBRRCTransfer              *SplitSRBRRCTransfer
	UEReportRRCTransfer              *UEReportRRCTransfer
	FastMCGRecoveryRRCTransferSNToMN *FastMCGRecoveryRRCTransfer
	FastMCGRecoveryRRCTransferMNToSN *FastMCGRecoveryRRCTransfer
	SDTSRBBetweenNewNodeOldNode      *SDTSRBBetweenNewNodeOldNode
	QoEMeasurementResults            *QoEMeasurementResults
}

func (msg *RRCTransfer) toIEs() []XnAPMessageIE {
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
	if msg.SplitSRBRRCTransfer != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SplitSRBRRCTransfer)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.SplitSRBRRCTransfer,
		})
	}
	if msg.UEReportRRCTransfer != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_UEReportRRCTransfer)},
			Criticality: Criticality{Value: CriticalityReject},
			Value:       msg.UEReportRRCTransfer,
		})
	}
	if msg.FastMCGRecoveryRRCTransferSNToMN != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_FastMCGRecoveryRRCTransferSNToMN)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.FastMCGRecoveryRRCTransferSNToMN,
		})
	}
	if msg.FastMCGRecoveryRRCTransferMNToSN != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_FastMCGRecoveryRRCTransferMNToSN)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.FastMCGRecoveryRRCTransferMNToSN,
		})
	}
	if msg.SDTSRBBetweenNewNodeOldNode != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SDTSRBBetweenNewNodeOldNode)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.SDTSRBBetweenNewNodeOldNode,
		})
	}
	if msg.QoEMeasurementResults != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_QoEMeasurementResults)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.QoEMeasurementResults,
		})
	}
	return ies
}

func (msg *RRCTransfer) Encode(w io.Writer) error {
	inner := per.NewEncoder(per.APER)
	if err := encodeBody(msg, inner); err != nil {
		return err
	}

	outer := per.NewEncoder(per.APER)

	if err := choiceXnMsg(outer, common.InitiatingMessage); err != nil {
		return err
	}
	if err := procedureCodeXnMsg(outer, common.ProcedureCode_RRCTransfer); err != nil {
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

func (msg *RRCTransfer) Decode(wire []byte) error {
	d := per.NewDecoder(wire, per.APER)

	seq := d.NewSequenceDecoder(per.SequenceConstraints{Extensible: true})
	if err := seq.DecodeExtensionBit(); err != nil {
		return fmt.Errorf("decode RRCTransfer extension bit: %w", err)
	}

	seqOf := d.NewSequenceOfDecoder(per.SizeConstraints{
		Extensible: false,
		Min:        common.Ptr(int64(0)),
		Max:        common.Ptr(int64(MaxProtocolIEs)),
	})
	n, err := seqOf.DecodeLength()
	if err != nil {
		return fmt.Errorf("decode RRCTransfer protocolIEs length: %w", err)
	}

	seen := make(map[int64]bool, n)
	for i := range n {
		id, err := d.DecodeInteger(protocolIEIDConstraints)
		if err != nil {
			return fmt.Errorf("decode RRCTransfer IE[%d] id: %w", i, err)
		}
		seen[id] = true

		var criticality Criticality
		if err := criticality.Decode(d); err != nil {
			return fmt.Errorf("decode RRCTransfer IE[%d] criticality: %w", i, err)
		}

		payload, err := d.DecodeOpenType()
		if err != nil {
			return fmt.Errorf("decode RRCTransfer IE[%d] value: %w", i, err)
		}
		ies := per.NewDecoder(payload, per.APER)

		switch id {
		case common.ProtocolIEID_MNGRANnodeUEXnAPID:
			err = msg.MNGRANnodeUEXnAPID.Decode(ies)
		case common.ProtocolIEID_SNGRANnodeUEXnAPID:
			err = msg.SNGRANnodeUEXnAPID.Decode(ies)
		case common.ProtocolIEID_SplitSRBRRCTransfer:
			v := &SplitSRBRRCTransfer{}
			err = v.Decode(ies)
			msg.SplitSRBRRCTransfer = v
		case common.ProtocolIEID_UEReportRRCTransfer:
			v := &UEReportRRCTransfer{}
			err = v.Decode(ies)
			msg.UEReportRRCTransfer = v
		case common.ProtocolIEID_FastMCGRecoveryRRCTransferSNToMN:
			v := &FastMCGRecoveryRRCTransfer{}
			err = v.Decode(ies)
			msg.FastMCGRecoveryRRCTransferSNToMN = v
		case common.ProtocolIEID_FastMCGRecoveryRRCTransferMNToSN:
			v := &FastMCGRecoveryRRCTransfer{}
			err = v.Decode(ies)
			msg.FastMCGRecoveryRRCTransferMNToSN = v
		case common.ProtocolIEID_SDTSRBBetweenNewNodeOldNode:
			v := &SDTSRBBetweenNewNodeOldNode{}
			err = v.Decode(ies)
			msg.SDTSRBBetweenNewNodeOldNode = v
		case common.ProtocolIEID_QoEMeasurementResults:
			v := &QoEMeasurementResults{}
			err = v.Decode(ies)
			msg.QoEMeasurementResults = v
		default:
			if criticality.Value == CriticalityReject {
				return fmt.Errorf("unsupported RRCTransfer IE id=%d with reject criticality", id)
			}
		}
		if err != nil {
			return fmt.Errorf("decode RRCTransfer IE id=%d: %w", id, err)
		}
	}

	if _, err := seq.DecodeExtensionAdditions(); err != nil {
		return fmt.Errorf("decode RRCTransfer extensions: %w", err)
	}

	for _, id := range []int64{
		common.ProtocolIEID_MNGRANnodeUEXnAPID,
		common.ProtocolIEID_SNGRANnodeUEXnAPID,
	} {
		if !seen[id] {
			return fmt.Errorf("missing mandatory RRCTransfer IE id=%d", id)
		}
	}

	return nil
}
