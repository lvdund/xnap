package ies

import (
	"fmt"
	"io"

	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

// PartialUEContextTransferAcknowledge ::= SEQUENCE {
//	protocolIEs  ProtocolIE-Container {{PartialUEContextTransferAcknowledge-IEs}},
//	...
// }

type PartialUEContextTransferAcknowledge struct {
	NewNGRANnodeUEXnAPID     NGRANnodeUEXnAPID
	OldNGRANnodeUEXnAPID     NGRANnodeUEXnAPID
	SDTDataForwardingDRBList *SDTDataForwardingDRBList
	CriticalityDiagnostics   *CriticalityDiagnostics
	SRSConfiguration         *SRSConfiguration
}

func (msg *PartialUEContextTransferAcknowledge) toIEs() []XnAPMessageIE {
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
	if msg.SDTDataForwardingDRBList != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SDTDataForwardingDRBList)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.SDTDataForwardingDRBList,
		})
	}
	if msg.CriticalityDiagnostics != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_CriticalityDiagnostics)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.CriticalityDiagnostics,
		})
	}
	if msg.SRSConfiguration != nil {
		ies = append(ies, XnAPMessageIE{
			ID:          ProtocolIEID{Value: int64(common.ProtocolIEID_SRSConfiguration)},
			Criticality: Criticality{Value: CriticalityIgnore},
			Value:       msg.SRSConfiguration,
		})
	}
	return ies
}

func (msg *PartialUEContextTransferAcknowledge) Encode(w io.Writer) error {
	inner := per.NewEncoder(per.APER)
	if err := encodeBody(msg, inner); err != nil {
		return err
	}

	outer := per.NewEncoder(per.APER)

	if err := choiceXnMsg(outer, common.SuccessfulOutcome); err != nil {
		return err
	}
	if err := procedureCodeXnMsg(outer, common.ProcedureCode_PartialUEContextTransfer); err != nil {
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

func (msg *PartialUEContextTransferAcknowledge) Decode(wire []byte) error {
	d := per.NewDecoder(wire, per.APER)

	seq := d.NewSequenceDecoder(per.SequenceConstraints{Extensible: true})
	if err := seq.DecodeExtensionBit(); err != nil {
		return fmt.Errorf("decode PartialUEContextTransferAcknowledge extension bit: %w", err)
	}

	seqOf := d.NewSequenceOfDecoder(per.SizeConstraints{
		Extensible: false,
		Min:        common.Ptr(int64(0)),
		Max:        common.Ptr(int64(MaxProtocolIEs)),
	})
	n, err := seqOf.DecodeLength()
	if err != nil {
		return fmt.Errorf("decode PartialUEContextTransferAcknowledge protocolIEs length: %w", err)
	}

	seen := make(map[int64]bool, n)
	for i := range n {
		id, err := d.DecodeInteger(protocolIEIDConstraints)
		if err != nil {
			return fmt.Errorf("decode PartialUEContextTransferAcknowledge IE[%d] id: %w", i, err)
		}
		seen[id] = true

		var criticality Criticality
		if err := criticality.Decode(d); err != nil {
			return fmt.Errorf("decode PartialUEContextTransferAcknowledge IE[%d] criticality: %w", i, err)
		}

		payload, err := d.DecodeOpenType()
		if err != nil {
			return fmt.Errorf("decode PartialUEContextTransferAcknowledge IE[%d] value: %w", i, err)
		}
		ies := per.NewDecoder(payload, per.APER)

		switch id {
		case common.ProtocolIEID_NewNGRANnodeUEXnAPID:
			err = msg.NewNGRANnodeUEXnAPID.Decode(ies)
		case common.ProtocolIEID_OldNGRANnodeUEXnAPID:
			err = msg.OldNGRANnodeUEXnAPID.Decode(ies)
		case common.ProtocolIEID_SDTDataForwardingDRBList:
			v := &SDTDataForwardingDRBList{}
			err = v.Decode(ies)
			msg.SDTDataForwardingDRBList = v
		case common.ProtocolIEID_CriticalityDiagnostics:
			v := &CriticalityDiagnostics{}
			err = v.Decode(ies)
			msg.CriticalityDiagnostics = v
		case common.ProtocolIEID_SRSConfiguration:
			v := &SRSConfiguration{}
			err = v.Decode(ies)
			msg.SRSConfiguration = v
		default:
			if criticality.Value == CriticalityReject {
				return fmt.Errorf("unsupported PartialUEContextTransferAcknowledge IE id=%d with reject criticality", id)
			}
		}
		if err != nil {
			return fmt.Errorf("decode PartialUEContextTransferAcknowledge IE id=%d: %w", id, err)
		}
	}

	if _, err := seq.DecodeExtensionAdditions(); err != nil {
		return fmt.Errorf("decode PartialUEContextTransferAcknowledge extensions: %w", err)
	}

	for _, id := range []int64{
		common.ProtocolIEID_NewNGRANnodeUEXnAPID,
		common.ProtocolIEID_OldNGRANnodeUEXnAPID,
	} {
		if !seen[id] {
			return fmt.Errorf("missing mandatory PartialUEContextTransferAcknowledge IE id=%d", id)
		}
	}

	return nil
}
